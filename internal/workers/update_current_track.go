package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/MowlCod/dbus-spotify/internal/domain"
)

type currentTrackStorage interface {
	SetTrack(ctx context.Context, track *domain.Track) error
}

type spotifyDBusConn interface {
	GetCurrentTrackInfo() (*domain.Track, error)
}

type UpdateCurrentTrackWorker struct {
	currentTrackStorage currentTrackStorage
	spotifyDBusConn     spotifyDBusConn
}

func NewUpdateCurrentTrackWorker(
	storage currentTrackStorage,
	conn spotifyDBusConn,
) *UpdateCurrentTrackWorker {
	return &UpdateCurrentTrackWorker{
		currentTrackStorage: storage,
		spotifyDBusConn:     conn,
	}
}

func (worker *UpdateCurrentTrackWorker) Start(ctx context.Context, updateInterval int) {
	ticker := time.NewTicker(time.Second * time.Duration(updateInterval))

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			currentTrack, err := worker.spotifyDBusConn.GetCurrentTrackInfo()

			if err != nil {
				fmt.Println("ERROR: can not get current track from spotify:", err)
				continue
			}

			err = worker.currentTrackStorage.SetTrack(ctx, currentTrack)

			if err != nil {
				fmt.Println("ERROR: can not save current track:", err)
			}
		}
	}
}
