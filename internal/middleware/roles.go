package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Requireadmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := GetRole(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
		if !strings.EqualFold(role, "admin") {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "this route can only access by admin",
			})
		}
		c.Next()
	}
}
