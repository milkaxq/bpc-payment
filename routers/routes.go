package routers

import (
	"github.com/gin-gonic/gin"
	docs "github.com/milkaxq/bpcpayment/docs"
	"github.com/milkaxq/bpcpayment/handlers/checkstatus"
	"github.com/milkaxq/bpcpayment/handlers/confirpayment"
	"github.com/milkaxq/bpcpayment/handlers/orderregistration"
	"github.com/milkaxq/bpcpayment/handlers/refund"
	"github.com/milkaxq/bpcpayment/handlers/resendpassword"
	"github.com/milkaxq/bpcpayment/handlers/submitcard"
	"github.com/milkaxq/bpcpayment/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	docs.SwaggerInfo.Title = "BPC payment service"
	docs.SwaggerInfo.Description = "This is simple bpc payment service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register-order", orderregistration.OrderRegistration)
		v1.POST("/submit-card", submitcard.SubmitCard)
		v1.POST("/confirm-payment", confirpayment.ConfirmPayment)
		v1.POST("/check-status", checkstatus.CheckStatus)
		v1.POST("/resend-password", resendpassword.ResendPassword)
		v1.POST("/refund", refund.Refund)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
