package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		_, extErr := c.Writer.Write([]byte("pong"))
		if extErr != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
	})
}
