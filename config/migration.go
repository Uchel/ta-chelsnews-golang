package config

import (
	"chelsnews/models/entity"
	"fmt"

	"log"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(&entity.News{}, entity.NewsCategory{}, entity.SubCategory{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Migration Success")

}
