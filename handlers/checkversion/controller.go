package checkversion

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Check EPG Versions
// @Description This API endpoint checks the versions of E-commerce Payment Gateway data for two URLs constructed from a provided base URL in *ENV*.
// @Tags epg
// @Produce json
// @Success 200 {object} EPGVersionsResponse "Success response containing versions for both URLs"
// @Failure 500 {object} EPGVersionsResponse "Check version function error"
// @Router /check-version [get]
func CheckVersion(c *gin.Context) {
	versions := checkEPGVersions()

	if versions.Error1 != "" && versions.Error2 != "" {
		c.JSON(http.StatusInternalServerError, versions)
		return
	}

	c.JSON(http.StatusOK, versions)
}
