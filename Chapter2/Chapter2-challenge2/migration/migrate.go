package main

import (
	"Chapter2-challenge2/config"
	"Chapter2-challenge2/models"
	"fmt"
	"log"
)

func init() {
	configDB, err := config.LoadConfig("..")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	config.ConnectDB(&configDB)
}

func main() {
	config.DB.AutoMigrate(&models.Book{})
	fmt.Println("? Migration complete")
}
