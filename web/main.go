package main

import (
	"go_login/controller"
	"go_login/logging"
	"go_login/utils"
	"log"
	"os"
)

var (
	file   *os.File
	logger *logging.Logger
)

func init() {
	log.Println("init main.")
	file = logging.SetLogger()
	logger = logging.GetLogger()
	logger.LogLevel = logging.INFO

	utils.LoadEnv("")
}

func main() {

	logger.Info("start serve.")

	logLevel := os.Getenv("APP_LOG_LEVEL")
	logger.SetLogLevel(logLevel)

	r := controller.GetRouter()
	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port)

	defer file.Close()
}
