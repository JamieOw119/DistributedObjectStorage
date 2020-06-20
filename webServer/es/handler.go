package es

import (
	"html/template"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("view/es.html")
	if e != nil {
		log.Println(e)
		return
	}
	t.Execute(w, nil)
}
