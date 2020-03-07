package main

import (
	// Private
	"BlogServer/utilities"

	//Public
	"gopkg.in/russross/blackfriday.v2"
)

type BlogPost struct {
	Metadata Post
	Body     string
}

type Post struct {
	Title             string
	CreatedAt         string
	Tags              []Tag
	ReadTimeInMinutes int
}

type Tag struct {
	Key  string
	Name string
}

func GetPost(postName string) BlogPost {
	//this article may help: https://www.thepolyglotdeveloper.com/2017/04/using-sqlite-database-golang-application/
	post := utilities.ReadFile("posts/" + postName + ".md")
	postBody := generateHtmlFromMarkdown(post)
	metadata := new(Post)
	postWordCount := utilities.WordCount(string(postBody))
	metadata.ReadTimeInMinutes = utilities.CalculateReadTime(postWordCount)

	return BlogPost{
		Body:     string(postBody),
		Metadata: *metadata,
	}
}

func generateHtmlFromMarkdown(markdown []byte) []byte {
	return blackfriday.Run(markdown)
}
