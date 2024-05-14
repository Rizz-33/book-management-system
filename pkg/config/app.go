package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "admin:password/simplerest?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
    return db
}