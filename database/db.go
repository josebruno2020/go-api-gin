package database

import (
	"fmt"

	"github.com/josebruno2020/go-api-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dns := fmt.Sprintf("host=localhost user=postgres password=postgres dbname=go_rest port=5433 sslmode=disable TimeZone=America/Sao_Paulo")
	DB, err = gorm.Open(postgres.Open(dns))

	if err != nil {
		panic(fmt.Sprintf("Error to connect to DB: %s", err.Error()))
	}

	DB.AutoMigrate(&models.Student{})
}
