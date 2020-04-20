package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)


type APTHandler struct{
	sourceBase map[string]string
	targetBase string
}

func (f APTHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Printf("request url %s", path)
	dist := path[:strings.Index(path, "/")]
	log.Printf("distribution %s", dist)
	source, ok := f.sourceBase[dist]
	if !ok {
		w.WriteHeader(400)
	} else {
		fileUrl := fmt.Sprintf("%s/%s", source, path)
		log.Printf("downloading %s", fileUrl)
		resp, err := http.Get(fileUrl)
		if err != nil {
			w.WriteHeader(500)
			log.Printf("fetch from %s failed %v", fileUrl, err)
			return
		}
		defer resp.Body.Close()

		filepath := fmt.Sprintf("%s%s", f.targetBase, path)
		parent := filepath[:strings.LastIndex(filepath, "/")]
		if _, err := os.Stat(parent); os.IsNotExist(err) {
			const PermDir = 0755
			err = os.MkdirAll(parent, PermDir)
			if err != nil {
				w.WriteHeader(500)
				log.Printf("create dir %s failed %v", parent, err)
				return
			}
		}
		file, err := os.Create(filepath)
		if err != nil {
			w.WriteHeader(500)
			log.Printf("write to %s failed %v", filepath, err)
			return
		}
		defer file.Close()
		out := io.MultiWriter(file, w)
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			w.WriteHeader(500)
			log.Printf("write respone stream failed %v", err)
		}
	}

}
