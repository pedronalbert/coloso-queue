package cache_test

import (
	"coloso-queue/cache"
	"testing"
)

import "coloso-queue/clients/mysql"

func TestSaveMasteries(t *testing.T) {
	var masteriesInDB cache.MasteriesPageRaw

	_ = cache.SaveMasteriesPage(masteriesTesting)

	mysql.Client.Where("summonerId = ?", masteriesTesting.SummonerID).First(&masteriesInDB)

	if masteriesInDB.ID == 0 {
		t.Fatalf("Masteries pages are not store on DB")
	}

	if masteriesInDB.Pages == "" {
		t.Fatalf("Masteries pages are not stored correctly on DB got nil")
	}
}
