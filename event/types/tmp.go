package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Temp", func() event.Event {
		return new(Temp)
	})
}

type Temp struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Temp) GetEvent() string {
	return e.Event
}

func (e Temp) GetTimestamp() time.Time {
	return e.Timestamp
}
