package utilities

import (
	"strings"
	"io/ioutil"
	"fmt"
)

func ReadFile(filePath string) []byte {
	contents, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
	}

	return contents
}

func ReplaceText(text string, replacementValue string, valueToReplace string) string{
	return strings.Replace(text, replacementValue, valueToReplace, -1)
}