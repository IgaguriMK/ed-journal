package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("DataScanned", func() event.Event {
		return new(DataScanned)
	})
}

type DataScanned struct {
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e DataScanned) GetEvent() string {
	return e.Event
}

func (e DataScanned) GetTimestamp() time.Time {
	return e.Timestamp
}
