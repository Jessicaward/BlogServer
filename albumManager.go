package main

import (
	"math/rand"
	"time"
	//"github.com/shkh/lastfm-go/lastfm"
)

type Album struct {
	Name string
	PositionOnLastFm int
	Artist string
	Length string
}

//Constructor for Album struct
func NewAlbum(name string, artist string, positionOnLastFm int, length string) *Album {
	album := Album{Name: name, PositionOnLastFm: positionOnLastFm, Length: length, Artist: artist}
	return &album
}

// Get all albums from Last.fm
func GetAlbums() []*Album {
	//todo: this is just a test, replace with Last.fm API call
	albums := []*Album{NewAlbum("Norman Fucking Rockwell", "Lana Del Rey", 10, "60:19"),
					  NewAlbum("Kid A", "Radiohead", 5, "47:11"),
					  NewAlbum("The Fragile", "Nine Inch Nails", 8, "60:43"),
					  NewAlbum("Art Angels", "Grimes", 2, "49:43")}

	return albums
}

func GetAlbumAtPosition(position int) *Album {
	//api := lastfm.New (ApiKey, ApiSecret)
	//todo: again, just test data.. still need to implement the Last.fm API.
	return NewAlbum("Lateralus", "Tool", 1, "49:19")
}

//Get random album from Last.fm
func GetRandomAlbum() *Album {
	rand.Seed(time.Now().Unix())
	return GetAlbumAtPosition(rand.Intn(2500))
}