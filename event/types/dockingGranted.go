package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("DockingGranted", func() event.Event {
		return new(DockingGranted)
	})
}

type DockingGranted struct {
	MarketID    int64     `json:"MarketID"`
	LandingPad  int64     `json:"LandingPad"`
	StationName string    `json:"StationName"`
	StationType string    `json:"StationType"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e DockingGranted) GetEvent() string {
	return e.Event
}

func (e DockingGranted) GetTimestamp() time.Time {
	return e.Timestamp
}
