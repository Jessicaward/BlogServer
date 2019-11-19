package main

import (
	"fmt"
)

func main() {
	var blog = GetPost("test-post")
	fmt.Println("title: " + blog.Metadata.Title)
	fmt.Println("created at: " + blog.Metadata.CreatedAt)
	fmt.Println("body: " + blog.Body)
	//fmt.Println("HTTP Server starting...")
	//HandleRoute()
}
