package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("DiscoveryScan", func() event.Event {
		return new(DiscoveryScan)
	})
}

type DiscoveryScan struct {
	Bodies        int64     `json:"Bodies"`
	Event         string    `json:"event"`
	SystemAddress int64     `json:"SystemAddress"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e DiscoveryScan) GetEvent() string {
	return e.Event
}

func (e DiscoveryScan) GetTimestamp() time.Time {
	return e.Timestamp
}
