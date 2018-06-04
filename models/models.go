package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MinerDetail struct {
	Id           int64  `gorm:"primary_key;AUTO_INCREMENT;column:Id"`
	MinerAddress string `gorm:"column:MinerAddress"`
	BlockNumber  int64  `gorm:"column:BlockNumber"`
	Reward       int64  `gorm:"column:Reward"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&MinerDetail{})
	log.Printf("Models Created")
	return db
}
