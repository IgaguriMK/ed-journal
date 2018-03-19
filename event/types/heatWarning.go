package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("HeatWarning", func() event.Event {
		return new(HeatWarning)
	})
}

type HeatWarning struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e HeatWarning) GetEvent() string {
	return e.Event
}

func (e HeatWarning) GetTimestamp() time.Time {
	return e.Timestamp
}
