package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("CockpitBreached", func() event.Event {
		return new(CockpitBreached)
	})
}

type CockpitBreached struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e CockpitBreached) GetEvent() string {
	return e.Event
}

func (e CockpitBreached) GetTimestamp() time.Time {
	return e.Timestamp
}
