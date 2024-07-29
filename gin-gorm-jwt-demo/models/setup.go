package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/userdemo?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Open DB failed:", err.Error())
	} else {
		log.Println("Open DB:", dsn)
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Token{})
}
