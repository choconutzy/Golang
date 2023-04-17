package main

import (
	"Chapter2-challenge2/config"
	"Chapter2-challenge2/routers"
	"log"
)

func main() {
	configDB, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	config.ConnectDB(&configDB)
	PORT := ":8000"
	routers.StartServer().Run(PORT)
}
