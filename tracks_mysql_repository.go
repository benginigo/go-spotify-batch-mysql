package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	queryParams = "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	nColumns    = 19
)

//go:embed sql/create_tracks.sql
var QueryCreateTracks string

func connectMysql() (*sql.DB, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	if username == "" || password == "" || host == "" || port == "" || database == "" {
		err := fmt.Errorf("missing environment variables")
		log.Fatal(err)
		return nil, err
	}

	// build the connection string using the environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	// Open the connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, err
}

func insertSpotifyTracksPrepareBatch(db *sql.DB, tracks []SpotifyTrack) error {
	args := make([]interface{}, len(tracks)*nColumns)

	iArgs := 0
	for _, track := range tracks {
		args[iArgs] = track.ArtistName
		args[iArgs+1] = track.TrackName
		args[iArgs+2] = track.TrackID
		args[iArgs+3] = track.Popularity
		args[iArgs+4] = track.Year
		args[iArgs+5] = track.Genre
		args[iArgs+6] = track.Danceability
		args[iArgs+7] = track.Energy
		args[iArgs+8] = track.Key
		args[iArgs+9] = track.Loudness
		args[iArgs+10] = track.Mode
		args[iArgs+11] = track.Speechiness
		args[iArgs+12] = track.Acousticness
		args[iArgs+13] = track.Instrumentalness
		args[iArgs+14] = track.Liveness
		args[iArgs+15] = track.Valence
		args[iArgs+16] = track.Tempo
		args[iArgs+17] = track.DurationMS
		args[iArgs+18] = track.TimeSignature

		iArgs += 19
	}
	_, err := db.Exec(queryParamsToBatches, args...)
	return err
}

func insertSpotifyTracksBatch(db *sql.DB, tracks []SpotifyTrack) error {
	queryLength := len(QueryCreateTracks) + len(tracks)*(len(queryParams)-1)
	builder := strings.Builder{}
	builder.Grow(queryLength)
	builder.WriteString(QueryCreateTracks)

	// initialize the arguments slice
	args := make([]interface{}, len(tracks)*nColumns)

	iArgs := 0
	// build the list of arguments and placeholders
	for iTrack, track := range tracks {
		builder.WriteString(queryParams)
		args[iArgs] = track.ArtistName
		args[iArgs+1] = track.TrackName
		args[iArgs+2] = track.TrackID
		args[iArgs+3] = track.Popularity
		args[iArgs+4] = track.Year
		args[iArgs+5] = track.Genre
		args[iArgs+6] = track.Danceability
		args[iArgs+7] = track.Energy
		args[iArgs+8] = track.Key
		args[iArgs+9] = track.Loudness
		args[iArgs+10] = track.Mode
		args[iArgs+11] = track.Speechiness
		args[iArgs+12] = track.Acousticness
		args[iArgs+13] = track.Instrumentalness
		args[iArgs+14] = track.Liveness
		args[iArgs+15] = track.Valence
		args[iArgs+16] = track.Tempo
		args[iArgs+17] = track.DurationMS
		args[iArgs+18] = track.TimeSignature

		iArgs += nColumns

		// Add a comma to separate blocks except for the last block
		if iTrack < len(tracks)-1 {
			builder.WriteString(",")
		}
	}
	query := builder.String()
	// execute the bulk insertion
	_, err := db.Exec(query, args...)
	return err
}

func prepareQueryParamsToBatch() string {
	queryLength := len(QueryCreateTracks) + batchSize*(len(queryParams)-1)
	params := strings.Builder{}
	params.Grow(queryLength)
	params.WriteString(QueryCreateTracks)
	for iTrack := 0; iTrack < batchSize; iTrack++ {
		params.WriteString(queryParams)
		if iTrack < batchSize-1 {
			params.WriteString(",")
		}
	}
	return params.String()
}
