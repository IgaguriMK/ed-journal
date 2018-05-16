package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Docked", func() event.Event {
		return new(Docked)
	})
}

type Docked struct {
	CockpitBreach     bool    `json:"CockpitBreach"`
	DistFromStarLS    float64 `json:"DistFromStarLS"`
	FactionState      string  `json:"FactionState"`
	MarketID          int64   `json:"MarketID"`
	StarSystem        string  `json:"StarSystem"`
	StationAllegiance string  `json:"StationAllegiance"`
	StationEconomy    string  `json:"StationEconomy"`
	StationEconomies  []struct {
		Name           string  `json:"Name"`
		Name_Localised string  `json:"Name_Localised"`
		Proportion     float64 `json:"Proportion"`
	} ` json:"StationEconomies"`
	StationEconomyLocalised    string    `json:"StationEconomy_Localised"`
	StationFaction             string    `json:"StationFaction"`
	StationGovernment          string    `json:"StationGovernment"`
	StationGovernmentLocalised string    `json:"StationGovernment_Localised"`
	StationName                string    `json:"StationName"`
	StationServices            []string  `json:"StationServices"`
	StationState               string    `json:"StationState"`
	StationType                string    `json:"StationType"`
	SystemAddress              int64     `json:"SystemAddress"`
	Event                      string    `json:"event"`
	Timestamp                  time.Time `json:"timestamp"`
}

func (e Docked) GetEvent() string {
	return e.Event
}

func (e Docked) GetTimestamp() time.Time {
	return e.Timestamp
}
