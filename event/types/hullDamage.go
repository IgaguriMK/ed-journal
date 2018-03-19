package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("HullDamage", func() event.Event {
		return new(HullDamage)
	})
}

type HullDamage struct {
	Fighter     bool      `json:"Fighter"`
	Health      float64   `json:"Health"`
	PlayerPilot bool      `json:"PlayerPilot"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e HullDamage) GetEvent() string {
	return e.Event
}

func (e HullDamage) GetTimestamp() time.Time {
	return e.Timestamp
}
