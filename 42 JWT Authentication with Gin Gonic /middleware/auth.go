package middleware

import (
	"net/http"

	"github.com/fbdaf/go-jwt-gin/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No authorization header provided"})
			ctx.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("first_name", claims.FirstName)
		ctx.Set("last_name", claims.LastName)
		ctx.Set("uid", claims.Uid)
		ctx.Set("user_type", claims.UserType)

		if claims.UserType != "ADMIN" && claims.UserType != "USER" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User is not authorized"})
			ctx.Abort()
			return
		}

		if claims.UserType == "USER" && ctx.Request.URL.Path == "/api-1" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "User is not authorized"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
