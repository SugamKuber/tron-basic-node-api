package routes

import (
	"github.com/gin-gonic/gin"
	"tron-api-server/internal/api"
	"tron-api-server/internal/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/h", api.CheckHealth)
	router.GET("/b/:address", middleware.CheckValidAddress(), api.HandlerGetBalance)
	router.GET("/txn/:address", middleware.CheckValidAddress(), api.HandlerGetTransactions)
}
