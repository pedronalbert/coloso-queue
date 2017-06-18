package models

import (
  "time"
)

// Summoner model
type Summoner struct{
  ID string `gorm:"primary_key"`
  AccountID string `gorm:"column:accountId"`
  Name string `gorm:"column:name"`
  SummonerLevel int `gorm:"column:summonerLevel"`
  ProfileIconID int `gorm:"column:profileIconId"`
  Region string `gorm:"column:region"`
  RevisionDate int `gorm:"column:revisionDate"`
  CreatedAt time.Time
  UpdatedAt time.Time
}

// Game model
type Game struct {
  ID string `gorm:"column:id"`
  SeasonID int `gorm:"column:seasonId"`
  QueueID int `gorm:"column:queueId"`
  GameVersion string `gorm:"column:gameVersion"`
  GameMode string `gorm:"column:gameMode"`
  MapID int `gorm:"column:mapId"`
  GameType string `gorm:"column:gameType"`
  Teams string `gorm:"column:teams"`
  ParticipantsIdentities string `gorm:"column:participantIdentities"`
  Participants string `gorm:"column:participants"`
  GameDuration int `gorm:"column:gameDuration"`
  GameCreation int `gorm:"column:gameCreation"`
  Fetched bool
  ApproxFetch int `gorm:"column:approxFetch"`
  CreatedAt time.Time
  UpdatedAt time.Time
}
