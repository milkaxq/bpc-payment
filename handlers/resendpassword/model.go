package resendpassword

type ResendPasswordRequest struct {
	// The same as MDORDER
	OrderID string `json:"order_id" binding:"required"`
}
