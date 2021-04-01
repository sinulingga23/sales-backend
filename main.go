package main

import (
	"log"
	"time"

	"sales-backend/model"
	"sales-backend/controller"
	"github.com/joho/godotenv"
)

func init() {
	load := godotenv.Load()
	if load != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {
	controller.RunServer()
}
