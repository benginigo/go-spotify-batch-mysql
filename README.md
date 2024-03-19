## deploying a MySQL database in Docker

```shell
docker run -d --name db-spotify-tracks \
    -p 3400:3306 \
    -e MYSQL_ROOT_PASSWORD=password \
    -e MYSQL_DATABASE=spotify \
    -d mysql:8.0
```

## creating a Spotify tracks table

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


## execute script
```shell
DB_USERNAME=root \
DB_PASSWORD=password \
DB_HOST=localhost \
DB_PORT=3400 \
DB_DATABASE=spotify \
go run .
```
