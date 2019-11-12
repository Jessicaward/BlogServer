package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandleRoute() {
	http.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my blog")
	})

	//static CSS and JS files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/album/", loadRandomAlbum)
	http.HandleFunc("/projects/", loadProjects)

	http.ListenAndServe(":80", nil)
}

func loadRandomAlbum(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("content/Layout.html", "content/RandomAlbum.html")
	album := loadAlbum()
	fmt.Println("loaded: " + album.Name)
	t.ExecuteTemplate(w, "layout", album)
}

func loadProjects(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("content/Layout.html", "content/Projects.html")
	t.Execute(w, "layout")
}
