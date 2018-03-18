package types

import ( "time"
"github.com/IgaguriMK/ed-journal/event" )

func init() {
	event.RegisterEvent("BuyExplorationData", func() event.Event {
		return new(BuyExplorationData)
	})
}

type BuyExplorationData struct {
	Cost      int64     `json:"Cost"`
	System    string    `json:"System"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e BuyExplorationData) GetEvent() string {
	return e.Event
}

func (e BuyExplorationData) GetTimestamp() time.Time {
	return e.Timestamp
}
