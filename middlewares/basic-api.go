package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func BasicAPIKeyAuth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		apiKeys := ctx.GetHeader("X-API-KEY")
		if apiKeys == ""{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"Error":"API Key Missing",
			})
		}

		if apiKeys != "123"{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"Error":"Invalid API Key",
			})
		}
		ctx.Next()
	}
}