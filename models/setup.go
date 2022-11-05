package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabse() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error load .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database :", Dbdriver)
		log.Fatal("Connection error:", err)
	} else {
		fmt.Println("Success connect to database :", Dbdriver)
	}

	DB.AutoMigrate(&User{})
}
