package main

import (
	"log"
	"net/http"
	"strconv"
)

func cmdServe() {
	log.Printf("Starting web server on port %d\n", Cfg().Port)
	panic(http.ListenAndServe(":"+strconv.Itoa(Cfg().Port), http.FileServer(http.Dir(Cfg().BuildDir))))
}

func extractPath(r *http.Request) string {
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}
	log.Println(path[len(path)-5:])
	if path[len(path)-5:] != ".html" {
		path = path + "/index.html"
	}
	return path
}
