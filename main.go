package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	constants.InitBase()
	routers.InitRoutes()
}
