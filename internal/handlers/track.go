package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MowlCod/dbus-spotify/internal/domain"
)

type spotifyDBusTransport interface {
	GetCurrentTrackInfo() (*domain.Track, error)
	NextTrack() error
	PrevTrack() error
	PlayTrack() error
	PauseTrack() error
}

type TrackHandler struct {
	dbusTransport spotifyDBusTransport
}

func NewCurrentTrackHandler(transport spotifyDBusTransport) *TrackHandler {
	return &TrackHandler{
		dbusTransport: transport,
	}
}

func (h *TrackHandler) GetCurrentTrack() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		track, err := h.dbusTransport.GetCurrentTrackInfo()

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := json.Marshal(track)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})
}

func (h *TrackHandler) NextTrack() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := h.dbusTransport.NextTrack()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func (h *TrackHandler) PrevTrack() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := h.dbusTransport.PrevTrack()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func (h *TrackHandler) Play() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := h.dbusTransport.PlayTrack()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func (h *TrackHandler) Pause() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := h.dbusTransport.PauseTrack()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
