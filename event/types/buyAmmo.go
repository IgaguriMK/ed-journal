package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("BuyAmmo", func() event.Event {
		return new(BuyAmmo)
	})
}

type BuyAmmo struct {
	Cost      int64     `json:"Cost"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e BuyAmmo) GetEvent() string {
	return e.Event
}

func (e BuyAmmo) GetTimestamp() time.Time {
	return e.Timestamp
}
