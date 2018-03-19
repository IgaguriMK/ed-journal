package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Powerplay", func() event.Event {
		return new(Powerplay)
	})
}

type Powerplay struct {
	Power       string    `json:"Power"`
	Rank        int64     `json:"Rank"`
	Merits      int64     `json:"Merits"`
	Votes       int64     `json:"Votes"`
	TimePledged int64     `json:"TimePledged"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e Powerplay) GetEvent() string {
	return e.Event
}

func (e Powerplay) GetTimestamp() time.Time {
	return e.Timestamp
}
