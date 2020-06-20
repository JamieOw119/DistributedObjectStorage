package versions

import (
	"lib/es"
	"lib/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"../../apiServer/objects"
)

func StartObjectScanner() {
	for {
		old := "%25"
		new := "%"
		old2 := "%2F"
		new2 := "/"
		files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")
		for i := range files {
			hash := strings.Split(filepath.Base(files[i]), ".")[0]
			escape_hash := strings.ReplaceAll(hash, old, new)
			escape_hash2 := strings.ReplaceAll(escape_hash, old2, new2)
			verify(escape_hash2, hash)
		}
		time.Sleep(24 * time.Hour)
	}
}

func verify(escape_hash2 string, hash string) {
	log.Println("verify", escape_hash2)
	size, e := es.SearchHashSize(escape_hash2)
	if e != nil {
		log.Println(e)
		return
	}
	stream, e := objects.GetStream(hash, size)
	if e != nil {
		log.Println(e)
		return
	}
	d := utils.CalculateHash(stream)
	if d != escape_hash2 {
		log.Printf("object hash mismatch, calculated=%s, requested=%s", d, escape_hash2)
	}
	stream.Close()
}
