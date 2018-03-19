package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("WingJoin", func() event.Event {
		return new(WingJoin)
	})
}

type WingJoin struct {
	Others    []string  `json:"Others"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e WingJoin) GetEvent() string {
	return e.Event
}

func (e WingJoin) GetTimestamp() time.Time {
	return e.Timestamp
}
