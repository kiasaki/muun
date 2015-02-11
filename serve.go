package main

import (
	"log"
	"net/http"
	"strconv"
)

func cmdServe() {
	http.HandleFunc("/posts", handlePost)
	http.HandleFunc("/", handlePage)

	log.Printf("Starting web server on port %s\n", Cfg().Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Cfg().Port), nil))
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

func handlePage(w http.ResponseWriter, r *http.Request) {
}

func handlePost(w http.ResponseWriter, r *http.Request) {
}
