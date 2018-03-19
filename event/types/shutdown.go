package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Shutdown", func() event.Event {
		return new(Shutdown)
	})
}

type Shutdown struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Shutdown) GetEvent() string {
	return e.Event
}

func (e Shutdown) GetTimestamp() time.Time {
	return e.Timestamp
}
