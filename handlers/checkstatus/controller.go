package checkstatus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
)

// Check Status checks order status from bank.
// @Summary Check order status. Make this request finally
// @Description Check's order status from bank.
// @Tags epg
// @Accept json
// @Produce json
// @Param requestBody body OrderStatusRequest true "check order status request body"
// @Success 200 {object} OrderStatusResponse "Successfully checked status"
// @Failure 400 {object} constants.HttpError "Invalid request body or parameters"
// @Router /check-status [post]
func CheckStatus(c *gin.Context) {
	var orderStatusRequest OrderStatusRequest

	if err := c.BindJSON(&orderStatusRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}

	orderStatusResponse, err := CheckOrderStatus(orderStatusRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orderStatusResponse)
}
