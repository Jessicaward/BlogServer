package main

import (
	"BlogServer/parser"
	"fmt"
)

func main() {
	filePath := "content/blog.jess"
	fmt.Println(filePath)
	result := parser.Parse(filePath)
	for x := 0; x < len(result); x++ {
		fmt.Println(result)
	}
}
