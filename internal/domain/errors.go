package domain

import "errors"

var (
	ErrTrackNotSet = errors.New("current track not set yet")
)
