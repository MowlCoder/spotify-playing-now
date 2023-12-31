package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/MowlCod/dbus-spotify/internal/config"
	"github.com/MowlCod/dbus-spotify/internal/dbusconn"
	"github.com/MowlCod/dbus-spotify/internal/handlers"
)

func main() {
	appCfg := &config.AppConfig{}
	appCfg.ParseFlags()

	_, err := exec.Command("pidof", "spotify").Output()
	if err != nil {
		fmt.Println("WARNING: It seems you haven't running Spotify. You have to run Spotify.")
	}

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
		"INFO: DBus spotify started on port %d\n",
		appCfg.HttpPort,
	)

	fmt.Println(`
GET  /current-track - Get current track info
POST /next-track - Switch to next track
POST /prev-track - Switch to previous track
POST /play - Start playing current track
POST /pause - Pause current track`)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", appCfg.HttpPort), nil); err != nil {
		log.Fatalln(err)
	}
}
