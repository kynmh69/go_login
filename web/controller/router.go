package controller

import (
	"go_login/logging"
	"go_login/model"
	"log"
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
	logger := logging.GetLogger()
	logger.Debug("Check Login")
	return func(c *gin.Context) {
		cookieKey := os.Getenv("LOGIN_USER_ID_KEY")

		id := model.GetSession(c, cookieKey)

		log.Println("id: ", id)
		if id == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			c.Next()
		}

	}
}

func checkLogout() gin.HandlerFunc {
	logger := logging.GetLogger()
	logger.Debug("Check Logout")
	return func(c *gin.Context) {
		logger := logging.GetLogger()
		cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
		id := model.GetSession(c, cookieKey)
		logger.Debug("id:", id)
		if id != nil {
			c.Redirect(http.StatusFound, "/home")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
