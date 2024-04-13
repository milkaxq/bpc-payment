package checkversion

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @Summary Check EPG Versions
// @Description This API endpoint checks the versions of E-commerce Payment Gateway data for two URLs constructed from a provided base URL in *ENV*.
// @Tags epg
// @Produce json
// @Success 200 {object} EPGVersionsResponse  "success response containing versions for both URLs"
// @Failure 400 {object} constants.HttpError "on invalid request body"
// @Failure 500 {object} EPGVersionsResponse "on check version function error"
// @Router /check-version [get]
func CheckVersion(c *gin.Context) {
	var request CheckVersionRequest

	request.BaseURL = os.Getenv("BASE_URL")
	versions := checkEPGVersions(request.BaseURL)
	if versions.Error1 != "" && versions.Error2 != "" {
		c.JSON(http.StatusInternalServerError, versions)
		return
	}

	c.JSON(http.StatusOK, versions)
}
