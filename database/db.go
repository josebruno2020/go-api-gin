package database

import (
	"fmt"
	"os"

	"github.com/josebruno2020/go-api-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo", host, port, user, password, dbName)
	DB, err = gorm.Open(postgres.Open(dns))

	if err != nil {
		panic(fmt.Sprintf("Error to connect to DB: %s", err.Error()))
	}

	DB.AutoMigrate(&models.Student{})
}
