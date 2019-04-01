package parser

import (
	"bufio"
	"os"
	"strings"
)

///Parse returns a string, which represents the filepath in which it created (HTML file).
func Parse(filePath string) []string {
	result := loadTextFile(filePath)
	splitResult := strings.Split(result, "\r\n")
	return splitResult
}

//Loads file in, returns result as string
func loadTextFile(filePath string) string {
	data := ""
	file, error := os.Open(filePath)
	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data += scanner.Text()
	}

	return string(data[:])
}
