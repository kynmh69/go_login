package controller

import (
	"html/template"
	"net/http"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	htmlFilePath := "template/logout.html"
	logger.Println("load html", htmlFilePath)
	t, err := template.ParseFiles(htmlFilePath)
	if err != nil {
		logger.Fatalln("error", err)
	}
	if err := t.Execute(w, nil); err != nil {
		logger.Fatalln("error", err)
	}
}
