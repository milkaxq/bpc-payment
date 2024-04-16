package constants

import (
	"log"
	"os"
)

type BaseStruct struct {
	// bank url < epg_server_IP_address>
	BaseURL string `json:"base_url" validate:"required" binding:"required"`
}

var Base BaseStruct

func InitBase() {
	if baseUrl, ok := os.LookupEnv("BASE_URL"); !ok {
		log.Panic("please specify base_url in .env file")
	} else {
		Base.BaseURL = baseUrl
	}
}
