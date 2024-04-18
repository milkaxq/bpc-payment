package config

type BankModel struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	OrderID      string `json:"order_id"`
	OTPRequestID string `json:"otp_request_id"`
}
