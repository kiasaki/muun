package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/kiasaki/batbelt/mst"
)

func main() {
	SetConfig(ParseFlag())

	if Cfg().Serve {
		cmdServe()
	} else {
		cmdBuild()
	}
}

func mustReadFile(filePath string) []byte {
	bytesRead, err := ioutil.ReadFile(filePath)
	mst.MustNotErr(err)
	return bytesRead
}

func globFilenamesForDir(dir string) []string {
	path := mst.MustString(filepath.Abs(dir))

	pages := []string{}
	pages = append(pages, mst.MustStringArray(filepath.Glob(filepath.Join(path, "*.md")))...)
	pages = append(pages, mst.MustStringArray(filepath.Glob(filepath.Join(path, "*.html")))...)

	return pages
}
