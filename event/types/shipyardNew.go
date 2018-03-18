package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("ShipyardNew", func() event.Event {
		return new(ShipyardNew)
	})
}

type ShipyardNew struct {
	NewShipID int64     `json:"NewShipID"`
	ShipType  string    `json:"ShipType"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e ShipyardNew) GetEvent() string {
	return e.Event
}

func (e ShipyardNew) GetTimestamp() time.Time {
	return e.Timestamp
}
