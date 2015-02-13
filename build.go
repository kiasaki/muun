package main

import (
	"log"
)

type BuildInfo struct {
	Pages       []Page
	CurrentPage Page
}

func (b BuildInfo) Posts() []Page {
	posts := []Page{}
	// Reverse the order as this is used on the index page to list recent posts
	for i := len(b.Pages) - 1; i > 0; i-- {
		if b.Pages[i].Type == PtPost {
			posts = append(posts, b.Pages[i])
		}
	}
	return posts
}

func cmdBuild() {
	pages := []Page{}

	for _, pType := range []PageType{PtPost, PtPage} {
		pageFiles := globFilenamesForDir(Cfg().DirForPageType(pType))
		log.Printf("Building %d %ss\n", len(pageFiles), pType)

		for _, page := range pageFiles {
			pages = append(pages, NewPage(page, pType))
		}
	}

	for _, page := range pages {
		buildInfo := BuildInfo{pages, page}
		page.WriteToBuildDir(buildInfo)
	}
}
