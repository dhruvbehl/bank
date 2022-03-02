package main

import (
	"github.com/dhruvbehl/bank/app"
	"github.com/dhruvbehl/bank/logger"
)

func main() {
	logPath := ""
	logger.Init(logPath)
	logger.Info("Starting application")
	app.Start()
}