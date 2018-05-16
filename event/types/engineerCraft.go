package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("EngineerCraft", func() event.Event {
		return new(EngineerCraft)
	})
}

type EngineerCraft struct {
	ApplyExperimentalEffect     string `json:"ApplyExperimentalEffect"`
	Blueprint                   string `json:"Blueprint"`
	BlueprintID                 int64  `json:"BlueprintID"`
	BlueprintName               string `json:"BlueprintName"`
	Engineer                    string `json:"Engineer"`
	EngineerID                  int64  `json:"EngineerID"`
	ExperimentalEffect          string `json:"ExperimentalEffect"`
	ExperimentalEffectLocalised string `json:"ExperimentalEffect_Localised"`
	Ingredients                 []struct {
		Count         int64  `json:"Count"`
		Name          string `json:"Name"`
		NameLocalised string `json:"Name_Localised"`
	} `json:"Ingredients"`
	Level     int64 `json:"Level"`
	Modifiers []struct {
		Label         string  `json:"Label"`
		LessIsGood    int64   `json:"LessIsGood"`
		OriginalValue float64 `json:"OriginalValue"`
		Value         float64 `json:"Value"`
	} `json:"Modifiers"`
	Module    string    `json:"Module"`
	Quality   float64   `json:"Quality"`
	Slot      string    `json:"Slot"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e EngineerCraft) GetEvent() string {
	return e.Event
}

func (e EngineerCraft) GetTimestamp() time.Time {
	return e.Timestamp
}
