package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/milkaxq/bpcpayment/handlers/checkversion"
)

func InitRoutes() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/check-version", checkversion.CheckVersion)
	}

	r.Run()
}
