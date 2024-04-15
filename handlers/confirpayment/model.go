package confirpayment

type ConfirmPaymentRequest struct {
	RequestID    string `json:"request_id"`
	PasswordEdit string `json:"passwordEdit"`
	MDORDER      string `json:"MDORDER"`
}
