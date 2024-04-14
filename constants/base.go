package constants

import "os"

type BaseStruct struct {
	// bank url < epg_server_IP_address>
	BaseURL string `json:"base_url" validate:"required" binding:"required"`
}

var Base BaseStruct

func InitBase() {
	Base.BaseURL = os.Getenv("BASE_URL")
}
