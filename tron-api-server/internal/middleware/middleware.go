package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CheckValidAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Param("address")
		if address == "" {
			c.JSON(400, gin.H{"error": "Address parameter is required"})
			c.Abort()
			return
		}

		fmt.Printf("Address: %s", address)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"address": address,
			"visible": true,
		})

		resp, err := http.Post(os.Getenv("LOCAL_NODE_API_URL"), "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Printf("Error making POST request: %v", err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v", err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Printf("Error unmarshalling JSON response: %v", err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		if valid, ok := result["result"].(bool); !ok || !valid {
			c.JSON(400, gin.H{"error": "Invalid address"})
			c.Abort()
			return
		}

		c.Next()
	}
}
