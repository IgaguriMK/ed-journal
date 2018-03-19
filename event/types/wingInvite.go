package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("WingInvite", func() event.Event {
		return new(WingInvite)
	})
}

type WingInvite struct {
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e WingInvite) GetEvent() string {
	return e.Event
}

func (e WingInvite) GetTimestamp() time.Time {
	return e.Timestamp
}
