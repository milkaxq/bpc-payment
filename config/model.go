package config

type BankModel struct {
	ApiClient    string `json:"api_client"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	OrderID      string `json:"order_id"`
	OTPRequestID string `json:"otp_request_id"`
}
