package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("ShipyardBuy", func() event.Event {
		return new(ShipyardBuy)
	})
}

type ShipyardBuy struct {
	MarketID           int64     `json:"MarketID"`
	SellOldShip        string    `json:"SellOldShip"`
	SellPrice          int64     `json:"SellPrice"`
	SellShipID         int64     `json:"SellShipID"`
	ShipPrice          int64     `json:"ShipPrice"`
	ShipType           string    `json:"ShipType"`
	ShipType_Localised string    `json:"ShipType_Localised"`
	StoreOldShip       string    `json:"StoreOldShip"`
	StoreShipID        int64     `json:"StoreShipID"`
	Event              string    `json:"event"`
	Timestamp          time.Time `json:"timestamp"`
}

func (e ShipyardBuy) GetEvent() string {
	return e.Event
}

func (e ShipyardBuy) GetTimestamp() time.Time {
	return e.Timestamp
}
