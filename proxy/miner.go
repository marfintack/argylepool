package proxy

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ethereum/ethash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/marfintack/argylepool/connector"
	"github.com/marfintack/argylepool/models"
)

var hasher = ethash.New()

func (s *ProxyServer) processShare(login, id, ip string, t *BlockTemplate, params []string) (bool, bool) {
	nonceHex := params[0]
	hashNoNonce := params[1]
	mixDigest := params[2]
	nonce, _ := strconv.ParseUint(strings.Replace(nonceHex, "0x", "", -1), 16, 64)
	shareDiff := s.config.Proxy.Difficulty

	h, ok := t.headers[hashNoNonce]
	if !ok {
		log.Printf("Stale share from %v@%v", login, ip)
		return false, false
	}

	share := Block{
		number:      h.height,
		hashNoNonce: common.HexToHash(hashNoNonce),
		difficulty:  big.NewInt(shareDiff),
		nonce:       nonce,
		mixDigest:   common.HexToHash(mixDigest),
	}

	block := Block{
		number:      h.height,
		hashNoNonce: common.HexToHash(hashNoNonce),
		difficulty:  h.diff,
		nonce:       nonce,
		mixDigest:   common.HexToHash(mixDigest),
	}

	if !hasher.Verify(share) {
		return false, false
	}

	if hasher.Verify(block) {
		ok, err := s.rpc().SubmitBlock(params)
		if err != nil {
			log.Printf("Block submission failure at height %v for %v: %v", h.height, t.Header, err)
		} else if !ok {
			log.Printf("Block rejected at height %v for %v", h.height, t.Header)
			return false, false
		} else {
			s.fetchBlockTemplate()
			exist, err := s.backend.WriteBlock(login, id, params, shareDiff, h.diff.Int64(), h.height, s.hashrateExpiration)
			if exist {
				return true, false
			}
			if err != nil {
				log.Println("Failed to insert block candidate into backend:", err)
			} else {
				// Getting Db connection
				db := connector.GetConnection()
				log.Printf("Inserted block %v to backend", h.height)
				minerRewardModel := models.MinerReward{}
				db.First(&minerRewardModel)
				reward := minerRewardModel.RewardValue
				log.Printf("Miner Reward is set to %s", reward)
				fmt.Println("Login %v", login)
				fmt.Println("Id %v ", ip)
				fmt.Println("Share Diff %v", shareDiff)
				fmt.Println("Height %v", h.height)
				fmt.Println("Hash Rate %v", s.hashrateExpiration)

				Miner := models.MinerDetail{MinerAddress: login, MinerIp: ip, HashRate: shareDiff, BlockNumber: h.height, Reward: reward}
				db.Save(&Miner)
				apiUrl := "https://admin.argylecoin.com"
				resource := "/transferTokenAdmin/"
				data := url.Values{}
				data.Set("tokens", reward)
				data.Add("toAddress", login)

				u, _ := url.ParseRequestURI(apiUrl)
				u.Path = resource
				urlStr := u.String() // 'https://api.com/user/'

				client := &http.Client{}
				newr, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
				newr.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				newr.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
				response, err := client.Do(newr)
				if err != nil {
					fmt.Printf("The HTTP request failed with error %s\n", err)
				} else {
					data, _ := ioutil.ReadAll(response.Body)
					fmt.Println("Argle Transferred Successfully To the Miner")
					fmt.Println(string(data))

					//respondJSON(w, http.StatusOK, string(data))
				}
			}
			log.Printf("Block found by miner %v@%v at height %d", login, ip, h.height)
		}
	} else {
		exist, err := s.backend.WriteShare(login, id, params, shareDiff, h.height, s.hashrateExpiration)
		if exist {
			fmt.Println("Block Is Mined Successfully")
			fmt.Println("Login %v", login)
			fmt.Println("Id %v ", ip)
			fmt.Println("Share Diff %v", shareDiff)
			fmt.Println("Height %v", h.height)
			fmt.Println("Hash Rate %v", s.hashrateExpiration)
			return true, false
		}
		if err != nil {
			log.Println("Failed to insert share data into backend:", err)
		}
	}
	return false, true
}
