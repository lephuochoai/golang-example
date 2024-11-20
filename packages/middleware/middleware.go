package middleware

import (
	"example/web-service-gin/packages/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helper.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Access token required"})
			context.Abort()
			return
		}
		user, _ := helper.CurrentUser(context)
		context.Set("userId", user.ID)
		context.Next()
	}
}
