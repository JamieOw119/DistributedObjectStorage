package objects

import (
	"fmt"
	"io"
	"lib/utils"
	"net/http"
	"net/url"

	"../locate"
)

func storeObject(r io.Reader, hash string, size int64) (int, error) {
	escapedhash := url.PathEscape(hash)
	if locate.Exist(escapedhash) {
		return http.StatusOK, nil
	}

	stream, e := putStream(escapedhash, size)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}

	reader := io.TeeReader(r, stream)
	d := utils.CalculateHash(reader)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	stream.Commit(true)
	return http.StatusOK, nil
}
