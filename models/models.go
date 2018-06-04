package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Miner struct {
	gorm.Model
	id           int
	MinerAddress string
	BlockNumber  int64
	Reward       int64
}
