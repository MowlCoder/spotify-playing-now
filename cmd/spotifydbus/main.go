package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MowlCod/dbus-spotify/internal/config"
	"github.com/MowlCod/dbus-spotify/internal/dbusconn"
	"github.com/MowlCod/dbus-spotify/internal/handlers"
	"github.com/MowlCod/dbus-spotify/internal/storage"
	"github.com/MowlCod/dbus-spotify/internal/workers"
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

	trackStorage, err := storage.NewCurrentTrackStorage()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create track storage:", err)
		os.Exit(1)
	}

	updateTrackWorker := workers.NewUpdateCurrentTrackWorker(
		trackStorage,
		conn,
	)

	go updateTrackWorker.Start(context.Background(), appCfg.UpdateInterval)

	currentTrackHandler := handlers.NewCurrentTrackHandler(trackStorage)

	http.Handle("/current-track", currentTrackHandler.GetCurrentTrack())

	fmt.Printf(
		"DBus spotify started on port %d.\nYou can get current track by request to /current-track endpoint.\n",
		appCfg.HttpPort,
	)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", appCfg.HttpPort), nil); err != nil {
		log.Fatalln(err)
	}
}
