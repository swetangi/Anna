package middleware

import (
	"anna/config"
	jwttoken "anna/jwt"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(appCtx *config.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwttoken.VerifyToken(tokenString, appCtx.JwtConfig.JwtSecretKey)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			fmt.Println("claims[email] ", claims["email"])
			ctx.Set("email", claims["email"])
			ctx.Next()
			v, ok := ctx.Get("email")
			fmt.Println("sss ", v, ok)
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	}
}
