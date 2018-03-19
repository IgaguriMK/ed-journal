package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("LeaveBody", func() event.Event {
		return new(LeaveBody)
	})
}

type LeaveBody struct {
	Body          string    `json:"Body"`
	BodyID        int64     `json:"BodyID"`
	StarSystem    string    `json:"StarSystem"`
	SystemAddress int64     `json:"SystemAddress"`
	Event         string    `json:"event"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e LeaveBody) GetEvent() string {
	return e.Event
}

func (e LeaveBody) GetTimestamp() time.Time {
	return e.Timestamp
}
