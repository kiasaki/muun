package main

import (
	"log"
)

func main() {
	SetConfig(ParseFlag())

	if Cfg().Serve {
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
