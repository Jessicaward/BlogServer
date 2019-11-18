package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandleRoute() {
	http.HandleFunc("/", loadHomePage)
	http.HandleFunc("/album/", loadRandomAlbum)
	http.HandleFunc("/projects/", loadProjects)
	http.HandleFunc("/social/", loadSocial)

	http.HandleFunc("/blog/", loadBlogPost)

	//static CSS and JS files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":80", nil)
}

func loadHomePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("content/Layout.html", "content/Home.html")
	t.Execute(w, "layout")
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

func loadSocial(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("content/Layout.html", "content/Social.html")
	t.Execute(w, "layout")
}

func loadBlogPost(w http.ResponseWriter, r *http.Request) {
	//todo: need a way to get post name from request
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, _ := template.ParseFiles("content/Layout.html", "content/Blog.html")
	post := GetPost("test-post")
	fmt.Println("loaded: " + post.Title)
	t.ExecuteTemplate(w, "layout", post)
}
