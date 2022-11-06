package controller

import (
	"go_login/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	loginCheckGroup := r.Group("/", checkLogin())
	{
		loginCheckGroup.GET("/home", GetHome)
		loginCheckGroup.GET("/logout", getLogout)

	}
	logoutCheckGroup := r.Group("/", checkLogout())
	{
		logoutCheckGroup.GET("/login", getLogin)
		logoutCheckGroup.POST("/login", postLogin)
		logoutCheckGroup.GET("/signup", getSignUp)
		logoutCheckGroup.POST("/signup", postSignUp)
	}

	return r
}

func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
		id := model.GetSession(c, cookieKey)
		if id == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func checkLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
		id := model.GetSession(c, cookieKey)
		if id != nil {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
