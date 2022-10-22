package middleware

import (
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "Token is required"})
			context.Abort()
			return
		}

		checkToken := strings.Fields(tokenString)
		if len(checkToken) != 2 {
			context.JSON(401, gin.H{"error": "Token is invalid"})
			context.Abort()
			return
		}

		err := helpers.ValidateToken(strings.Split(tokenString, "Bearer ")[1])
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
