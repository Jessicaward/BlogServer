package server

import (
	"net/http"
	"fmt"
)

func HandleRoute(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to my site")
	})

	http.ListenAndServe(":80", nil)
}