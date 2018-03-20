package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Undocked", func() event.Event {
		return new(Undocked)
	})
}

type Undocked struct {
	MarketID    int64     `json:"MarketID"`
	StationName string    `json:"StationName"`
	StationType string    `json:"StationType"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e Undocked) GetEvent() string {
	return e.Event
}

func (e Undocked) GetTimestamp() time.Time {
	return e.Timestamp
}
