package cache_test

import (
	"coloso-queue/cache"
	"testing"
)

import "coloso-queue/clients/mysql"

func TestSaveRunesPages(t *testing.T) {
	var runesInDb cache.RunesPagesRaw

	_ = cache.SaveRunesPages(runesTesting)

	mysql.Client.Where("summonerId = ?", runesTesting.SummonerID).First(&runesInDb)

	if runesInDb.ID == 0 {
		t.Fatalf("Rune pages are not store on DB")
	}

	if runesInDb.Pages == "" {
		t.Fatalf("Rune pages are not stored correctly on DB got nil")
	}
}
