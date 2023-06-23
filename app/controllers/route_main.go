package controllers

import (
	"net/http"
)

// この関数の引数はパターン
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "hello", "layout", "public_navbar", "top")
}
