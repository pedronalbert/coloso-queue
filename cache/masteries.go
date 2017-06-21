package cache

import (
	"coloso-queue/clients/mysql"
	"coloso-queue/models"
	"coloso-queue/utils"
	"time"
)

// MasteriesPageRaw - Runespages json as string
type MasteriesPageRaw struct {
	ID         int    `gorm:"column:id;AUTO_INCREMENT"`
	SummonerID string `gorm:"column:summonerId"`
	Pages      string `gorm:"column:pages;type:string" json:"pages"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// TableName - Runespages tablename
func (MasteriesPageRaw) TableName() string {
	return "masteries_pages"
}

// SaveMasteriesPage - Save pages in db
func SaveMasteriesPage(newMasteries models.MasteriesPages) models.MasteriesPages {
	var rawRunes MasteriesPageRaw

	rawRunes = MasteriesPageRaw{
		ID:         newMasteries.ID,
		SummonerID: newMasteries.SummonerID,
	}

	mysql.Client.First(&rawRunes)

	// Assign values to save
	rawRunes.SummonerID = newMasteries.SummonerID
	rawRunes.Pages = utils.StructToString(newMasteries.Pages)

	if rawRunes.ID == 0 {
		log.Debugf("Masteries not found in cache, creating new")

		mysql.Client.Create(&rawRunes)

		// Assign values saved to returned object
		newMasteries.ID = rawRunes.ID
		newMasteries.CreatedAt = rawRunes.CreatedAt
		newMasteries.UpdatedAt = rawRunes.UpdatedAt

		return newMasteries
	}

	log.Debugf("Masteries already stored in cache, updating data")

	mysql.Client.Save(&rawRunes)

	newMasteries.UpdatedAt = rawRunes.UpdatedAt

	return newMasteries
}
