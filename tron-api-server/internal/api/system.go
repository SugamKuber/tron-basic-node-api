package api

import (
	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
    c.JSON(200, gin.H{"status":true })
}

