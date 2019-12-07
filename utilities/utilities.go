package utilities

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func ReadFile(filePath string) []byte {
	contents, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
	}

	return contents
}

func CalculateReadTime(numberOfWords int) int {
	return numberOfWords / 250
}

func WordCount(value string) int {
	return len(regexp.MustCompile(`[\S]+`).FindAllString(value, -1))
}

func ReplaceText(text string, replacementValue string, valueToReplace string) string {
	return strings.Replace(text, replacementValue, valueToReplace, -1)
}
