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
	var rawMasteries MasteriesPageRaw

	rawMasteries = MasteriesPageRaw{
		ID:         newMasteries.ID,
		SummonerID: newMasteries.SummonerID,
	}

	mysql.Client.First(&rawMasteries)

	// Assign values to save
	rawMasteries.SummonerID = newMasteries.SummonerID
	rawMasteries.Pages = utils.StructToString(newMasteries.Pages)

	if rawMasteries.ID == 0 {
		log.Debugf("Masteries not found in cache, creating new")

		mysql.Client.Create(&rawMasteries)

		// Assign values saved to returned object
		newMasteries.ID = rawMasteries.ID
		newMasteries.CreatedAt = rawMasteries.CreatedAt
		newMasteries.UpdatedAt = rawMasteries.UpdatedAt

		return newMasteries
	}

	log.Debugf("Masteries already stored in cache, updating data")

	mysql.Client.Save(&rawMasteries)

	newMasteries.ID = rawMasteries.ID
	newMasteries.CreatedAt = rawMasteries.CreatedAt
	newMasteries.UpdatedAt = rawMasteries.UpdatedAt

	return newMasteries
}
