package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("SupercruiseEntry", func() event.Event {
		return new(SupercruiseEntry)
	})
}

type SupercruiseEntry struct {
	StarSystem string    `json:"StarSystem"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e SupercruiseEntry) GetEvent() string {
	return e.Event
}

func (e SupercruiseEntry) GetTimestamp() time.Time {
	return e.Timestamp
}
