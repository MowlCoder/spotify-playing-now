package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MowlCod/dbus-spotify/internal/domain"
)

type spotifyDBusTransport interface {
	GetCurrentTrackInfo() (*domain.Track, error)
}

type CurrentTrackHandler struct {
	dbusTransport spotifyDBusTransport
}

func NewCurrentTrackHandler(transport spotifyDBusTransport) *CurrentTrackHandler {
	return &CurrentTrackHandler{
		dbusTransport: transport,
	}
}

func (h *CurrentTrackHandler) GetCurrentTrack() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(405)
			return
		}

		track, err := h.dbusTransport.GetCurrentTrackInfo()

		if err != nil {
			w.WriteHeader(400)
			return
		}

		body, err := json.Marshal(track)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})
}
