# playlist-killer

A hacky Golang program to clean up my Spotify playlists. **WARNING - This program can and will delete all your playlists. Use caution!**

## Problem

I have been using Spotify for well over a decade, including many years before they provided the ability to save albums and follow artists separately, rather than having to create a playlist for each. As a result, I have a long list of "album" playlists that I want rid of. Instead of right-clicking each individually, it gave a good opportunity to check out the Spotify API and practice some Go.

## Runing the Application

### Prerequisites:
* Golang
* A Spotify account

### Running Playlist Killer:
* Go to the [Spotify API reference](https://developer.spotify.com/console/get-current-user-playlists/)
* Click "Get Token"
* Select `playlist-read-private` and `playlist-modify-private`, click "Request Token", then copy it for later
* Build the program with `go build .`
* Run `./playlist-killer <YOUR TOKEN HERE>`

### Using Playlist Killer:

The program will use the user token to fetch all playlists and write them to a file. It will then open the file using the user's default editor (i.e. `$EDITOR`). The user will see a CSV containing the playlist's title and uuid. Delete the rows you **wish to keep** and when ready, save the file and quit the editor. The program will then print the total number of playlists to delete and will ask for confirmation. **This is your last opportunity to abort**. The playlists will then be deleted.
