package utils

import (
	"aha/settings"
	"fmt"
	"github.com/jinzhu/gorm"
)

func CreateDBConn() *gorm.DB {
	dbConfig := settings.GetDatabaseConfig()
	args := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Name)

	dbConn, err := gorm.Open("mysql", args)
	if err != nil {
		panic(err)
	}

	return dbConn
}