package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d,err := gorm.Open("mysql", "nandini:nandini584/Knowhow?charset=utf8&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	db=d
}

func GetDb() *gorm.DB(
	return db
)