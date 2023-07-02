package security

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	success  = "success"
	fail     = "error"
	apiToken = "1"
)

func TokenA() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token-A")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "API token required"})
			return
		}
		// добавляем в контекст ключ, чтобы в дальнейшем можно было получить его из контекста
		pID, err := strconv.Atoi(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
			return
		}

		c.Set("pID", pID)

		if token != apiToken {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
			return
		}
		c.Next()
	}
}
