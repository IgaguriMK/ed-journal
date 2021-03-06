package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("MarketBuy", func() event.Event {
		return new(MarketBuy)
	})
}

type MarketBuy struct {
	BuyPrice       int64     `json:"BuyPrice"`
	Count          int64     `json:"Count"`
	MarketID       int64     `json:"MarketID"`
	TotalCost      int64     `json:"TotalCost"`
	Type           string    `json:"Type"`
	Type_Localised string    `json:"Type_Localised"`
	Event          string    `json:"event"`
	Timestamp      time.Time `json:"timestamp"`
}

func (e MarketBuy) GetEvent() string {
	return e.Event
}

func (e MarketBuy) GetTimestamp() time.Time {
	return e.Timestamp
}
