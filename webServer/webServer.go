package main

import (
	"log"
	"net/http"
	"os"

	"./delete"
	"./download"
	"./es"
	"./index"
	"./journal"
	"./mq"
	"./upload"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index.Handler)
	http.HandleFunc("/filelist", index.Handler)
	http.HandleFunc("/upload", upload.Handler)
	http.HandleFunc("/download", download.Handler)
	http.HandleFunc("/delete", delete.Handler)
	http.HandleFunc("/mq", mq.Handler)
	http.HandleFunc("/es", es.Handler)
	http.HandleFunc("/journal", journal.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
