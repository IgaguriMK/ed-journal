package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("SelfDestruct", func() event.Event {
		return new(SelfDestruct)
	})
}

type SelfDestruct struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e SelfDestruct) GetEvent() string {
	return e.Event
}

func (e SelfDestruct) GetTimestamp() time.Time {
	return e.Timestamp
}
