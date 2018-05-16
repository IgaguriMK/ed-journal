package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Cargo", func() event.Event {
		return new(Cargo)
	})
}

type Cargo struct {
	Inventory []struct {
		Count         int64  `json:"Count"`
		Name          string `json:"Name"`
		NameLocalised string `json:"Name_Localised"`
		Stolen        int64  `json:"Stolen"`
	} `json:"Inventory"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Cargo) GetEvent() string {
	return e.Event
}

func (e Cargo) GetTimestamp() time.Time {
	return e.Timestamp
}
