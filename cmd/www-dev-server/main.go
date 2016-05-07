package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

var staticPrefixes = []string{
	"/assets/",
	"/app/",
	"/bower_components/",
}

func main() {
	root := os.Getenv("WWW_ROOT")
	static := http.FileServer(http.Dir(fmt.Sprintf("%s", root)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, prefix := range staticPrefixes {
			if strings.HasPrefix(r.URL.Path, prefix) {
				static.ServeHTTP(w, r)
				return
			}
		}

		bytes, _ := ioutil.ReadFile(fmt.Sprintf("%s/index.html", root))
		w.Write(bytes)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil); err != nil {
		fmt.Println(err)
	}
}
