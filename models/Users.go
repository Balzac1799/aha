package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(30);not null"`
	Password string `gorm:"type:varchar(128);not null"`
	AhaMoments []AhaMoment `gorm:"foreignkey:Author"`
}