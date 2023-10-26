# Spotify Playing Now

## Overview
The idea of this project is convert spotify dbus interface to http interface. You can serve this app on the chosen port and you will be able to get current track by using GET http request and control tracks (next, prev, play, pause) by using POST requests.

## Build & Run

```shell
make build
./bin/spotify-dbus -p 7766
```

## Usage

App is listening requests to `GET /current-track`
```shell
curl "127.0.0.1:7766/current-track"
```
Response will be like this:
```json
{
  "title": "Haunting Me",
  "url": "https://open.spotify.com/track/7zHehwGCfT6befxyOenhgE",
  "artist": ["Snow Strippers"],
  "album": "The Snow Strippers",
  "art_url": "https://i.scdn.co/image/ab67616d0000b27380ab44e51da54976bfeeb1c4",
  "state":"Playing",
  "position": 882000,
  "max_length": 144378000
}
```

Possible status codes:

- `200` - OK
- `400` - Program can't get track from spotify
- `405` - Method not allowed
- `500` - Error when trying to prepare response

Other handlers:
- `POST /next-track` - Play next track
- `POST /prev-track` - Play previous track
- `POST /play` - Play current track
- `POST /pause` - Pause current track

They have similar status codes to `GET /current-track` and response empty body.

## Config
You can execute program with flag `-h` and you'll see all flag options.

## TODO

- ~~Add support to track control (Play, Pause, Next, Prev)~~