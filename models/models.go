package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MinerDetail struct {
	Id           int64  `gorm:"column:Id"`
	MinerAddress string `gorm:"column:MinerAddress"`
	BlockNumber  int64  `gorm:"column:BlockNumber"`
	Reward       int64  `gorm:"column:Reward"`
}
