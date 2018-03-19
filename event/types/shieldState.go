package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("ShieldState", func() event.Event {
		return new(ShieldState)
	})
}

type ShieldState struct {
	ShieldsUp bool      `json:"ShieldsUp"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e ShieldState) GetEvent() string {
	return e.Event
}

func (e ShieldState) GetTimestamp() time.Time {
	return e.Timestamp
}
