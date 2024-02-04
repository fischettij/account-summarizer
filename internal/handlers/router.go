package handlers

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine, summarizerHandler *SummarizerHandler) {
	router.POST("/summary/send", summarizerHandler.SendEmail)
}
