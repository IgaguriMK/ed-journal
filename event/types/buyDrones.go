package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("BuyDrones", func() event.Event {
		return new(BuyDrones)
	})
}

type BuyDrones struct {
	BuyPrice  int64     `json:"BuyPrice"`
	Count     int64     `json:"Count"`
	TotalCost int64     `json:"TotalCost"`
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e BuyDrones) GetEvent() string {
	return e.Event
}

func (e BuyDrones) GetTimestamp() time.Time {
	return e.Timestamp
}
