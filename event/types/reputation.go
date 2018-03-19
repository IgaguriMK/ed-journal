package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Reputation", func() event.Event {
		return new(Reputation)
	})
}

type Reputation struct {
	Empire     float64   `json:"Empire"`
	Federation float64   `json:"Federation"`
	Alliance   float64   `json:"Alliance"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e Reputation) GetEvent() string {
	return e.Event
}

func (e Reputation) GetTimestamp() time.Time {
	return e.Timestamp
}
