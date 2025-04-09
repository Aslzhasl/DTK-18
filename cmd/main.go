package main

import (
	"guilt-type-service/config"
	"guilt-type-service/database"
	"guilt-type-service/routes"

	"log"
)

func main() {
	config.LoadEnv()
	DB := database.ConnectDB()
	r := routes.SetupRouter(DB)
	if err := r.Run(); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
