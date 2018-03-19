package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("ModuleInfo", func() event.Event {
		return new(ModuleInfo)
	})
}

type ModuleInfo struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e ModuleInfo) GetEvent() string {
	return e.Event
}

func (e ModuleInfo) GetTimestamp() time.Time {
	return e.Timestamp
}
