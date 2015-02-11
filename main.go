package main

import (
	"flag"
	"log"
)

var (
	fServe = flag.Bool("serve", false, "If provided an http server will be started")
	fPort  = flag.String("port", "8080", "Port to listen on if a server is started")
)

func main() {
	flag.Parse()

	if *fServe {
		cmdServe()
	} else {
		cmdBuild()
	}
}

func mustString(s string, err error) string {
	if err != nil {
		log.Panic(err)
	}
	return s
}

func mustStringArray(s []string, err error) []string {
	if err != nil {
		log.Panic(err)
	}
	return s
}
