package main

import (
	"strings"

	"github.com/yterajima/go-sitemap"
)

var root string

func GetDirectory(url string) map[string][]string {
	root = strings.Replace(url, "sitemap.xml", "", 1)

	sitemap, err := sitemap.Get(url, nil)
	if err != nil {
		panic(err)
	}

	return BuildDirectory(sitemap)
}

func BuildDirectory(sitemap sitemap.Sitemap) map[string][]string {
	directory := make(map[string][]string)

	for _, url := range sitemap.URL {
		path := strings.TrimPrefix(url.Loc, root)
		path = strings.Split(path, "/")[0]

		directory[path] = nil
	}

	for _, url := range sitemap.URL {
		path := strings.TrimPrefix(url.Loc, root)
		parts := strings.Split(path, "/")
		if len(parts) < 2 {
			continue
		}

		path1 := parts[0]
		path2 := parts[1]

		if len(parts) > 2 {
			for i := 1; i < len(parts); i++ {
				path2 = path2 + "/" + parts[i]
			}
		}

		if path2 != "" {
			directory[path1] = append(directory[path1], path2)
		}
	}
	return directory
}
