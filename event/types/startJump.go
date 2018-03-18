package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("StartJump", func() event.Event {
		return new(StartJump)
	})
}

type StartJump struct {
	JumpType   string    `json:"JumpType"`
	StarClass  string    `json:"StarClass"`
	StarSystem string    `json:"StarSystem"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e StartJump) GetEvent() string {
	return e.Event
}

func (e StartJump) GetTimestamp() time.Time {
	return e.Timestamp
}
