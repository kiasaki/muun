package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	SetConfig(ParseFlag())

	if Cfg().Serve {
		cmdServe()
	} else {
		cmdBuild()
	}
}

func assertNotErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func mustInt(i int, err error) int {
	if err != nil {
		log.Panic(err)
	}
	return i
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

func mustReadFile(filePath string) []byte {
	bytesRead, err := ioutil.ReadFile(filePath)
	assertNotErr(err)
	return bytesRead
}

func globFilenamesForDir(dir string) []string {
	path := mustString(filepath.Abs(dir))

	pages := []string{}
	pages = append(pages, mustStringArray(filepath.Glob(filepath.Join(path, "*.md")))...)
	pages = append(pages, mustStringArray(filepath.Glob(filepath.Join(path, "*.html")))...)

	return pages
}
