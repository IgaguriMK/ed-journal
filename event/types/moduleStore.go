package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("ModuleStore", func() event.Event {
		return new(ModuleStore)
	})
}

type ModuleStore struct {
	EngineerModifications string    `json:"EngineerModifications"`
	Ship                  string    `json:"Ship"`
	ShipID                int64     `json:"ShipID"`
	Slot                  string    `json:"Slot"`
	StoredItem            string    `json:"StoredItem"`
	StoredItemLocalised   string    `json:"StoredItem_Localised"`
	Event                 string    `json:"event"`
	Timestamp             time.Time `json:"timestamp"`
}

func (e ModuleStore) GetEvent() string {
	return e.Event
}

func (e ModuleStore) GetTimestamp() time.Time {
	return e.Timestamp
}
