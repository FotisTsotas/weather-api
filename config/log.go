package config

import (
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func InitLogger() {
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal("Could not create logs directory: ", err)
	}
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Could not open log file: ", err)
	}

	logWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(logWriter)
	gin.DefaultWriter = logWriter
	gin.DefaultErrorWriter = logWriter

	slog.SetDefault(slog.New(slog.NewTextHandler(logWriter, nil)))
}
