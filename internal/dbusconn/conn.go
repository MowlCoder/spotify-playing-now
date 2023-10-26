package dbusconn

import (
	"github.com/godbus/dbus/v5"

	"github.com/MowlCod/dbus-spotify/internal/domain"
)

// Commands:
// Toggle track: dbus-send --print-reply --dest=org.mpris.MediaPlayer2.spotify /org/mpris/MediaPlayer2 org.mpris.MediaPlayer2.Player.PlayPause
// Get Metadata: dbus-send --print-reply --dest=org.mpris.MediaPlayer2.spotify /org/mpris/MediaPlayer2 org.freedesktop.DBus.Properties.Get string:org.mpris.MediaPlayer2.Player string:Metadata

type DbusConn struct {
	conn *dbus.Conn
}

func New() (*DbusConn, error) {
	conn, err := dbus.ConnectSessionBus()

	if err != nil {
		return nil, err
	}

	return &DbusConn{
		conn: conn,
	}, err
}

func (dc *DbusConn) GetCurrentTrackInfo() (*domain.Track, error) {
	var metaDataReply map[string]dbus.Variant
	var playbackStatusReplay dbus.Variant
	var positionReplay dbus.Variant

	obj := dc.conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2")
	err := obj.Call("org.freedesktop.DBus.Properties.Get", 0, "org.mpris.MediaPlayer2.Player", "Metadata").Store(&metaDataReply)

	if err != nil {
		return nil, err
	}

	err = obj.Call("org.freedesktop.DBus.Properties.Get", 0, "org.mpris.MediaPlayer2.Player", "PlaybackStatus").Store(&playbackStatusReplay)

	if err != nil {
		return nil, err
	}

	err = obj.Call("org.freedesktop.DBus.Properties.Get", 0, "org.mpris.MediaPlayer2.Player", "Position").Store(&positionReplay)

	if err != nil {
		return nil, err
	}

	return &domain.Track{
		Title:     metaDataReply["xesam:title"].Value().(string),
		Album:     metaDataReply["xesam:album"].Value().(string),
		Artist:    metaDataReply["xesam:artist"].Value().([]string),
		ArtURL:    metaDataReply["mpris:artUrl"].Value().(string),
		URL:       metaDataReply["xesam:url"].Value().(string),
		Position:  positionReplay.Value().(int64),
		MaxLength: metaDataReply["mpris:length"].Value().(uint64),
		State:     playbackStatusReplay.Value().(string),
	}, nil
}

func (dc *DbusConn) Close() {
	dc.Close()
}
