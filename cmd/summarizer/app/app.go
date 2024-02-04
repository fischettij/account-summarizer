package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/fischettij/account-summarizer/internal/fileprocessor"
	"github.com/fischettij/account-summarizer/internal/handlers"
	"github.com/fischettij/account-summarizer/pkg/summarizer"
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

	fileParser, err := fileprocessor.NewFileParser(logger, config.FilesDirectory)
	if err != nil {
		logger.Fatal(err.Error())
	}

	summarizerManager, err := summarizer.NewManager(fileParser)
	if err != nil {
		logger.Fatal(err.Error())
	}

	summarizerHandler, err := handlers.NewSummarizerHandler(summarizerManager)
	if err != nil {
		logger.Fatal(err.Error())
	}

	handlers.ConfigureRoutes(r, summarizerHandler)

	err = r.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info(fmt.Sprintf("application started and listening on port %s", config.Port))

}
