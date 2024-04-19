package checkstatus

import "fmt"

type OrderStatusRequest struct {
	// bank merchant user name
	Username string `json:"userName"`
	// bank merchant passowrd
	Password string `json:"password"`
	// Order id the same as order number
	OrderID string `json:"orderId" binding:"required"`
}

type OrderStatusResponse struct {
	ErrorCode             string                `json:"errorCode"`
	ErrorMessage          string                `json:"errorMessage"`
	OrderNumber           string                `json:"orderNumber"`
	OrderStatus           int                   `json:"orderStatus"`
	ActionCode            int                   `json:"actionCode"`
	ActionCodeDescription string                `json:"actionCodeDescription"`
	Amount                float64               `json:"amount"`
	Currency              string                `json:"currency"`
	Date                  int64                 `json:"date"`
	OrderDescription      string                `json:"orderDescription"`
	Ip                    string                `json:"ip"`
	MerchantOrderParams   []*MerchantOrderParam `json:"merchantOrderParams"`
	Attributes            []*Attribute          `json:"attributes"`
	CardAuthInfo          *CardAuthInfo         `json:"cardAuthInfo"`
	FraudLevel            int                   `json:"fraudLevel"`
}

type MerchantOrderParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CardAuthInfo struct {
	Expiration              string `json:"expiration"`
	CardholderName          string `json:"cardholderName"`
	AuthorizationResponseId string `json:"authorizationResponseId"`
	Pan                     string `json:"pan"`
}

func (c OrderStatusResponse) getErrorString() string {
	switch c.ErrorCode {
	case "0":
		return "No system error."
	case "1":
		return "Expected orderId or orderNumber."
	case "2":
		return "Order is declined due to an error in the payment credential."
	case "5":
		return "Access is denied. The user who processes payments for the merchant must change their password."
	case "6":
		return "Unregistered orderId."
	case "7":
		return "System error."
	default:
		return fmt.Sprintf("Unknown error code: %s", c.ErrorCode)
	}
}
