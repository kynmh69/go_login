package controller

import (
	"go_login/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", nil)
}

func postSignUp(ctx *gin.Context) {
	id := ctx.PostForm("user_id")
	pw := ctx.PostForm("password")
	user, err := model.SignUp(id, pw)
	if err != nil {
		ctx.Redirect(301, "/login")
		return
	}
	ctx.HTML(http.StatusOK, "home.html", gin.H{"user": user})
}

func getLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func postLogin(ctx *gin.Context) {
	id := ctx.PostForm("user_id")
	pw := ctx.PostForm("password")

	user, err := model.Login(id, pw)
	if err != nil {
		ctx.Redirect(301, "/login")
		return
	}
	ctx.HTML(http.StatusOK, "home.html", gin.H{"user": user})
}
