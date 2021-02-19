package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=gin_rummy password=postgres sslmode=disable")
	if err != nil {
		fmt.Println("Something new")
		fmt.Println("Failed to connect to database!")
		panic(err.Error())
	}

	fmt.Println("Connected to database successfully!")

	db.AutoMigrate(&User{})

	return db
}
