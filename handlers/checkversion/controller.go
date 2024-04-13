package checkversion

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/constants"
)

// @Summary Check EPG Versions
// @Description This API endpoint checks the versions of E-commerce Payment Gateway data for two URLs constructed from a provided base URL.
// @Tags epg
// @Accept json
// @Produce json
// @Param request body CheckVersionRequest true "Request containing the base URL"
// @Success 200 {object} EPGVersionsResponse  "success response containing versions for both URLs"
// @Failure 400 {object} constants.HttpError "on invalid request body"
// @Failure 500 {object} constants.HttpError "on check version function error"
// @Router /check-version [post]
func CheckVersion(c *gin.Context) {
	var request CheckVersionRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequestBody + err.Error()})
		return
	}

	versions, err := checkEPGVersions(request.BaseURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, versions)
}
