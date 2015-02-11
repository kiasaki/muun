package main

import (
	"html/template"
)

type PageType string

const (
	PtPage PageType = "page"
	PtPost          = "post"
)

type Page struct {
	Title    string
	Contents template.HTML
}
