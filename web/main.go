package main

import (
	"html/template"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	htmlFilePath := "src/template/login.html"
	log.Println("load html", htmlFilePath)
	t, err := template.ParseFiles(htmlFilePath)
	if err != nil {
		log.Fatalln("error", err)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatalln("error", err)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	htmlFilePath := "src/template/logout.html"
	log.Println("load html", htmlFilePath)
	t, err := template.ParseFiles(htmlFilePath)
	if err != nil {
		log.Fatalln("error", err)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatalln("error", err)
	}
}

func main() {
	log.Println("start serve.")
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.ListenAndServe(":8000", nil)
}
