package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Died", func() event.Event {
		return new(Died)
	})
}

type Died struct {
	KillerName          string    `json:"KillerName"`
	KillerNameLocalised string    `json:"KillerName_Localised"`
	KillerRank          string    `json:"KillerRank"`
	KillerShip          string    `json:"KillerShip"`
	Event               string    `json:"event"`
	Timestamp           time.Time `json:"timestamp"`
}

func (e Died) GetEvent() string {
	return e.Event
}

func (e Died) GetTimestamp() time.Time {
	return e.Timestamp
}
