package refund

type RefundRequestBody struct {
	Username string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
	OrderID  string `json:"orderId" binding:"required"`
	Amount   string `json:"amount" binding:"required"`
}

type RefundBankResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
