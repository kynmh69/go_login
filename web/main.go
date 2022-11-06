package main

import (
	"go_login/controller"
	"go_login/logging"
	"go_login/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	file   *os.File
	logger *logging.Logger
)

func init() {
	log.Println("init main.")
	file = logging.SetLogger()
	logger = logging.GetLogger()
	logger.LogLevel = logging.DEBUG

	utils.LoadEnv("")

	gin.SetMode(gin.ReleaseMode)
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
