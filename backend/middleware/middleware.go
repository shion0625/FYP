package middleware

import (
	"net/http"
	token "gihtub.com/SherzodAbdullajonov/ecommerce-yt/tokens"
	"github.com/gin-gonic/gin"
)


func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ClientToken := ctx.Request.Header.Get("token")
		if ClientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no Authorizaton header"})
			ctx.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		if err != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("uid", claims.Uid)
		ctx.Next()
	}
}