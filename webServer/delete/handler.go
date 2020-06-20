package delete

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query()["name"][0]
	req, e := http.NewRequest("DELETE", "http://"+os.Getenv("API_SERVER")+"/objects/"+name, nil)
	if e != nil {
		log.Println(e)
		return
	}
	client := http.Client{}
	_, e = client.Do(req)
	if e != nil {
		log.Println(e)
		fmt.Println(e)
		return
	}
	log.Println("Delele file", name)
	t, e := template.ParseFiles("view/delete.html")
	if e != nil {
		log.Println(e)
		return
	}
	t.Execute(w, nil)
}
