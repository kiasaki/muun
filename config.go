package main

import (
	"flag"
	"path/filepath"
)

type Config struct {
	Serve      bool
	Port       int
	BuildDir   string
	LayoutFile string
	PagesDir   string
	PostsDir   string
	SiteTitle  string
}

var globalConfig Config

func Cfg() Config {
	return globalConfig
}

func SetConfig(c Config) {
	globalConfig = c
}

func ParseFlag() Config {
	var (
		fServe      = flag.Bool("serve", false, "If provided an http server will be started")
		fPort       = flag.Int("port", 8080, "Port to listen on if a server is started")
		fBuildDir   = flag.String("build-dir", "build", "Directory to output builded website to")
		fLayoutFile = flag.String("layout-file", "_layout.html", "HTML layout for all pages")
		fPagesDir   = flag.String("pages-dir", "pages", "Directory containing pages")
		fPostsDir   = flag.String("posts-dir", "posts", "Directory containing posts")
		fSiteTitle  = flag.String("site-title", "", "Site title, passed on to layout")
	)
	flag.Parse()
	c := Config{
		*fServe,
		*fPort,
		*fBuildDir,
		*fLayoutFile,
		*fPagesDir,
		*fPostsDir,
		*fSiteTitle,
	}
	return c
}

func (c Config) LayoutFullpath() string {
	return mustString(filepath.Abs(Cfg().LayoutFile))
}

func (c Config) DirForPageType(pt PageType) string {
	if pt == PtPage {
		return c.PagesDir
	} else {
		return c.PostsDir
	}
}
