package controller

import (
	"go_login/logging"
	"go_login/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getSignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", nil)
}

func postSignUp(ctx *gin.Context) {
	logger := logging.GetLogger()

	id := ctx.PostForm("user_id")
	pw := ctx.PostForm("password")

	logger.Info("request to create user:", id)

	_, err := model.SignUp(id, pw)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/signup")
		return
	}
	ctx.Redirect(http.StatusFound, "/login")
}

func getLogin(ctx *gin.Context) {
	logger := logging.GetLogger()
	logger.Debug("call get Login")
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func postLogin(ctx *gin.Context) {
	logger := logging.GetLogger()

	id := ctx.PostForm("user_id")
	pw := ctx.PostForm("password")

	user, err := model.Login(id, pw)

	logger.Info("Login: ", id)

	if err != nil {
		logger.Warn("Login failure:", id)
		ctx.Redirect(http.StatusFound, "/login")
		return
	}
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	model.NewSession(ctx, cookieKey, user.UserId)

	ctx.HTML(http.StatusOK, "home.html", gin.H{"user": user})
}
