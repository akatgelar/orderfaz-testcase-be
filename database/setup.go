package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
 
var DB_POSTGRES *gorm.DB
 
func ConnnectDatabasePostgres() { 
	err := godotenv.Load()
	if err != nil { 
	  fmt.Println(err)
	}
	
	DB_POSTGRES_HOST := os.Getenv("DB_POSTGRES_HOST")
	DB_POSTGRES_PORT := os.Getenv("DB_POSTGRES_PORT")
	DB_POSTGRES_USER := os.Getenv("DB_POSTGRES_USER")
	DB_POSTGRES_PASS := os.Getenv("DB_POSTGRES_PASS")
	DB_POSTGRES_DATABASE := os.Getenv("DB_POSTGRES_DATABASE")
 
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host="+DB_POSTGRES_HOST+" user="+DB_POSTGRES_USER+" password="+DB_POSTGRES_PASS+" dbname="+DB_POSTGRES_DATABASE+" port="+DB_POSTGRES_PORT+" sslmode=disable TimeZone=Asia/Jakarta", 
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
  

	if err != nil {
        fmt.Println(err)
    }
 
	DB_POSTGRES = db
}
