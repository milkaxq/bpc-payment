package resendpassword

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
)

// @Summary Confirm Payment Make This Request Third
// @Description Confirm Payment by writing otp code.
// @Tags epg
// @Accept json
// @Produce json
// @Param requestBody body ResendPasswordRequest true "Resend Password Request body"
// @Success 200 {object} constants.ResponseWithMessage "Just message that says it was succesfully"
// @Failure 500 {object} constants.HttpError "Some Confirmation Error"
// @Router /resend-password [post]
func ResendPassword(c *gin.Context) {
	var resendPasswordRequest ResendPasswordRequest

	if err := c.BindJSON(&resendPasswordRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}

	bankModel, err := utils.GetBankModel(resendPasswordRequest.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if bankModel.OTPRequestID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrOTPRequestID})
		return
	}

	resp, err := resendOTPRequest(bankModel.OTPRequestID, bankModel.ApiClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp})
}
