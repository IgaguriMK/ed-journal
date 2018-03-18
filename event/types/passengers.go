package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Passengers", func() event.Event {
		return new(Passengers)
	})
}

type Passengers struct {
	Manifest []struct {
		Count     int64  `json:"Count"`
		MissionID int64  `json:"MissionID"`
		Type      string `json:"Type"`
		Vip       bool   `json:"VIP"`
		Wanted    bool   `json:"Wanted"`
	} `json:"Manifest"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Passengers) GetEvent() string {
	return e.Event
}

func (e Passengers) GetTimestamp() time.Time {
	return e.Timestamp
}
