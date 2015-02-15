package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/kiasaki/batbelt/mst"
	"github.com/russross/blackfriday"
)

type PageType string

const (
	PtPage PageType = "page"
	PtPost          = "post"
)

type Page struct {
	Title    string
	Filename string
	Type     PageType
	Contents string
}

func NewPage(filename string, pt PageType) Page {
	var title string
	ext := filepath.Ext(filename)

	contents := string(mustReadFile(filename))
	// HTML is skipped, it'll be rendered using html/template later
	if ext == ".md" {
		// Use markdown h1 tag as page title
		contentLines := strings.Split(string(contents), "\n")
		title = strings.Replace(contentLines[0], "# ", "", -1)
		contents = strings.Join(contentLines[1:], "\n")
		contents = string(blackfriday.MarkdownCommon([]byte(contents)))

		// parse with blackfriday
		contents = fmt.Sprintf(`
			{{define "title"}}%s{{end}}
			{{define "contents"}}%s{{end}}
		`, Cfg().SiteTitle+" "+title, contents)
	}

	return Page{
		Title:    title,
		Filename: filename,
		Type:     pt,
		Contents: contents,
	}
}

func (p Page) OutFilename() string {
	return strings.Replace(filepath.Base(p.Filename), ".md", ".html", -1)
}

func (p Page) Link() string {
	return "/" + strings.Replace(p.OutFilename(), "-", "/", 3)
}

func (p Page) IsPost() bool {
	return p.Type == PtPost
}

func (p Page) IsMarkdown() bool {
	return filepath.Ext(p.Filename) == ".md"
}

func (p Page) DateFromFilename() time.Time {
	fn := p.OutFilename()
	y := mst.MustInt(strconv.Atoi(fn[:4]))
	m := time.Month(mst.MustInt(strconv.Atoi(fn[5:7])))
	d := mst.MustInt(strconv.Atoi(fn[8:10]))
	loc, err := time.LoadLocation("UTC")
	mst.MustNotErr(err)
	return time.Date(y, m, d, 0, 0, 0, 0, loc)
}

func (p Page) DateFormatted() string {
	d := p.DateFromFilename()
	return fmt.Sprintf("%d %s %d", d.Day(), d.Month().String()[:3], d.Year())
}

func (p Page) WriteToBuildDir(bi BuildInfo) {
	t, err := template.New("contents").Parse(p.Contents)
	mst.MustNotErr(err)

	t, err = t.New("layout").ParseFiles(Cfg().LayoutFullpath())
	mst.MustNotErr(err)

	var doc bytes.Buffer
	err = t.ExecuteTemplate(&doc, "layout", bi)
	mst.MustNotErr(err)

	writeToDisk(p.Link(), doc.Bytes())
}

func writeToDisk(destinationFilename string, contents []byte) {
	buildDir := mst.MustString(filepath.Abs(Cfg().BuildDir))
	finalFilePath := filepath.Join(buildDir, destinationFilename)

	// Ensure build dir exists
	err := os.MkdirAll(filepath.Dir(finalFilePath), 0755)
	mst.MustNotErr(err)

	// Write file
	err = ioutil.WriteFile(finalFilePath, contents, 0644)
	mst.MustNotErr(err)
}
