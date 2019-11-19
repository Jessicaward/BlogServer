package main

import (
	// Private
	"BlogServer/utilities"
	"database/sql"
	"log"

	//Public
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/russross/blackfriday.v2"
)

type BlogPost struct {
	Metadata Post
	Body     string
}

type Post struct {
	Title     string
	CreatedAt string
	Tags      []Tag
}

type Tag struct {
	Key  string
	Name string
}

func GetPost(postName string) BlogPost {
	post := utilities.ReadFile("posts/" + postName + ".md")
	postBody := generateHtmlFromMarkdown(post)
	metadata := new(Post)
	//open database connection
	database, err := sql.Open("sqlite3", "blog.db")

	if err == nil {
		log.Fatal(err)
	}

	rows, _ := database.Query("SELECT Title, CreatedAt FROM Post;")
	for rows.Next() {
		rows.Scan(metadata.Title, metadata.CreatedAt)
	}

	return BlogPost{
		Body:     string(postBody),
		Metadata: *metadata,
	}
}

func generateHtmlFromMarkdown(markdown []byte) []byte {
	return blackfriday.Run(markdown)
}
