package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MowlCod/dbus-spotify/internal/config"
	"github.com/MowlCod/dbus-spotify/internal/dbusconn"
	"github.com/MowlCod/dbus-spotify/internal/handlers"
)

func main() {
	appCfg := &config.AppConfig{}
	appCfg.ParseFlags()

	conn, err := dbusconn.New()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	currentTrackHandler := handlers.NewCurrentTrackHandler(conn)

	http.Handle("/current-track", currentTrackHandler.GetCurrentTrack())
	http.Handle("/next-track", currentTrackHandler.NextTrack())
	http.Handle("/prev-track", currentTrackHandler.PrevTrack())
	http.Handle("/play", currentTrackHandler.Play())
	http.Handle("/pause", currentTrackHandler.Pause())

	fmt.Printf(
		"DBus spotify started on port %d.\nYou can get current track by request to /current-track endpoint.\n",
		appCfg.HttpPort,
	)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", appCfg.HttpPort), nil); err != nil {
		log.Fatalln(err)
	}
}
