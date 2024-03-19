# Importing One Million Spotify Tracks into MySQL Using Goroutines in Go

## Introduction:
In this tutorial, we'll explore how to efficiently import a large dataset of Spotify tracks into a MySQL database using Go. We'll leverage goroutines to parallelize the data import process, allowing us to import one million tracks quickly and effectively.

## Project Overview:
The project consists of Go scripts that interact with a MySQL database and a CSV file containing Spotify track data. We'll break down the project structure and explain each component:

### main.go: 
This is the main entry point of the project. It orchestrates the data import process, reads data from the CSV file, prepares parameters, starts worker goroutines, and manages the overall execution flow.

### tracks_csv_repository.go: 
This file handles reading the Spotify track data from the CSV file. It opens the file, reads the tracks, and sends batches of tracks to worker goroutines for processing.

### tracks_entity.go: 
This file defines the SpotifyTrack struct and contains functions for creating Spotify tracks from CSV records.

### tracks_mysql_repository.go: 
This file is responsible for interacting with the MySQL database. It establishes a database connection, prepares batch insert queries, and executes bulk insertions of Spotify tracks into the database.

## Execution Flow:

The main.go script reads Spotify track data from the CSV file using tracks_csv_repository.go.
It prepares parameters and starts worker goroutines to process the data concurrently.
Each worker goroutine receives batches of tracks, prepares batch insert queries, and inserts the tracks into the MySQL database using tracks_mysql_repository.go.
The project leverages goroutines and parallel processing to efficiently import one million tracks into the MySQL database.
Conclusion:
In this post, we've demonstrated how to import a large dataset of Spotify tracks into a MySQL database using Go and goroutines. By parallelizing

the data import process, we were able to import one million tracks quickly and efficiently in approximately 78 to 80 seconds. This project showcases the power and flexibility of Go for handling large-scale data processing tasks.
You can find the complete source code and instructions for running the project on GitHub https://github.com/benginigo/go-spotify-batch-mysql.

## Deploying a MySQL database in Docker

```shell
docker run -d --name db-spotify-tracks \
    -p 3400:3306 \
    -e MYSQL_ROOT_PASSWORD=password \
    -e MYSQL_DATABASE=spotify \
    -d mysql:8.0
```

## Creating a Spotify tracks table

```mysql
create table if not exists spotify.spotify_tracks
(
    ArtistName       varchar(255) null,
    TrackName        text         null,
    TrackID          varchar(255) not null
        primary key,
    Popularity       int          null,
    Year             int          null,
    Genre            varchar(255) null,
    Danceability     double       null,
    Energy           double       null,
    KeyVal           int          null,
    Loudness         double       null,
    Mode             int          null,
    Speechiness      double       null,
    Acousticness     double       null,
    Instrumentalness double       null,
    Liveness         double       null,
    Valence          double       null,
    Tempo            double       null,
    DurationMS       int          null,
    TimeSignature    int          null
);
```

## Unzip tracks_csv_repository.zip


## Execute script
```shell
DB_USERNAME=root \
DB_PASSWORD=password \
DB_HOST=localhost \
DB_PORT=3400 \
DB_DATABASE=spotify \
go run .
```
