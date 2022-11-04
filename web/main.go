package main

import (
	"go_login/web/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	file := logging.SetLogger()
	logger := logging.GetLogger()
	logger.LogLevel = logging.INFO
	logger.Info("start serve.")

	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"login.html",
			nil,
		)
	})
	r.Run(":8000")
	// http.HandleFunc("/login", controller.LoginHandler)
	// http.HandleFunc("/logout", controller.LogoutHandler)
	// http.ListenAndServe(":8000", nil)

	defer file.Close()
}
