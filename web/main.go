package main

import (
	"go_login/web/controller"
	"go_login/web/logging"
	"net/http"
)

func main() {
	file := logging.SetLogger()
	logger := logging.GetLogger()
	logger.Debug("start serve.")
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)
	http.ListenAndServe(":8000", nil)

	defer file.Close()
}
