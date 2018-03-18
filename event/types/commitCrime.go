package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("CommitCrime", func() event.Event {
		return new(CommitCrime)
	})
}

type CommitCrime struct {
	CrimeType string    `json:"CrimeType"`
	Faction   string    `json:"Faction"`
	Fine      int64     `json:"Fine"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e CommitCrime) GetEvent() string {
	return e.Event
}

func (e CommitCrime) GetTimestamp() time.Time {
	return e.Timestamp
}
