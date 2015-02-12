package main

import (
	"log"
)

type BuildInfo struct {
	Pages []Page
}

func (b BuildInfo) Posts() []Page {
	posts := []Page{}
	for _, page := range b.Pages {
		if page.Type == PtPost {
			posts = append(posts, page)
		}
	}
	return posts
}

func cmdBuild() {
	pages := []Page{}

	for _, pType := range []PageType{PtPost, PtPage} {
		pageFiles := globFilenamesForDir(Cfg().DirForPageType(pType))
		log.Printf("Building %d %ss\n", len(pages)+1, pType)

		for _, page := range pageFiles {
			pages = append(pages, NewPage(page, pType))
		}
	}

	buildInfo := BuildInfo{pages}
	for _, page := range pages {
		page.WriteToBuildDir(buildInfo)
	}
}
