package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Promotion", func() event.Event {
		return new(Promotion)
	})
}

type Promotion struct {
	Combat     int64     `json:"Combat"`
	Explore    int64     `json:"Explore"`
	Federation int64     `json:"Federation"`
	Trade      int64     `json:"Trade"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e Promotion) GetEvent() string {
	return e.Event
}

func (e Promotion) GetTimestamp() time.Time {
	return e.Timestamp
}
