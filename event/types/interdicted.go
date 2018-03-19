package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Interdicted", func() event.Event {
		return new(Interdicted)
	})
}

type Interdicted struct {
	Faction     string    `json:"Faction"`
	Interdictor string    `json:"Interdictor"`
	IsPlayer    bool      `json:"IsPlayer"`
	Submitted   bool      `json:"Submitted"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e Interdicted) GetEvent() string {
	return e.Event
}

func (e Interdicted) GetTimestamp() time.Time {
	return e.Timestamp
}
