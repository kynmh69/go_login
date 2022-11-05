package main

import (
	"go_login/logging"
	"go_login/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	file := logging.SetLogger()
	logger := logging.GetLogger()
	logger.LogLevel = logging.DEBUG
	logger.Info("start serve.")

	utils.LoadEnv()

	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"login.html",
			nil,
		)
	})
	r.GET("/logout", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"logout.html",
			nil,
		)
	})

	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port)
	// http.HandleFunc("/login", controller.LoginHandler)
	// http.HandleFunc("/logout", controller.LogoutHandler)
	// http.ListenAndServe(":8000", nil)

	defer file.Close()
}
