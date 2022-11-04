package main

import (
	"go_login/web/logging"
	"html/template"
	"log"
	"net/http"
)



func main() {
	file := logging.SetLogger(logger)
	logger.Println("start serve.")
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.ListenAndServe(":8000", nil)

	defer file.Close()
}
