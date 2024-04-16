package confirpayment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
)

// Last step of payment
// @Summary Confirm Payment
// @Description Confirm Payment by writing otp code.
// @Tags epg
// @Accept json
// @Produce json
// @Param requestBody body ConfirmPaymentRequest true "Payment Confirmation request body"
// @Success 200 {string} string "A string of success payment"
// @Failure 500 {object} constants.HttpError "Some Confirmation Error"
// @Router /confirm-payment[post]
func ConfirmPayment(c *gin.Context) {
	var confirPaymentRequest ConfirmPaymentRequest

	if err := c.BindJSON(&confirPaymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}

	paRes, err := ConfirmPaymenRequest(confirPaymentRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := FinishPayment(paRes, confirPaymentRequest.MDORDER)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
