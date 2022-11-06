package controller

import (
	"go_login/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getLogout(ctx *gin.Context) {
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	model.DeleteSession(ctx, cookieKey)
	ctx.Redirect(http.StatusFound, "/login")
}
