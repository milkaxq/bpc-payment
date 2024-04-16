package submitcard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
)

type RequestIdResponse struct {
	RequestID string `json:"request_id"`
}

// Submit Card handles submit form of card
// @Summary Card Submission
// @Description Submit card data to move next into step.
// @Tags epg
// @Accept json
// @Produce json
// @Param requestBody body SubmitCardRequest true "Card submission request body"
// @Success 200 {object} RequestIdResponse "Coming request id of otp to use in next request confirmPayment"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	request_id, err := GetOTPPassword(submitCardResponse, submitCardRequest.MDORDER)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"request_id": request_id})
}
