package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
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
	r.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"signup.html",
			nil,
		)
	})
	r.POST("/signup", func(ctx *gin.Context) {
		ctx.Redirect(
			http.StatusFound,
			"login",
		)
	})
	return r
}
