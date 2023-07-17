package main

import (
	"github.com/yterajima/go-sitemap"
)

type Directory struct {
	Folders []Folder
}

type Folder struct {
	FolderName string
	Folders    []Folder
	Files      []string
}

var root string

func GetSitemap(url string) {
	root = url

	sitemap, err := sitemap.Get(url, nil)
	if err != nil {
		panic(err)
	}

	BuildDirectory(sitemap)
}

func BuildDirectory(sitemap sitemap.Sitemap) {
}
