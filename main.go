package main

import (
	"github.com/josebruno2020/go-api-gin/database"
	"github.com/josebruno2020/go-api-gin/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequest()
}
