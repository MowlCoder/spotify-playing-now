.PHONY:
.SILENT:

build:
	go build -o ./bin/spotify-dbus ./cmd/spotifydbus/main.go

run: build
	./bin/spotify-dbus