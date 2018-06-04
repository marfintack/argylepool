package models

import "github.com/jinzhu/gorm"

type Miner struct {
	gorm.Model
	MinerAddress string
	BlockNumber  int64
	Reward       int64
}
