package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("MaterialDiscarded", func() event.Event {
		return new(MaterialDiscarded)
	})
}

type MaterialDiscarded struct {
	Category  string    `json:"Category"`
	Count     int64     `json:"Count"`
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e MaterialDiscarded) GetEvent() string {
	return e.Event
}

func (e MaterialDiscarded) GetTimestamp() time.Time {
	return e.Timestamp
}
