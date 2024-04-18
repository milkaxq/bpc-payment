package constants

import (
	"log"
	"os"
)

type BaseStruct struct {
	// bank url < epg_server_IP_address>
	SenagatBaseURL string
}

var Base BaseStruct

func InitBase() {
	if senagatBaseUrl, ok := os.LookupEnv("SENAGAT_BASE_URL"); !ok {
		log.Panic("please specify base_url in .env file")
	} else {
		Base.SenagatBaseURL = senagatBaseUrl
	}
}
