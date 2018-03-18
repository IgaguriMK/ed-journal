package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("SellDrones", func() event.Event {
		return new(SellDrones)
	})
}

type SellDrones struct {
	Count     int64     `json:"Count"`
	SellPrice int64     `json:"SellPrice"`
	TotalSale int64     `json:"TotalSale"`
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e SellDrones) GetEvent() string {
	return e.Event
}

func (e SellDrones) GetTimestamp() time.Time {
	return e.Timestamp
}
