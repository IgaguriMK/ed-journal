package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("CollectCargo", func() event.Event {
		return new(CollectCargo)
	})
}

type CollectCargo struct {
	Stolen        bool      `json:"Stolen"`
	Type          string    `json:"Type"`
	TypeLocalised string    `json:"Type_Localised"`
	Event         string    `json:"event"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e CollectCargo) GetEvent() string {
	return e.Event
}

func (e CollectCargo) GetTimestamp() time.Time {
	return e.Timestamp
}
