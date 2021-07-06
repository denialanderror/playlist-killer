package main

import (
	"fmt"
	"log"
	"os"
	"sync"
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
		var wg sync.WaitGroup
		for _, id := range ids {
			wg.Add(1)
			go func(id string) {
				user.UnfollowPlaylist(id)
				wg.Done()
			}(id)
		}
		wg.Wait()
		log.Println("Playlists killed!")
	}
}
