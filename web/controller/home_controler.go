package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"home.html",
		nil,
	)
}
