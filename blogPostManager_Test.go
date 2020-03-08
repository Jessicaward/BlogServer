package main

import (
	"BlogServer/utilities"
	"strings"
	"testing"
)

func TestMarkdownParser(t *testing.T) {
	testFile := utilities.ReadFile("posts/test_post.md")

	//testFile is required, otherwise rest of test is invalid
	if testFile == nil {
		t.Errorf("Test post was not found")
	}

	html := generateHtmlFromMarkdown(testFile)

	if html == nil || len(html) == 0 {
		t.Errorf("HTML was not generated at all")
	}

	//test file includes a bold tag
	if !strings.Contains(string(html), "<strong>") {
		t.Fail()
	}
}
