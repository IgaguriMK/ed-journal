package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Commander", func() event.Event {
		return new(Commander)
	})
}

type Commander struct {
	Event     string    `json:"event"`
	Name      string    `json:"Name"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Commander) GetEvent() string {
	return e.Event
}

func (e Commander) GetTimestamp() time.Time {
	return e.Timestamp
}
