package main

import (
	"greening-india/internal/config"
	"greening-india/internal/router"

)

func main() {
	config.ConnectDB()
	config.RunMigrations()

	router := router.SetupRouter()

	router.Run(":8080")
}