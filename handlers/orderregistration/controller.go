package orderregistration

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
	"github.com/milkaxq/bpcpayment/utils"
)

// OrderRegistration handles the registration of an order.
// @Summary Order Registration Make this Request First
// @Description Register an order for payment processing.
// @Tags epg
// @Accept json
// @Produce json
// @Param requestBody body OrderRegistrationRequest true "Order Registration Request Body"
// @Success 200 {object} OrderRegistrationResponse "Successful order registration"
// @Failure 400 {object} constants.HttpError "Invalid request body or parameters"
// @Router /register-order [post]
func OrderRegistration(c *gin.Context) {
	var orderRegistrationRequest OrderRegistrationRequest

	if err := c.BindJSON(&orderRegistrationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}

	orderRegistrationRequest.OrderNumber = utils.GenerateOrderNumber(1, 32)
	orderRegistrationRequest.Currency = "934"
	orderRegistrationRequest.ReturnURL = "/"

	urlParams := utils.StructToURLParams(orderRegistrationRequest)

	resp, err := makeOrderRegistration(orderRegistrationRequest.ApiClient, urlParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = createNewOrder(orderRegistrationRequest, resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
