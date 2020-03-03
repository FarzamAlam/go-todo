package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load() // Loads the .env file
	if e != nil {
		fmt.Println("Error while loading the environment variable:", e)
	}
	// Get db details
	username := os.Getenv("db_username")
	password := os.Getenv("db_password")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbType := os.Getenv("db_type")
	dbName := os.Getenv("db_name")
	dbURI := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", username, password, dbHost, dbPort, dbName)
	fmt.Println("dbURI: ", dbURI)
	conn, err := gorm.Open(dbType, dbURI)
	if err != nil {
		log.Fatal("Not able to connect to database:", err)
	}

	db = conn
	db.Debug().AutoMigrate(&Task{}, &Project{})
	fmt.Println("Connection successful!")
}

func GetDB() *gorm.DB {
	return db
}
