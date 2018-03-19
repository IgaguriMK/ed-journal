package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("ShipyardSell", func() event.Event {
		return new(ShipyardSell)
	})
}

type ShipyardSell struct {
	SellShipID int64     `json:"SellShipID"`
	ShipPrice  int64     `json:"ShipPrice"`
	ShipType   string    `json:"ShipType"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e ShipyardSell) GetEvent() string {
	return e.Event
}

func (e ShipyardSell) GetTimestamp() time.Time {
	return e.Timestamp
}
