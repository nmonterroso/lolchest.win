package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), http.FileServer(http.Dir(os.Getenv("WWW_ROOT"))))
}
