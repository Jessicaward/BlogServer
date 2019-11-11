package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func HandleRoute() *http.Server{
	server := &http.Server{Addr: ":80"}
	
	// http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprintf(w, "Welcome to my site")
	// })

	http.HandleFunc("/blog/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to my blog")
	})

	post := BuildPage("posts/testpost.html")
	//todo: test route, remove
	http.HandleFunc("/test/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, post)
	})

	//static CSS and JS files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//get random album
	http.HandleFunc("/album/", loadRandomAlbum)

	http.ListenAndServe(":80", nil)

    // returning reference so caller can call Shutdown()
	return server
}

func loadRandomAlbum (w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("content/Layout.html", "content/RandomAlbum.html")
	album := loadAlbum()
	fmt.Println(album.Name)
	t.ExecuteTemplate(w, "layout", album)
}

func loadAlbum() *Album {
	album := GetAlbumAtPosition(1)
	return &Album{Name: album.Name,
				  Artist: album.Artist,
				  Length: album.Length,
				  PositionOnLastFm: album.PositionOnLastFm}
}