package main

import (
	"github.com/joho/godotenv"
	"github.com/josebruno2020/go-api-gin/database"
	"github.com/josebruno2020/go-api-gin/routes"
)

func main() {
	godotenv.Load()
	database.ConnectDB()
	r := routes.HandleRequest()

	r.Run(":3000")
}
