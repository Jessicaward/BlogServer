package main

import (
	// Private
	"BlogServer/utilities"

	//Public
	"gopkg.in/russross/blackfriday.v2"
)

type BlogPost struct {
	Title string
	Views int
	Body  string
}

func GetPost(postName string) BlogPost {
	//posts will be in XML format
	//containing metadata on post, as well as markdown (actual post)
	post := utilities.ReadFile(postName + ".html")

	//todo: introduce XML parser

	return BlogPost{
		Body:  string(post),
		Title: "test",
	}
}

func generateHtmlFromMarkdown(markdown []byte) []byte {
	return blackfriday.Run(markdown)
}
