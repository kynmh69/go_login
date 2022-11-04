package main

import (
	"go_login/web/controller"
	"go_login/web/logging"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	file := logging.SetLogger()
	logger := logging.GetLogger()
	logger.LogLevel = logging.INFO
	logger.Info("start serve.")
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)
	http.ListenAndServe(":8000", nil)

	quit := make(chan os.Signal, 10)
	signal.Notify(quit, os.Interrupt)

	<-quit

	logger.Info("finish serve.")

	defer file.Close()
}
