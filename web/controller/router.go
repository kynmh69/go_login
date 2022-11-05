package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/login", getLogin)
	r.POST("/login", postLogin)
	r.GET("/logout", getLogout)
	r.GET("/signup", getSignUp)
	r.POST("/signup", postSignUp)

	return r
}
