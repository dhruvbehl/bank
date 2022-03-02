package main

import (
	"os"

	"github.com/dhruvbehl/bank/app"
	"github.com/dhruvbehl/bank/domain"
	"github.com/dhruvbehl/bank/logger"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	logPath := os.Getenv("LOG_PATH")

	if len(host) == 0 || len(port) == 0 || len(dbUsername) == 0 || len(dbPassword) == 0 || len(dbHost) == 0 || len(dbPort) == 0 || len(dbName) == 0 {
		panic("Application needs HOST, PORT, DBUSERNAME, DBPASSWORD, DBHOST, DBPORT & DBNAME as environment variables")
	}

	envVar := domain.Environment{
		HOST:       host,
		PORT:       port,
		DBUSERNAME: dbUsername,
		DBPASSWORD: dbPassword,
		DBHOST:     dbHost,
		DBPORT:     dbPort,
		DBNAME:     dbName,
		LOGPATH:    logPath,
	}

	logger.Init(logPath)
	logger.Info("Starting application")
	app.Start(envVar)
}