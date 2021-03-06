package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("SupercruiseExit", func() event.Event {
		return new(SupercruiseExit)
	})
}

type SupercruiseExit struct {
	Body          string    `json:"Body"`
	BodyID        int64     `json:"BodyID"`
	BodyType      string    `json:"BodyType"`
	StarSystem    string    `json:"StarSystem"`
	SystemAddress int64     `json:"SystemAddress"`
	Event         string    `json:"event"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e SupercruiseExit) GetEvent() string {
	return e.Event
}

func (e SupercruiseExit) GetTimestamp() time.Time {
	return e.Timestamp
}
