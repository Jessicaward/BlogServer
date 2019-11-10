package main

import (
	"net/http"
	"fmt"
)

func HandleRoute(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to my site")
	})

	http.HandleFunc("/blog/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to my blog")
	})

	post := BuildPage("posts/testpost.html")
	//todo: test route, remove
	http.HandleFunc("/test/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, post)
	})

	http.ListenAndServe(":80", nil)
}