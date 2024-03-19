package main

import (
	"fmt"
	"strconv"
)

type SpotifyTrack struct {
	ArtistName       string  `csv:"artist_name"`
	TrackName        string  `csv:"track_name"`
	TrackID          string  `csv:"track_id"`
	Popularity       int     `csv:"popularity"`
	Year             int     `csv:"year"`
	Genre            string  `csv:"genre"`
	Danceability     float64 `csv:"danceability"`
	Energy           float64 `csv:"energy"`
	Key              int     `csv:"key"`
	Loudness         float64 `csv:"loudness"`
	Mode             int     `csv:"mode"`
	Speechiness      float64 `csv:"speechiness"`
	Acousticness     float64 `csv:"acousticness"`
	Instrumentalness float64 `csv:"instrumentalness"`
	Liveness         float64 `csv:"liveness"`
	Valence          float64 `csv:"valence"`
	Tempo            float64 `csv:"tempo"`
	DurationMS       int     `csv:"duration_ms"`
	TimeSignature    int     `csv:"time_signature"`
}

func createSpotifyTrack(record []string) (*SpotifyTrack, error) {
	popularity, err := strconv.Atoi(record[4])
	if err != nil {
		return nil, fmt.Errorf("error converting Popularity: %w", err)
	}

	year, err := strconv.Atoi(record[5])
	if err != nil {
		return nil, fmt.Errorf("error converting Year: %w", err)
	}

	key, err := strconv.Atoi(record[9])
	if err != nil {
		return nil, fmt.Errorf("error converting Key: %w", err)
	}

	mode, err := strconv.Atoi(record[11])
	if err != nil {
		return nil, fmt.Errorf("error converting Mode: %w", err)
	}

	durationMS, err := strconv.Atoi(record[18])
	if err != nil {
		return nil, fmt.Errorf("error converting Duration: %w", err)
	}

	timeSignature, err := strconv.Atoi(record[19])
	if err != nil {
		return nil, fmt.Errorf("error converting Time Signature: %w", err)
	}

	spotifyTrack := SpotifyTrack{
		ArtistName:       record[1],
		TrackName:        record[2],
		TrackID:          record[3],
		Popularity:       popularity,
		Year:             year,
		Genre:            record[6],
		Danceability:     parseFloat(record[7]),
		Energy:           parseFloat(record[8]),
		Key:              key,
		Loudness:         parseFloat(record[10]),
		Mode:             mode,
		Speechiness:      parseFloat(record[12]),
		Acousticness:     parseFloat(record[13]),
		Instrumentalness: parseFloat(record[14]),
		Liveness:         parseFloat(record[15]),
		Valence:          parseFloat(record[16]),
		Tempo:            parseFloat(record[17]),
		DurationMS:       durationMS,
		TimeSignature:    timeSignature,
	}

	return &spotifyTrack, nil
}

func parseFloat(value string) float64 {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0
	}
	return result
}
