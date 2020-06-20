package versions

import (
	"lib/es"
	"log"
	"time"
)

const MIN_VERSION_COUNT = 5

func StartDeleteOldMetadata() {
	for {
		buckets, e := es.SearchVersionStatus(MIN_VERSION_COUNT + 1)
		if e != nil {
			log.Println(e)
			return
		}
		for i := range buckets {
			bucket := buckets[i]
			for v := 0; v < bucket.Doc_count-MIN_VERSION_COUNT; v++ {
				println(bucket.Key, v+int(bucket.Min_version.Value))
				es.DelMetadata(bucket.Key, v+int(bucket.Min_version.Value))
			}
		}
		time.Sleep(5 * time.Minute)
	}
}
