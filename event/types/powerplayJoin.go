package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("PowerplayJoin", func() event.Event {
		return new(PowerplayJoin)
	})
}

type PowerplayJoin struct {
	Power     string    `json:"Power"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e PowerplayJoin) GetEvent() string {
	return e.Event
}

func (e PowerplayJoin) GetTimestamp() time.Time {
	return e.Timestamp
}
