package main

import (
	"go_login/controller"
	"go_login/logging"
	"go_login/model"
	"go_login/utils"
	"os"
)

func main() {
	file := logging.SetLogger()
	logger := logging.GetLogger()
	logger.LogLevel = logging.INFO
	logger.Info("start serve.")

	utils.LoadEnv("")

	logLevel := os.Getenv("APP_LOG_LEVEL")
	logger.SetLogLevel(logLevel)

	model.ConnectDb()

	r := controller.GetRouter()
	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port)

	defer file.Close()
}
