package main

import (
	"embed"
	"log"
	"net/http"
)

////go:embed index.html favicon.ico robots.txt *.css
var content embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(content)))
	log.Fatal(http.ListenAndServe(":11789", nil))
}
