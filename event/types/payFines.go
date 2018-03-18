package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("PayFines", func() event.Event {
		return new(PayFines)
	})
}

type PayFines struct {
	Amount    int64     `json:"Amount"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e PayFines) GetEvent() string {
	return e.Event
}

func (e PayFines) GetTimestamp() time.Time {
	return e.Timestamp
}
