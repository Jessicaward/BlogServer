package main

import (
	"BlogServer/utilities"
)

func BuildPage(pagePath string) string{
	layout := utilities.ReadFile("content/layout.html")
	post := utilities.ReadFile(pagePath)
	return utilities.ReplaceText(string(layout), string(post), "{text}")
}