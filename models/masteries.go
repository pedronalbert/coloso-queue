package models

import "time"

// Mastery - Mastery
type Mastery struct {
	ID   int `json:"id"`
	Rank int `json:"rank"`
}

// MasteryPage - MasteryPage
type MasteryPage struct {
	Current   bool      `json:"current"`
	Masteries []Mastery `json:"masteries"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
}

// MasteriesPages - MasteriesPages
type MasteriesPages struct {
	ID         int           `gorm:"column:id;AUTO_INCREMENT"`
	SummonerID string        `gorm:"column:summonerId"`
	Pages      []MasteryPage `gorm:"column:pages" json:"pages"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
