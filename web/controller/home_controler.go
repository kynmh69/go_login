package controller

import (
	"go_login/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetHome(ctx *gin.Context) {
	var user *model.User
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")

	userId := model.GetSession(ctx, cookieKey)

	if userId != nil {
		user, _ = model.GetOneUser(userId.(string))
	}
	ctx.HTML(
		http.StatusOK,
		"home.html",
		gin.H{
			"user": user,
		},
	)
}
