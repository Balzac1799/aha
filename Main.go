package main

import (
	"aha/urls"
	"aha/utils"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	db := utils.CreateDBConn()
	defer db.Close()
	server := urls.SetUpUrls()
	server.Run()
}
