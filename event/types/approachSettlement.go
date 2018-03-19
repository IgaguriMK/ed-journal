package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("ApproachSettlement", func() event.Event {
		return new(ApproachSettlement)
	})
}

type ApproachSettlement struct {
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e ApproachSettlement) GetEvent() string {
	return e.Event
}

func (e ApproachSettlement) GetTimestamp() time.Time {
	return e.Timestamp
}
