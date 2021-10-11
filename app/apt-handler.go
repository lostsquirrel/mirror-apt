package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type APTHandler struct {
	sourceBase map[string]string
	targetBase string
}

func (f APTHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Printf("request url %s", path)
	distIndex := strings.Index(path, "/")
	if distIndex == -1 {
		log.Printf("unknown address %s", path)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	dist := path[:distIndex]
	log.Printf("distribution %s", dist)
	source, ok := f.sourceBase[dist]
	if !ok {
		w.WriteHeader(400)
		log.Printf("%s not found", dist)
	} else {
		fileUrl := fmt.Sprintf("%s/%s", source, path[distIndex+1:])
		log.Printf("downloading %s", fileUrl)
		resp, err := http.Get(fileUrl)
		if err != nil {
			w.WriteHeader(500)
			log.Printf("fetch from %s failed %v", fileUrl, err)
			return
		}
		if resp.StatusCode > 299 {
			w.WriteHeader(resp.StatusCode)
			log.Printf("fetch from %s failed", fileUrl)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Printf("close failed %v", err)
			}
		}(resp.Body)

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			w.WriteHeader(500)
			log.Printf("write respone stream failed %v", err)
		}

	}

}
