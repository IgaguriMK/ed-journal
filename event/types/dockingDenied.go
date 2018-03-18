package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("DockingDenied", func() event.Event {
		return new(DockingDenied)
	})
}

type DockingDenied struct {
	Reason      string    `json:"Reason"`
	StationName string    `json:"StationName"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e DockingDenied) GetEvent() string {
	return e.Event
}

func (e DockingDenied) GetTimestamp() time.Time {
	return e.Timestamp
}
