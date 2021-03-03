package main

import (
	"net/http"
	"embed"
)

//go:embed public/*

var content embed.FS

func main() {
	http.HandleFunc("/", http.FileServer(content))
	http.ListenAndServe(":1789", nil)
}
