package journal

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.EscapedPath(), "/")[2]
	filename = os.Getenv("LOG_PREFIX") + filename + ".log"
	f, e := os.Open(filename)
	defer f.Close()
	if e != nil {
		log.Println(e)
		return
	}
	io.Copy(w, f)
}
