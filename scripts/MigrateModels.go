package scripts

import (
	"aha/models"
	"github.com/jinzhu/gorm"
)

func Migrate() {
	db, err := gorm.Open("mysql", "root:development123@(127.0.0.1:3306)/aha_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.AhaMoment{})
}
