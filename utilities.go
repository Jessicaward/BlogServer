package main

import (
	"strings"
)

func replaceText(text string, replacementValue string, valueToReplace string) string{
	return strings.Replace(text, replacementValue, valueToReplace, -1)
}