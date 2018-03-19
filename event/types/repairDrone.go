package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("RepairDrone", func() event.Event {
		return new(RepairDrone)
	})
}

type RepairDrone struct {
	HullRepaired float64   `json:"HullRepaired"`
	Event        string    `json:"event"`
	Timestamp    time.Time `json:"timestamp"`
}

func (e RepairDrone) GetEvent() string {
	return e.Event
}

func (e RepairDrone) GetTimestamp() time.Time {
	return e.Timestamp
}
