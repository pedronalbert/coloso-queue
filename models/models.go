package models

import (
	"time"
)

// Summoner model
type Summoner struct {
	ID            string `gorm:"primary_key"`
	AccountID     string `gorm:"column:accountId"`
	Name          string `gorm:"column:name"`
	SummonerLevel int    `gorm:"column:summonerLevel"`
	ProfileIconID int    `gorm:"column:profileIconId"`
	Region        string `gorm:"column:region"`
	RevisionDate  int    `gorm:"column:revisionDate"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// RunePageSlot - RunesPages.Pages.Slot
type RunePageSlot struct {
	RuneSlotID int `json:"runeSlotId"`
	RuneID     int `json:"runeId"`
}

// RunePage - RunesPages.Pages
type RunePage struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Current bool           `json:"current"`
	Slots   []RunePageSlot `json:"slots"`
}

// RunesPage - Esquema de respuesta de runas
type RunesPage struct {
	ID         int        `gorm:"column:id;AUTO_INCREMENT"`
	SummonerID string     `gorm:"column:summonerId"`
	Pages      []RunePage `gorm:"column:pages" json:"pages"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Game model
type Game struct {
	ID                     string `gorm:"column:id"`
	SeasonID               int    `gorm:"column:seasonId"`
	QueueID                int    `gorm:"column:queueId"`
	GameVersion            string `gorm:"column:gameVersion"`
	GameMode               string `gorm:"column:gameMode"`
	MapID                  int    `gorm:"column:mapId"`
	GameType               string `gorm:"column:gameType"`
	Teams                  string `gorm:"column:teams"`
	ParticipantsIdentities string `gorm:"column:participantIdentities"`
	Participants           string `gorm:"column:participants"`
	GameDuration           int    `gorm:"column:gameDuration"`
	GameCreation           int    `gorm:"column:gameCreation"`
	Fetched                bool
	ApproxFetch            int `gorm:"column:approxFetch"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
