package refund

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
)

// @Summary Check order status. Make this request finally
// @Description Check's order status from bank.
// @Tags epg
// @Accept json
// @Produce json
// @Param requestBody body RefundRequestBody true "amount and order id of refunding order"
// @Success 200 {object} constants.ResponseWithMessage "Just message that says it was succesfully"
// @Failure 400 {object} constants.HttpError "Invalid request body or parameters"
// @Router /refund [post]
func Refund(c *gin.Context) {
	var refundRequest RefundRequestBody

	if err := c.BindJSON(&refundRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}

	bankModel, err := utils.GetBankModel(refundRequest.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	resp, err := makeRefundRequest(refundRequest, bankModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp})
}
