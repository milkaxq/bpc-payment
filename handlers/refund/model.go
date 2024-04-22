package refund

type RefundRequestBody struct {
	OrderID string `json:"order_id"`
	Amount  string `json:"amount"`
}

type RefundRequestBodyOfBank struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	OrderID  string `json:"orderId"`
	Amount   string `json:"amount"`
}

type RefundBankResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
