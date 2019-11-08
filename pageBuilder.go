package main

import (
	"io/ioutil"
	"fmt"
)

//todo: move this to utilities
func readFile(filePath string) []byte {
	contents, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
	}

	return contents
}

func buildPage(pagePath string) string{
	layout := readFile("/content/layout.html")
	post := readFile(pagePath)
	return replaceText(string(layout), string(post), "{BLOGPOST}")
}