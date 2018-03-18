package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Scanned", func() event.Event {
		return new(Scanned)
	})
}

type Scanned struct {
	ScanType  string    `json:"ScanType"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Scanned) GetEvent() string {
	return e.Event
}

func (e Scanned) GetTimestamp() time.Time {
	return e.Timestamp
}
