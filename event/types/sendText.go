package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("SendText", func() event.Event {
		return new(SendText)
	})
}

type SendText struct {
	Message   string    `json:"Message"`
	To        string    `json:"To"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e SendText) GetEvent() string {
	return e.Event
}

func (e SendText) GetTimestamp() time.Time {
	return e.Timestamp
}
