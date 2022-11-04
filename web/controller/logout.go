package controller

import (
	"go_login/web/logging"
	"html/template"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	htmlFilePath := "template/logout.html"
	logger.Info("load html", htmlFilePath)
	t, err := template.ParseFiles(htmlFilePath)
	if err != nil {
		logger.Fatalln("error", err)
	}
	if err := t.Execute(w, nil); err != nil {
		logger.Fatalln("error", err)
	}
}
