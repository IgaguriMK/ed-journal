package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("ModuleSell", func() event.Event {
		return new(ModuleSell)
	})
}

type ModuleSell struct {
	MarketID          int64     `json:"MarketID"`
	SellItem          string    `json:"SellItem"`
	SellItemLocalised string    `json:"SellItem_Localised"`
	SellPrice         int64     `json:"SellPrice"`
	Ship              string    `json:"Ship"`
	ShipID            int64     `json:"ShipID"`
	Slot              string    `json:"Slot"`
	Event             string    `json:"event"`
	Timestamp         time.Time `json:"timestamp"`
}

func (e ModuleSell) GetEvent() string {
	return e.Event
}

func (e ModuleSell) GetTimestamp() time.Time {
	return e.Timestamp
}
