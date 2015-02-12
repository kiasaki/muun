package main

import (
	"log"
)

type BuildInfo struct {
	Pages []Page
}

func cmdBuild() {
	pages := []Page{}

	for _, pType := range []PageType{PtPost, PtPage} {
		pageFiles := globFilenamesForDir(Cfg().DirForPageType(pType))
		log.Printf("Building %d %ss\n", len(pages), pType)

		for _, page := range pageFiles {
			pages = append(pages, NewPage(page, pType))
		}
	}

	buildInfo := BuildInfo{pages}
	for _, page := range pages {
		page.WriteToBuildDir(buildInfo)
	}
}
