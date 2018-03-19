package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Missions", func() event.Event {
		return new(Missions)
	})
}

type Missions struct {
	Active []struct {
	} `json:"Active"`
	Complete []struct {
	} `json:"Complete"`
	Failed []struct {
	} `json:"Failed"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Missions) GetEvent() string {
	return e.Event
}

func (e Missions) GetTimestamp() time.Time {
	return e.Timestamp
}
