package middlewares

import (
	"mygram/helpers"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "UNAUTHORIZED",
				"message": "token not found",
			})
			return
		}

		bearer := strings.HasPrefix(token, "Bearer")
		if !bearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "UNAUTHORIZED",
				"message": "no bearer",
			})
			return
		}

		tokenStr := strings.Split(token, "Bearer ")[1]
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "UNAUTHORIZED",
				"msg":   "token not found after bearer",
			})
			return
		}

		claims, err := helpers.VerifyToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "UNAUTHORIZED",
				"msg":   err.Error(),
			})
			return
		}

		var data = claims.(jwt.MapClaims)

		ctx.Set("email", data["email"])
		ctx.Set("id", data["id"])
		ctx.Next()
	}
}
