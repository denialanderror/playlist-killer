package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	token := os.Args[1]
	user := GetUser(token)

	log.Println("Fetching playlists for user token...")
	playlists := user.GetPlaylists()
	WriteToFile(playlists)

	OpenInEditor()

	ids := ReadIDsToDelete()
	log.Printf("Killing %d playlists. Do you wish to proceed [y/n]?", len(ids)-1)

	var yesno string
	fmt.Scanln(&yesno)

	if yesno == "y" {
		for _, id := range ids {
			user.UnfollowPlaylist(id)
		}
		log.Println("Playlists killed!")
	}
}
