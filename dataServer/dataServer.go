package main

import (
	"log"
	"net/http"
	"os"

	"./heartbeat"
	"./locate"
	"./objects"
	"./temp"
	//"./versions"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	//go versions.StartDeleteOldMetadata()
	//go versions.StartDeleteOrphanObject()
	//go versions.StartObjectScanner()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
