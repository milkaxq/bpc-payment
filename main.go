package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/milkaxq/bpcpayment/config"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/routers"
)

func main() {
	config.InitDB()
	initEnv()
	constants.InitBase()
	routers.InitRoutes()

}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
