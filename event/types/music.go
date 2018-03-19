package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Music", func() event.Event {
		return new(Music)
	})
}

type Music struct {
	MusicTrack string    `json:"MusicTrack"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e Music) GetEvent() string {
	return e.Event
}

func (e Music) GetTimestamp() time.Time {
	return e.Timestamp
}
