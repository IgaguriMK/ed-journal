package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("LaunchSRV", func() event.Event {
		return new(LaunchSRV)
	})
}

type LaunchSRV struct {
	Loadout          string    `json:"Loadout"`
	PlayerControlled bool      `json:"PlayerControlled"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e LaunchSRV) GetEvent() string {
	return e.Event
}

func (e LaunchSRV) GetTimestamp() time.Time {
	return e.Timestamp
}
