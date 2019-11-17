package main

import (
	// Private
	"BlogServer/utilities"
	"strings"

	//Public
	"gopkg.in/russross/blackfriday.v2"
)

type BlogPost struct {
	Title string
	Body  string
}

func GetPost(postName string) BlogPost {
	post := utilities.ReadFile("posts/" + postName + ".md")
	postBody := generateHtmlFromMarkdown(post)
	postTitle := getTitleFromMarkdown(post)

	return BlogPost{
		Body:  string(postBody),
		Title: postTitle,
	}
}

func generateHtmlFromMarkdown(markdown []byte) []byte {
	return blackfriday.Run(markdown)
}

func getTitleFromMarkdown(markdown []byte) string {
	//todo: this returns the title + the remainder of the post
	return strings.Trim(string(markdown), " #")
}
