package configs

import (
	"fmt"
	"log"
	"os"

	"prakerja/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	loadEnv()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	ConnectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser,
			dbPassword,
			dbHost,
			dbPort,
			dbDatabase,
		)
	var err error
	DB, err = gorm.Open(mysql.Open(ConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitialMigration()
}
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed env file")
	}
}

func InitialMigration() {
	err := DB.AutoMigrate(&models.Student{})
	if err != nil {
		log.Println(err)
		fmt.Println("AutoMigrate Failed!")
	} else {
		fmt.Println("AutoMigrate Succeed! ")
	}
}
