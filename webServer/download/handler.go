package download

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	req, e := http.Get("http://" + os.Getenv("API_SERVER") + "/objects/" + url.PathEscape(r.URL.Query()["name"][0]) + "?version=" + r.URL.Query()["version"][0])
	if e != nil {
		log.Println(e)
		return
	}
	w.Header().Set("content-disposition", "attachment;filename="+r.URL.Query()["name"][0])
	io.Copy(w, req.Body)
}
