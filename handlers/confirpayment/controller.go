package confirpayment

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
)

// Last step of payment
// @Summary Confirm Payment Make This Request Third
// @Description Confirm Payment by writing otp code.
// @Tags epg
// @Accept json
// @Produce json
// @Param requestBody body ConfirmPaymentRequest true "Payment Confirmation request body"
// @Success 200 {object} constants.ResponseWithMessage "Just message that says it was succesfully"
// @Failure 500 {object} constants.HttpError "Some Confirmation Error"
// @Router /confirm-payment [post]
func ConfirmPayment(c *gin.Context) {
	var confirPaymentRequest ConfirmPaymentRequest

	if err := c.BindJSON(&confirPaymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}

	bankModel, expiresAt, err := utils.GetBankModel(confirPaymentRequest.MDORDER)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Println(bankModel)
	fmt.Println(expiresAt)
	confirPaymentRequest.RequestID = bankModel.OTPRequestID

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

	c.JSON(http.StatusOK, gin.H{"message": resp})
}
