package middlewares

import (
	"fmt"
	"hacktiv8-final-project/helpers"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Token Not Found",
			})
			return
		}

		bearer := strings.HasPrefix(token, "Bearer")

		if !bearer {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Bearer Not FOund",
			})
			return
		}
		tokenStr := strings.Split(token, "Bearer ")[1]

		if tokenStr == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Token STR",
			})
			return
		}

		claims, err := helpers.VerifyToken(tokenStr)

		// fmt.Println(tokenStr)

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		var data = claims.(jwt.MapClaims)
		userId := data["id"]
		ctx.Set("id", fmt.Sprint(userId))
		ctx.Set("email", data["id"])
		ctx.Next()
	}
}
