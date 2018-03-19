package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Touchdown", func() event.Event {
		return new(Touchdown)
	})
}

type Touchdown struct {
	Latitude         float64   `json:"Latitude"`
	Longitude        float64   `json:"Longitude"`
	PlayerControlled bool      `json:"PlayerControlled"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e Touchdown) GetEvent() string {
	return e.Event
}

func (e Touchdown) GetTimestamp() time.Time {
	return e.Timestamp
}
