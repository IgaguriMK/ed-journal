package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Statistics", func() event.Event {
		return new(Statistics)
	})
}

type Statistics struct {
	BankAccount struct {
		CurrentWealth          int64 `json:"Current_Wealth"`
		InsuranceClaims        int64 `json:"Insurance_Claims"`
		SpentOnAmmoConsumables int64 `json:"Spent_On_Ammo_Consumables"`
		SpentOnFuel            int64 `json:"Spent_On_Fuel"`
		SpentOnInsurance       int64 `json:"Spent_On_Insurance"`
		SpentOnOutfitting      int64 `json:"Spent_On_Outfitting"`
		SpentOnRepairs         int64 `json:"Spent_On_Repairs"`
		SpentOnShips           int64 `json:"Spent_On_Ships"`
	} `json:"Bank_Account"`
	Cqc struct {
		CQCCreditsEarned int64   `json:"CQC_Credits_Earned"`
		CqcKd            float64 `json:"CQC_KD"`
		CQCKills         int64   `json:"CQC_Kills"`
		CQCTimePlayed    int64   `json:"CQC_Time_Played"`
		CqcWl            int64   `json:"CQC_WL"`
	} `json:"CQC"`
	Combat struct {
		AssassinationProfits int64 `json:"Assassination_Profits"`
		Assassinations       int64 `json:"Assassinations"`
		BountiesClaimed      int64 `json:"Bounties_Claimed"`
		BountyHuntingProfit  int64 `json:"Bounty_Hunting_Profit"`
		CombatBondProfits    int64 `json:"Combat_Bond_Profits"`
		CombatBonds          int64 `json:"Combat_Bonds"`
		HighestSingleReward  int64 `json:"Highest_Single_Reward"`
		SkimmersKilled       int64 `json:"Skimmers_Killed"`
	} `json:"Combat"`
	Crafting struct {
		CountOfUsedEngineers  int64 `json:"Count_Of_Used_Engineers"`
		RecipesGenerated      int64 `json:"Recipes_Generated"`
		RecipesGeneratedRank1 int64 `json:"Recipes_Generated_Rank_1"`
		RecipesGeneratedRank2 int64 `json:"Recipes_Generated_Rank_2"`
		RecipesGeneratedRank3 int64 `json:"Recipes_Generated_Rank_3"`
		RecipesGeneratedRank4 int64 `json:"Recipes_Generated_Rank_4"`
		RecipesGeneratedRank5 int64 `json:"Recipes_Generated_Rank_5"`
	} `json:"Crafting"`
	Crew  struct{} `json:"Crew"`
	Crime struct {
		BountiesReceived int64 `json:"Bounties_Received"`
		Fines            int64 `json:"Fines"`
		HighestBounty    int64 `json:"Highest_Bounty"`
		Notoriety        int64 `json:"Notoriety"`
		TotalBounties    int64 `json:"Total_Bounties"`
		TotalFines       int64 `json:"Total_Fines"`
	} `json:"Crime"`
	Exploration struct {
		ExplorationProfits        int64   `json:"Exploration_Profits"`
		GreatestDistanceFromStart float64 `json:"Greatest_Distance_From_Start"`
		HighestPayout             int64   `json:"Highest_Payout"`
		PlanetsScannedToLevel2    int64   `json:"Planets_Scanned_To_Level_2"`
		PlanetsScannedToLevel3    int64   `json:"Planets_Scanned_To_Level_3"`
		SystemsVisited            int64   `json:"Systems_Visited"`
		TimePlayed                int64   `json:"Time_Played"`
		TotalHyperspaceDistance   int64   `json:"Total_Hyperspace_Distance"`
		TotalHyperspaceJumps      int64   `json:"Total_Hyperspace_Jumps"`
	} `json:"Exploration"`
	MaterialTraderStats struct {
		MaterialsTraded int64 `json:"Materials_Traded"`
		TradesCompleted int64 `json:"Trades_Completed"`
	} `json:"Material_Trader_Stats"`
	Mining struct {
		MaterialsCollected int64 `json:"Materials_Collected"`
		MiningProfits      int64 `json:"Mining_Profits"`
		QuantityMined      int64 `json:"Quantity_Mined"`
	} `json:"Mining"`
	Multicrew struct {
		MulticrewCreditsTotal     int64 `json:"Multicrew_Credits_Total"`
		MulticrewFighterTimeTotal int64 `json:"Multicrew_Fighter_Time_Total"`
		MulticrewFinesTotal       int64 `json:"Multicrew_Fines_Total"`
		MulticrewGunnerTimeTotal  int64 `json:"Multicrew_Gunner_Time_Total"`
		MulticrewTimeTotal        int64 `json:"Multicrew_Time_Total"`
	} `json:"Multicrew"`
	Passengers struct {
		PassengersMissionsAccepted    int64 `json:"Passengers_Missions_Accepted"`
		PassengersMissionsBulk        int64 `json:"Passengers_Missions_Bulk"`
		PassengersMissionsDelivered   int64 `json:"Passengers_Missions_Delivered"`
		PassengersMissionsDisgruntled int64 `json:"Passengers_Missions_Disgruntled"`
		PassengersMissionsEjected     int64 `json:"Passengers_Missions_Ejected"`
		PassengersMissionsVIP         int64 `json:"Passengers_Missions_VIP"`
	} `json:"Passengers"`
	SearchAndRescue struct {
		SearchRescueCount  int64 `json:"SearchRescue_Count"`
		SearchRescueProfit int64 `json:"SearchRescue_Profit"`
		SearchRescueTraded int64 `json:"SearchRescue_Traded"`
	} `json:"Search_And_Rescue"`
	Smuggling struct {
		AverageProfit            float64 `json:"Average_Profit"`
		BlackMarketsProfits      int64   `json:"Black_Markets_Profits"`
		BlackMarketsTradedWith   int64   `json:"Black_Markets_Traded_With"`
		HighestSingleTransaction int64   `json:"Highest_Single_Transaction"`
		ResourcesSmuggled        int64   `json:"Resources_Smuggled"`
	} `json:"Smuggling"`
	Trading struct {
		AverageProfit            float64 `json:"Average_Profit"`
		HighestSingleTransaction int64   `json:"Highest_Single_Transaction"`
		MarketProfits            int64   `json:"Market_Profits"`
		MarketsTradedWith        int64   `json:"Markets_Traded_With"`
		ResourcesTraded          int64   `json:"Resources_Traded"`
	} `json:"Trading"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Statistics) GetEvent() string {
	return e.Event
}

func (e Statistics) GetTimestamp() time.Time {
	return e.Timestamp
}
