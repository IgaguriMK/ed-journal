package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("RepairAll", func() event.Event {
		return new(RepairAll)
	})
}

type RepairAll struct {
	Cost      int64     `json:"Cost"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e RepairAll) GetEvent() string {
	return e.Event
}

func (e RepairAll) GetTimestamp() time.Time {
	return e.Timestamp
}
