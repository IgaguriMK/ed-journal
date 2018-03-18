package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Materials", func() event.Event {
		return new(Materials)
	})
}

type Materials struct {
	Encoded []struct {
		Count int64  `json:"Count"`
		Name  string `json:"Name"`
	} `json:"Encoded"`
	Manufactured []struct {
		Count int64  `json:"Count"`
		Name  string `json:"Name"`
	} `json:"Manufactured"`
	Raw []struct {
		Count int64  `json:"Count"`
		Name  string `json:"Name"`
	} `json:"Raw"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Materials) GetEvent() string {
	return e.Event
}

func (e Materials) GetTimestamp() time.Time {
	return e.Timestamp
}
