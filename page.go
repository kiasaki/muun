package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

type PageType string

const (
	PtPage PageType = "page"
	PtPost          = "post"
)

type Page struct {
	Filename string
	Type     PageType
	Contents string
}

func NewPage(filename string, pt PageType) Page {
	ext := filepath.Ext(filename)

	contents := string(mustReadFile(filename))
	// HTML is skipped, it'll be rendered using html/template later
	if ext == ".md" {
		// Use markdown h1 tag as page title
		contentLines := strings.Split(string(contents), "\n")
		title := Cfg().SiteTitle + strings.Replace(contentLines[0], "# ", " ", -1)
		contents = string(blackfriday.MarkdownCommon([]byte(contents)))

		// parse with blackfriday
		contents = fmt.Sprintf(`
			{{define "title"}}%s{{end}}
			{{define "contents"}}%s{{end}}
		`, title, contents)
	}

	return Page{
		Filename: filename,
		Type:     pt,
		Contents: contents,
	}
}

func (p Page) WriteToBuildDir(bi BuildInfo) {
	t, err := template.New("contents").Parse(p.Contents)
	assertNotErr(err)

	t, err = t.New("layout").ParseFiles(Cfg().LayoutFullpath())
	assertNotErr(err)

	var doc bytes.Buffer
	err = t.ExecuteTemplate(&doc, "layout", bi)
	assertNotErr(err)

	base := filepath.Base(p.Filename)
	writeToDisk(base, doc.Bytes())
}

func writeToDisk(originalFilename string, contents []byte) {
	buildDir := mustString(filepath.Abs(Cfg().BuildDir))
	destinationFilename := strings.Replace(originalFilename, ".md", ".html", -1)
	finalFilePath := filepath.Join(buildDir, destinationFilename)

	// Ensure build dir exists
	err := os.MkdirAll(buildDir, 0755)
	assertNotErr(err)

	// Write file
	err = ioutil.WriteFile(finalFilePath, contents, 0644)
	assertNotErr(err)
}
