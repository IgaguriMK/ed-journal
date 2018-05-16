package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("MaterialCollected", func() event.Event {
		return new(MaterialCollected)
	})
}

type MaterialCollected struct {
	Category      string    `json:"Category"`
	Count         int64     `json:"Count"`
	Name          string    `json:"Name"`
	NameLocalised string    `json:"Name_Localised"`
	Event         string    `json:"event"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e MaterialCollected) GetEvent() string {
	return e.Event
}

func (e MaterialCollected) GetTimestamp() time.Time {
	return e.Timestamp
}
