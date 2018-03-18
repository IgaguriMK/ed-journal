package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("WingLeave", func() event.Event {
		return new(WingLeave)
	})
}

type WingLeave struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e WingLeave) GetEvent() string {
	return e.Event
}

func (e WingLeave) GetTimestamp() time.Time {
	return e.Timestamp
}
