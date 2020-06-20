package upload

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Sizer interface {
	Size() int64
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, header, e := r.FormFile("upload")
	if e != nil {
		log.Println(e)
		return
	}
	defer f.Close()
	h := sha256.New()
	io.Copy(h, f)
	d := base64.StdEncoding.EncodeToString(h.Sum(nil))
	f.Seek(0, 0)
	dat, _ := ioutil.ReadAll(f)
	req, e := http.NewRequest("PUT", "http://"+os.Getenv("API_SERVER")+"/objects/"+url.PathEscape(header.Filename), bytes.NewBuffer(dat))
	if e != nil {
		log.Println(e)
		return
	}
	req.Header.Set("digest", "SHA-256="+d)
	client := http.Client{}
	log.Println("uploading file", header.Filename, "hash", d, "size", f.(Sizer).Size())
	_, e = client.Do(req)
	if e != nil {
		log.Println(e)
		return
	}
	log.Println("uploaded")
	t, e := template.ParseFiles("view/upload.html")
	if e != nil {
		log.Println(e)
		return
	}
	t.Execute(w, nil)
}
