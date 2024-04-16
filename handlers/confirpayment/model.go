package confirpayment

type ConfirmPaymentRequest struct {
	// Request id comes from previouse request of card submission
	RequestID string `json:"request_id"`
	// otp password that comes to phone number
	PasswordEdit string `json:"passwordEdit"`
	// md order which was created on order registration
	MDORDER string `json:"MDORDER"`
}
