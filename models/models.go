package models

import (
  "time"
)

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
