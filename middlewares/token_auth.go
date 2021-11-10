package middlewares

import (
	"payment-monitoring/auth"
	"payment-monitoring/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ResponseErrorCustom{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
