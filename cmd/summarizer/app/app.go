package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/fischettij/account-summarizer/internal/handlers"
)

func Start() {
	r := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	config, err := LoadConfig()
	if err != nil {
		logger.Fatal(err.Error())
	}

	handlers.ConfigureRoutes(r)

	err = r.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info(fmt.Sprintf("application started and listening on port %s", config.Port))

}
