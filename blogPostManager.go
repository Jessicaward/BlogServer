package main

import (
	// Private
	"BlogServer/utilities"
	"database/sql"

	//Public
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
	database, _ := sql.Open("sqlite3", "./nraboy.db")

	rows, _ := database.Query("SELECT Title, CreatedAt FROM Post")
	var title string
	var createdAt string
	for rows.Next() {
		rows.Scan(&title, &createdAt)
	}

	metadata.Title = title
	metadata.CreatedAt = createdAt

	return BlogPost{
		Body:     string(postBody),
		Metadata: Post (
			Title: title,
			CreatedAt: createdAt,
		)
	}
}

func generateHtmlFromMarkdown(markdown []byte) []byte {
	return blackfriday.Run(markdown)
}
