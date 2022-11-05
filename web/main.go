package main

import (
	"go_login/controller"
	"go_login/logging"
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

	r := controller.GetRouter()
	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port)
	// http.HandleFunc("/login", controller.LoginHandler)
	// http.HandleFunc("/logout", controller.LogoutHandler)
	// http.ListenAndServe(":8000", nil)

	defer file.Close()
}
