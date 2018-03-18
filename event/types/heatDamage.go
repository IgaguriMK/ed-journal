package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("HeatDamage", func() event.Event {
		return new(HeatDamage)
	})
}

type HeatDamage struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e HeatDamage) GetEvent() string {
	return e.Event
}

func (e HeatDamage) GetTimestamp() time.Time {
	return e.Timestamp
}
