package domain

import (
	"fmt"
	"strings"
)

type Track struct {
	Title     string   `json:"title"`
	URL       string   `json:"url"`
	Artist    []string `json:"artist"`
	Album     string   `json:"album"`
	ArtURL    string   `json:"art_url"`
	State     string   `json:"state"`
	Position  int64    `json:"position"`
	MaxLength uint64   `json:"max_length"`
}

func (t *Track) GetPreview() string {
	return fmt.Sprintf(`
Artists: %s
Track: %s
Album: %s
State: %s,
Link: %s`, strings.Join(t.Artist, ", "), t.Title, t.Album, t.State, t.URL)
}
