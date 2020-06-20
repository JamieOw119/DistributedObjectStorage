package main

import (
	"log"
	"net/http"
	"os"

	"./heartbeat"
	"./journal"
	"./locate"
	"./objects"
	"./temp"
	"./versions"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	http.HandleFunc("/journal/", journal.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
