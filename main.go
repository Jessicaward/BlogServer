package main

import (
	"fmt"
	"bufio"
	"os"
	"net/http"
	"context"
)

var state = "started"

func main() {
	fmt.Println("Blog server starting...")
	server := HandleRoute()

	//get commands
	for {
		state = getServerState()
		if (state != "stopped"){
			break
		}	
	}

	stopServer(server)
}

func getServerState() string{
	//get user command
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("BlogServer: ")
	command, _ := reader.ReadString('\n')

	switch command {
	case "stop":
	case "quit":
		return "stopped"
	}

	return "running"
}

func stopServer(server *http.Server) {
	if err := server.Shutdown(context.TODO()); err != nil {
        panic(err)
    }
}