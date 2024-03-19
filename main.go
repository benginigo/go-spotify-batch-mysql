package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	batchSize      = 3000
	nWorkers       = 15
	nSpotifyTracks = 1000000
)

var (
	nBatches             = (nSpotifyTracks + batchSize - 1) / batchSize
	queryParamsToBatches = ""
	jobs                 = make(chan []SpotifyTrack, nBatches)
)

func main() {
	db, err := connectMysql()
	if err != nil {
		return
	}
	defer func() {
		_ = db.Close()
	}()

	startTime := time.Now()

	// read data
	readTracks()

	// prepare params
	queryParamsToBatches = prepareQueryParamsToBatch()

	// start workers
	done := make(chan struct{}, nWorkers)
	for iWorker := 0; iWorker < nWorkers; iWorker++ {
		go spotifyTracksWorker(db, jobs, done)
	}

	// wait to workers
	for iWorker := 0; iWorker < nWorkers; iWorker++ {
		<-done
	}

	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime).Seconds())
}

func spotifyTracksWorker(db *sql.DB, batch <-chan []SpotifyTrack, done chan<- struct{}) {
	for tracks := range batch {
		if len(tracks) == batchSize {
			err := insertSpotifyTracksPrepareBatch(db, tracks)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := insertSpotifyTracksBatch(db, tracks)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	done <- struct{}{}
}
