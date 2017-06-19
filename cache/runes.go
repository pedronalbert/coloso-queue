package cache

import (
	"coloso-queue/clients/mysql"
	"coloso-queue/models"
	"coloso-queue/utils"
	"time"
)

// RunesPagesRaw - Runespages json as string
type RunesPagesRaw struct {
	ID         int    `gorm:"column:id;AUTO_INCREMENT"`
	SummonerID string `gorm:"column:summonerId"`
	Pages      string `gorm:"column:pages;type:string" json:"pages"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// TableName - Runespages tablename
func (RunesPagesRaw) TableName() string {
	return "runes_pages"
}

// SaveRunesPages - Save pages in db
func SaveRunesPages(newRunes models.RunesPages) models.RunesPages {
	var rawRunes RunesPagesRaw

	rawRunes = RunesPagesRaw{
		ID:         newRunes.ID,
		SummonerID: newRunes.SummonerID,
	}

	mysql.Client.First(&rawRunes)

	// Assign values to save
	rawRunes.SummonerID = newRunes.SummonerID
	rawRunes.Pages = utils.StructToString(newRunes.Pages)

	if rawRunes.ID == 0 {
		log.Debugf("Runes not found in cache, creating new")

		mysql.Client.Create(&rawRunes)

		// Assign values saved to returned object
		newRunes.ID = rawRunes.ID
		newRunes.CreatedAt = rawRunes.CreatedAt
		newRunes.UpdatedAt = rawRunes.UpdatedAt

		return newRunes
	}

	log.Debugf("Runes already stored in cache, updating data")

	mysql.Client.Save(&rawRunes)

	newRunes.UpdatedAt = rawRunes.UpdatedAt

	return newRunes
}
