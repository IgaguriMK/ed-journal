package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("DockingRequested", func() event.Event {
		return new(DockingRequested)
	})
}

type DockingRequested struct {
	MarketID    int64     `json:"MarketID"`
	StationName string    `json:"StationName"`
	StationType string    `json:"StationType"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e DockingRequested) GetEvent() string {
	return e.Event
}

func (e DockingRequested) GetTimestamp() time.Time {
	return e.Timestamp
}
