package journal

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var urldict map[string]string = map[string]string{
	"webserver1":  "http://123.56.162.83:9007/journal/webServer",
	"webserver2":  "http://39.100.9.17:9008/journal/webServer",
	"webserver3":  "http://39.100.10.126:9009/journal/webServer",
	"apiserver1":  "http://123.56.162.83:9007/journal/apiServer",
	"apiserver2":  "http://39.100.9.17:9008/journal/apiServer",
	"apiserver3":  "http://39.100.10.126:9009/journal/apiServer",
	"dataserver1": "http://123.56.162.83:9007/journal/dataServer1",
	"dataserver2": "http://123.56.162.83:9007/journal/dataServer2",
	"dataserver3": "http://39.100.9.17:9008/journal/dataServer1",
	"dataserver4": "http://39.100.9.17:9008/journal/dataServer2",
	"dataserver5": "http://39.100.10.126:9009/journal/dataServer1",
	"dataserver6": "http://39.100.10.126:9009/journal/dataServer2",
}

func Handler(w http.ResponseWriter, r *http.Request) {
	servername := r.URL.Query()["target"][0]
	targeturl := urldict[servername]

	req, e := http.Get(targeturl)
	if e != nil {
		log.Println(e)
		return
	}

	s, e := ioutil.ReadAll(req.Body)
	if e != nil {
		log.Println(e)
		return
	}

	t, e := template.ParseFiles("view/journal.html")
	if e != nil {
		log.Println(e)
		return
	}
	t.Execute(w, template.HTML(s))
}
