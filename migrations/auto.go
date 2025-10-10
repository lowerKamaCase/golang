package main

import (
	"lowerkamacase/golang/pkg/link"
	"lowerkamacase/golang/pkg/product"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&link.Link{})
	db.AutoMigrate(&product.Product{})
}
