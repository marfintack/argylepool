package connector

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/marfintack/argylepool/config"
)

func getConnection() *gorm.DB {
	config := config.GetConfig()
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database %s", err)
	}
	return db
}
