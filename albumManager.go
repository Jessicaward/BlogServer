package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/shkh/lastfm-go/lastfm"
)

type Album struct {
	Name             string
	PositionOnLastFm string
	Artist           string
	ImageUrl         string
	PlayCount        string
}

//Constructor for Album struct
func NewAlbum(name string, artist string, positionOnLastFm string, imageUrl string, playCount string) *Album {
	newImageUrl := strings.Replace(imageUrl, "34s", "128s", -1)
	album := Album{Name: name,
		PositionOnLastFm: positionOnLastFm,
		Artist:           artist,
		ImageUrl:         newImageUrl,
		PlayCount:        playCount}
	return &album
}

// Get all albums from Last.fm
func GetAlbums() []*Album {
	lastFmDetails := GetLastFmConfiguration()
	var albums []*Album

	api := lastfm.New(lastFmDetails.ApiKey, lastFmDetails.ApiSecret)
	result, _ := api.User.GetTopAlbums(lastfm.P{"user": "jessicaward25"}) //discarding error

	for _, album := range result.Albums {
		albums = append(albums, NewAlbum(album.Name, album.Artist.Name, album.Rank, album.Images[0].Url, album.PlayCount))
	}

	return albums
}

//Get random album at specified position
func GetAlbumAtPosition(position int) *Album {
	lastFmDetails := GetLastFmConfiguration()
	var album *Album

	api := lastfm.New(lastFmDetails.ApiKey, lastFmDetails.ApiSecret)
	result, _ := api.User.GetTopAlbums(lastfm.P{"user": "jessicaward25", "limit": "1", "page": strconv.Itoa(position)}) //discarding error

	for _, r := range result.Albums {
		album = NewAlbum(r.Name, r.Artist.Name, r.Rank, r.Images[0].Url, r.PlayCount)
	}

	return album
}

//Get random album from Last.fm
func GetRandomAlbum() *Album {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(1000)
	fmt.Println("random index: " + strconv.Itoa(rnd))
	return GetAlbumAtPosition(rnd)
}

func loadAlbum() *Album {
	album := GetRandomAlbum()
	return &Album{Name: album.Name,
		Artist:           album.Artist,
		PositionOnLastFm: album.PositionOnLastFm,
		ImageUrl:         album.ImageUrl,
		PlayCount:        album.PlayCount,
	}
}
