package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"tron-api-server/internal/service"
)

func HandlerGetBalance(c *gin.Context) {
	address := c.Param("address")

	result, err := service.FetchBalance(address)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func HandlerGetTransactions(c *gin.Context) {
	address := c.Param("address")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid limit"})
		return
	}

	start, err := strconv.Atoi(c.DefaultQuery("start", "0"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid start"})
		return
	}

	transactions, err := service.FetchTransactions(address, limit, start)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, transactions)
}
