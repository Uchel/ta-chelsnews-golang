package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseInit() *gorm.DB {
	var db *gorm.DB
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/latihan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to database")
	return db
}
