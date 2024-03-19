package main

import (
	"encoding/csv"
	"log"
	"os"
)

const (
	fileNameSource = "spotify_data.csv"
)

func openCSVFile() (*csv.Reader, *os.File, error) {
	file, err := os.Open(fileNameSource)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	reader := csv.NewReader(file)
	// read header
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
		return nil, file, err
	}
	return reader, file, nil
}

func readTracks() {
	reader, file, err := openCSVFile()
	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()
	if err != nil {
		return
	}

	// read tracks
	var record []string
	spotifyTracks := make([]SpotifyTrack, batchSize)
	var spotifyTrack *SpotifyTrack
	index := 0
	iRead := 0
	for {
		if record, err = reader.Read(); err != nil {
			break
		}
		if spotifyTrack, err = createSpotifyTrack(record); err != nil {
			break
		}
		spotifyTracks[iRead] = *spotifyTrack
		index++
		iRead = index % batchSize
		if iRead == 0 {
			jobs <- spotifyTracks
			spotifyTracks = make([]SpotifyTrack, batchSize)
		}
		if index == nSpotifyTracks {
			break
		}
	}
	if err != nil {
		log.Panic(err)
		return
	}
	if iRead > 0 {
		jobs <- spotifyTracks[0:iRead]
	}
	close(jobs)
}
