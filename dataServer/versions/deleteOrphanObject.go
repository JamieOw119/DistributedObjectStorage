package versions

import (
	"lib/es"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func StartDeleteOrphanObject() {
	for {
		old := "%2F"
		new := "/"
		files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/object/*")
		for i := range files {
			hash := strings.Split(filepath.Base(files[i]), ".")[0]
			eshash := strings.ReplaceAll(hash, old, new)
			hashInMetadata, e := es.HasHash(eshash)
			if e != nil {
				log.Println(e)
				return
			}
			if !hashInMetadata {
				del(hash)
			}
		}
		time.Sleep(5 * time.Minute)
	}
}

func del(hash string) {
	log.Println("delete", hash)
	url := "http://" + os.Getenv("LOCATE_ADDRESS") + "/objects/" + hash
	request, _ := http.NewRequest("DELETE", url, nil)
	client := http.Client{}
	client.Do(request)
}
