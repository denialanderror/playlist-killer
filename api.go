package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type playlistList struct {
	Items []Playlist `json:"items"`
	Next  string     `json:"next"`
}

// Playlist holds the Name and ID of the playlist
type Playlist struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// User holds the information to make an http request for the user
type User struct {
	token  string
	client *http.Client
}

// GetUser builds the user client
func GetUser(token string) User {
	return User{token, &http.Client{}}
}

// GetPlaylists returns the name and ID for all playlists for the user
func (user *User) GetPlaylists() []Playlist {
	playlists := []Playlist{}
	url := "https://api.spotify.com/v1/me/playlists"

	for len(url) > 0 {
		res := user.call("GET", url)

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal("Unable to parse response: ", err)
		}

		var data playlistList
		err = json.Unmarshal(body, &data)

		if err != nil {
			log.Fatal("Unable to decode json: ", err)
		}

		url = data.Next
		playlists = append(playlists, data.Items...)
	}

	return playlists
}

// UnfollowPlaylist 'deletes' a playlist for the user
func (user *User) UnfollowPlaylist(id string) {
	res := user.call("DELETE", "https://api.spotify.com/v1/playlists/"+id+"/followers")
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("Non-OK status received for delete request:", res.StatusCode)
	}
}

func (user *User) call(method string, url string) *http.Response {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal("Error initating request: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+user.token)

	res, err := user.client.Do(req)
	if err != nil {
		log.Fatal("Error making request: ", err)
	}

	return res
}
