package submitcard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
)

// Submit Card handles submit form of card
// @Summary Card Submission Make This Request Second
// @Description Submit card data to move next into step.
// @Tags epg
// @Accept json
// @Produce json
// @Param requestBody body SubmitCardRequest true "Card submission request body"
// @Success 200 {object} constants.ResponseWithMessage "Just message that says it was succesfully"
// @Failure 400 {object} constants.HttpError "Submit card to bank error"
// @Failure 500 {object} constants.HttpError "Get otp error"
// @Router /submit-card [post]
func SubmitCard(c *gin.Context) {
	var submitCardRequest SubmitCardRequest

	if err := c.BindJSON(&submitCardRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}
	// impossibly important
	submitCardRequest.Language = "ru"

	bankModel, err := utils.GetBankModel(submitCardRequest.MDORDER)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	submitCardResponse, err := SubmitCardToBank(bankModel.ApiClient, submitCardRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestID, err := getOTPPassword(submitCardResponse, submitCardRequest.MDORDER, bankModel.ApiClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bankModel.OTPRequestID = requestID
	err = addOTPRequestID(bankModel.OrderID, bankModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succesfully sended OTP"})
}
