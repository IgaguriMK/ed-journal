package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("MissionRedirected", func() event.Event {
		return new(MissionRedirected)
	})
}

type MissionRedirected struct {
	MissionID             int64     `json:"MissionID"`
	Name                  string    `json:"Name"`
	NewDestinationStation string    `json:"NewDestinationStation"`
	NewDestinationSystem  string    `json:"NewDestinationSystem"`
	OldDestinationStation string    `json:"OldDestinationStation"`
	OldDestinationSystem  string    `json:"OldDestinationSystem"`
	Event                 string    `json:"event"`
	Timestamp             time.Time `json:"timestamp"`
}

func (e MissionRedirected) GetEvent() string {
	return e.Event
}

func (e MissionRedirected) GetTimestamp() time.Time {
	return e.Timestamp
}
