package controller

import (
	"go_login/web/logging"
	"net/http"
	"text/template"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	htmlFilePath := "template/login.html"
	logger.Println("load html", htmlFilePath)
	t, err := template.ParseFiles(htmlFilePath)
	if err != nil {
		logger.Fatalln("error", err)
	}
	if err := t.Execute(w, nil); err != nil {
		logger.Fatalln("error", err)
	}
}
