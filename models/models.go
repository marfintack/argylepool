package models

import "github.com/jinzhu/gorm"

type Miner struct {
	gorm.Model
	id           int
	MinerAddress string
	BlockNumber  int64
	Reward       int64
}
