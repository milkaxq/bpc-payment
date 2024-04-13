package checkversion

type CheckVersionRequest struct {
	// bank url < epg_server_IP_address>
	BaseURL string `json:"base_url" validate:"required" binding:"required"`
}

type EPGVersionsResponse struct {
	Version1 string `json:"version1"`
	Version2 string `json:"version2"`
	Error1   string `json:"error1,omitempty"`
	Error2   string `json:"error2,omitempty"`
}
