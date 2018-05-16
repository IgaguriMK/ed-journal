package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Loadout", func() event.Event {
		return new(Loadout)
	})
}

type Loadout struct {
	HullValue int64 `json:"HullValue"`
	Modules   []struct {
		AmmoInClip        int64  `json:"AmmoInClip"`
		AmmoInHopper      int64  `json:"AmmoInHopper"`
		EngineerBlueprint string `json:"EngineerBlueprint"`
		Engineering       struct {
			BlueprintID                 int64  `json:"BlueprintID"`
			BlueprintName               string `json:"BlueprintName"`
			Engineer                    string `json:"Engineer"`
			EngineerID                  int64  `json:"EngineerID"`
			ExperimentalEffect          string `json:"ExperimentalEffect"`
			ExperimentalEffectLocalised string `json:"ExperimentalEffect_Localised"`
			Level                       int64  `json:"Level"`
			Modifiers                   []struct {
				Label         string  `json:"Label"`
				LessIsGood    int64   `json:"LessIsGood"`
				OriginalValue float64 `json:"OriginalValue"`
				Value         float64 `json:"Value"`
			} `json:"Modifiers"`
			Quality float64 `json:"Quality"`
		} `json:"Engineering"`
		EngineerLevel int64   `json:"EngineerLevel"`
		Health        float64 `json:"Health"`
		Item          string  `json:"Item"`
		On            bool    `json:"On"`
		Priority      int64   `json:"Priority"`
		Slot          string  `json:"Slot"`
		Value         int64   `json:"Value"`
	} `json:"Modules"`
	ModulesValue int64     `json:"ModulesValue"`
	Rebuy        int64     `json:"Rebuy"`
	Ship         string    `json:"Ship"`
	ShipID       int64     `json:"ShipID"`
	ShipIdent    string    `json:"ShipIdent"`
	ShipName     string    `json:"ShipName"`
	Event        string    `json:"event"`
	Timestamp    time.Time `json:"timestamp"`
}

func (e Loadout) GetEvent() string {
	return e.Event
}

func (e Loadout) GetTimestamp() time.Time {
	return e.Timestamp
}
