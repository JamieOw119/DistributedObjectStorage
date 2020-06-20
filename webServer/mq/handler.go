package mq

import (
	"html/template"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("view/mq.html")
	if e != nil {
		log.Println(e)
		return
	}
	t.Execute(w, nil)
}
