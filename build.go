package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

func cmdBuild() {
	log.Println("Discovering pages")
	wd := mustString(os.Getwd())

	pages := []string{}
	pagesMdPath := filepath.Join(wd, "pages", "*.md")
	pagesHtmlPath := filepath.Join(wd, "pages", "*.html")
	pages = append(pages, mustStringArray(filepath.Glob(pagesMdPath))...)
	pages = append(pages, mustStringArray(filepath.Glob(pagesHtmlPath))...)

	log.Println("-----")
	for _, page := range pages {
		processFile(page, "page")
	}
	log.Println("-----")
}

func processFile(path string, category PageType) {
	var contents string

	base := filepath.Base(path)
	ext := filepath.Ext(path)

	log.Println(base)

	var rawFileContents []byte
	if bytesRead, err := ioutil.ReadFile(path); err != nil {
		log.Panic(err)
	} else {
		rawFileContents = bytesRead
	}

	if ext == ".md" {
		// parse with blackfriday
		contents = string(blackfriday.MarkdownCommon(rawFileContents))
	} else if ext == ".html" {
		// parse with html/template
		contents = string(rawFileContents)
	}

	// Wrap contents in layout
	page := Page{
		Cfg().SiteTitle,
		template.HTML(contents),
	}
	buildDir := mustString(filepath.Abs(Cfg().BuildDir))
	finalFilePath := filepath.Join(buildDir, strings.Replace(base, ".md", ".html", -1))
	finalContent := wrapContentsInLayout(page)

	// Ensure build dir exists
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		log.Panic(err)
	}

	// Write file
	if err := ioutil.WriteFile(finalFilePath, []byte(finalContent), 0755); err != nil {
		log.Panic(err)
	}
}

func wrapContentsInLayout(page Page) string {
	var doc bytes.Buffer

	t := template.New("layout")
	t, err := t.ParseFiles(mustString(filepath.Abs(Cfg().LayoutFile)))
	err = t.ExecuteTemplate(&doc, "layout", page)
	if err != nil {
		log.Panic(err)
	}

	return doc.String()
}
