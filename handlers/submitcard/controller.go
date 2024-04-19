package submitcard

import (
	"fmt"
	"net/http"
	"time"

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

	urlParams := utils.StructToURLParams(submitCardRequest)

	submitCardResponse, err := SubmitCardToBank(urlParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bankModel, expiresAt, err := utils.GetBankModel(submitCardRequest.MDORDER)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(bankModel)
	fmt.Println(expiresAt)
	requestID, err := GetOTPPassword(submitCardResponse, submitCardRequest.MDORDER)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bankModel.OTPRequestID = requestID

	// need to put minutes because in utility function i add entry with minustes
	err = addOTPRequestID(bankModel.OrderID, bankModel, (expiresAt-time.Now().Unix())/60)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succesfully sended OTP"})
}
