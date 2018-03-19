package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("ApproachBody", func() event.Event {
		return new(ApproachBody)
	})
}

type ApproachBody struct {
	Body          string    `json:"Body"`
	BodyID        int64     `json:"BodyID"`
	StarSystem    string    `json:"StarSystem"`
	SystemAddress int64     `json:"SystemAddress"`
	Event         string    `json:"event"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e ApproachBody) GetEvent() string {
	return e.Event
}

func (e ApproachBody) GetTimestamp() time.Time {
	return e.Timestamp
}
