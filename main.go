package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed public
var content embed.FS

func main() {
	fsys, err := fs.Sub(content, "public")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(fsys)))
	log.Fatal(http.ListenAndServe(":11789", nil))
}
