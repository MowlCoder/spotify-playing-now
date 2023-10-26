package storage

import (
	"context"

	domain2 "github.com/MowlCod/dbus-spotify/internal/domain"
)

type CurrentTrackStorage struct {
	currentTrack *domain2.Track
}

func NewCurrentTrackStorage() (*CurrentTrackStorage, error) {
	return &CurrentTrackStorage{
		currentTrack: nil,
	}, nil
}

func (storage *CurrentTrackStorage) SetTrack(ctx context.Context, track *domain2.Track) error {
	storage.currentTrack = track

	return nil
}

func (storage *CurrentTrackStorage) GetTrack(ctx context.Context) (*domain2.Track, error) {
	if storage.currentTrack == nil {
		return nil, domain2.ErrTrackNotSet
	}

	return storage.currentTrack, nil
}
