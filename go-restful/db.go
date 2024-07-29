package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456789@/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("fail to connect db: " + err.Error())
	}

	db.AutoMigrate(&todoModel{})
}
