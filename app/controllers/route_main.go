package controllers

import (
	"log"
	"net/http"
	"text/template"
)

// この関数の引数はパターン
func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, "hello")
}
