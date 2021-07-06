package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"os/exec"
)

const (
	filename      = "playlists.temp"
	defaultEditor = "vim"
)

// WriteToFile writes the playlists to a CSV for user editing
func WriteToFile(playlists []Playlist) {
	log.Println(len(playlists))
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Unable to create temp file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"-- Delete the playlists you wish to retain. Everything else will be killed! --", ""})

	for _, playlist := range playlists {
		err := writer.Write([]string{playlist.Name, playlist.ID})
		if err != nil {
			log.Fatal("Unable to write to temp file: ", err)
		}
	}
}

// OpenInEditor opens the temp file in an editor for user editing
func OpenInEditor() {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = defaultEditor
	}
	// Get the full executable path for the editor.
	executable, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal("Unable to locate editor executable: ", err)
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal("Error occurred in opening file for edit: ", err)
	}
}

// ReadIDsToDelete reads the edited temp file and extracts the ID strings
func ReadIDsToDelete() []string {
	toDelete := []string{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to read temp file: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for i := 0; ; i++ {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil && i != 0 {
			log.Fatal("Error reading CSV: ", err)
		}

		toDelete = append(toDelete, line[1])
	}

	return toDelete
}
