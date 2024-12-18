package constants

import (
	"log"
	"os"
)

type ResponseWithMessage struct {
	Message string `json:"message"`
}

type BaseStruct struct {
	// bank url < epg_server_IP_address>
	SenagatBaseURL string
	HalkBaseURL    string
	RysgalBaseURL  string
}

var Base BaseStruct

func InitBase() {
	if senagatBaseUrl, ok := os.LookupEnv("SENAGAT_BASE_URL"); !ok {
		log.Panic("please specify base_url in .env file")
	} else {
		Base.SenagatBaseURL = senagatBaseUrl
	}
	if halkbankBaseUrl, ok := os.LookupEnv("HALKBANK_BASE_URL"); !ok {
		log.Panic("please specify base_url in .env file")
	} else {
		Base.HalkBaseURL = halkbankBaseUrl
	}
	if rysgalBaseUrl, ok := os.LookupEnv("RYSGAL_BASE_URL"); !ok {
		log.Panic("please specify base_url in .env file")
	} else {
		Base.RysgalBaseURL = rysgalBaseUrl
	}
}
