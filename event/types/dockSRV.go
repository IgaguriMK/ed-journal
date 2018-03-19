package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("DockSRV", func() event.Event {
		return new(DockSRV)
	})
}

type DockSRV struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e DockSRV) GetEvent() string {
	return e.Event
}

func (e DockSRV) GetTimestamp() time.Time {
	return e.Timestamp
}
