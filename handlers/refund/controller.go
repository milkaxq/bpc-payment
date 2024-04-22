package refund

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
)

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
