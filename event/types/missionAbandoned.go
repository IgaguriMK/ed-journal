package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("MissionAbandoned", func() event.Event {
		return new(MissionAbandoned)
	})
}

type MissionAbandoned struct {
	MissionID int64     `json:"MissionID"`
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e MissionAbandoned) GetEvent() string {
	return e.Event
}

func (e MissionAbandoned) GetTimestamp() time.Time {
	return e.Timestamp
}
