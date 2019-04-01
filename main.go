package main

import (
	"fmt"
	"BlogServer/parser"
)

func main(){
	result := Parse("BlogServer/Content/Blog.jess")
	for x := 0; x < len(result); x++{
		fmt.Println(result)
	}
}