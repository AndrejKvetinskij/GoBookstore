package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// 'mqbook'@'%' IDENTIFIED BY 'Mqbook@123'
func Connect() {
	d, err := gorm.Open("mysql", "mqbook1:Mqbook@1234@/bookstoreDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
