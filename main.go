package main

import (
	"fmt"
	"BlogServer/server"
)

func main() {
	fmt.Println("Blog server starting...")
	server.HandleRoute()
}