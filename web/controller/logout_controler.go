package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getLogout(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/login")
}
