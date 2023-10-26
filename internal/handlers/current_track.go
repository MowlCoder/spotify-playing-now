package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MowlCod/dbus-spotify/internal/domain"
)

type currentTrackStorage interface {
	GetTrack(ctx context.Context) (*domain.Track, error)
}

type CurrentTrackHandler struct {
	storage currentTrackStorage
}

func NewCurrentTrackHandler(storage currentTrackStorage) *CurrentTrackHandler {
	return &CurrentTrackHandler{
		storage: storage,
	}
}

func (h *CurrentTrackHandler) GetCurrentTrack() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(405)
			return
		}

		track, err := h.storage.GetTrack(r.Context())

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
