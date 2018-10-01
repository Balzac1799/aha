package scripts

import (
	"aha/models"
	"github.com/jinzhu/gorm"
)

func CreateUsers() {
	user := models.User{UserName:"高小琴", Password:"123456"}
	user1 := models.User{UserName:"祁同伟", Password:"123456"}
	user2 := models.User{UserName:"赵立春", Password:"123456"}
	user3 := models.User{UserName:"沙瑞金", Password:"123456"}
	user4 := models.User{UserName:"侯亮平", Password:"123456"}
	user5 := models.User{UserName:"陈海", Password:"123456"}
	user6 := models.User{UserName:"高玉良", Password:"123456"}
	user7 := models.User{UserName:"李达康", Password:"123456"}
	user8 := models.User{UserName:"赵瑞龙", Password:"123456"}

	db, err := gorm.Open("mysql", "root:development123@(127.0.0.1:3306)/aha_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.NewRecord(user)
	db.NewRecord(user1)
	db.NewRecord(user2)
	db.NewRecord(user3)
	db.NewRecord(user4)
	db.NewRecord(user5)
	db.NewRecord(user6)
	db.NewRecord(user7)
	db.NewRecord(user8)

	db.Create(&user)
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)
	db.Create(&user4)
	db.Create(&user5)
	db.Create(&user6)
	db.Create(&user7)
	db.Create(&user8)
}