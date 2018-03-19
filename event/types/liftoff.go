package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Liftoff", func() event.Event {
		return new(Liftoff)
	})
}

type Liftoff struct {
	Latitude         float64   `json:"Latitude"`
	Longitude        float64   `json:"Longitude"`
	PlayerControlled bool      `json:"PlayerControlled"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e Liftoff) GetEvent() string {
	return e.Event
}

func (e Liftoff) GetTimestamp() time.Time {
	return e.Timestamp
}
