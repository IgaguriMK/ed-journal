package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("MaterialDiscovered", func() event.Event {
		return new(MaterialDiscovered)
	})
}

type MaterialDiscovered struct {
	Category        string    `json:"Category"`
	DiscoveryNumber int64     `json:"DiscoveryNumber"`
	Name            string    `json:"Name"`
	Event           string    `json:"event"`
	Timestamp       time.Time `json:"timestamp"`
}

func (e MaterialDiscovered) GetEvent() string {
	return e.Event
}

func (e MaterialDiscovered) GetTimestamp() time.Time {
	return e.Timestamp
}
