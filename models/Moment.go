package models

import "github.com/jinzhu/gorm"

type AhaMoment struct {
	gorm.Model
	CategoryCode int
	CategoryValue string `gorm:"type:varchar(12);default:'一般';not null"`
	Text string `gorm:"type:varchar(300);default:'默认值';not null"`
	Tags string `gorm:"type:varchar(10);default:'默认值';not null"`
	WeatherCode int
	WeatherValue string `gorm:"type:varchar(12);default:'阴天';not null"`
	Location string `gorm:"type:varchar(16);default:'出租屋';not null"`
	Author int `gorm:"type:int(10) unsigned"`
}