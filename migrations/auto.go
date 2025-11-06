package main

import (
	"lowerkamacase/golang/internal/user"
	"lowerkamacase/golang/pkg/link"
	"lowerkamacase/golang/pkg/stat"
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

	db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
}
