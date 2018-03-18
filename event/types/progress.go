package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Progress", func() event.Event {
		return new(Progress)
	})
}

type Progress struct {
	Cqc        int64     `json:"CQC"`
	Combat     int64     `json:"Combat"`
	Empire     int64     `json:"Empire"`
	Explore    int64     `json:"Explore"`
	Federation int64     `json:"Federation"`
	Trade      int64     `json:"Trade"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e Progress) GetEvent() string {
	return e.Event
}

func (e Progress) GetTimestamp() time.Time {
	return e.Timestamp
}
