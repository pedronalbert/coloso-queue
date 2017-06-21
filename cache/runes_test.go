package cache_test

import (
	"coloso-queue/cache"
	"testing"
)

import "coloso-queue/clients/mysql"

func TestSaveRunesPages(t *testing.T) {
	var runesInDb cache.RunesPagesRaw

	runesTesting = cache.SaveRunesPages(runesTesting)

	mysql.Client.Where("summonerId = ?", runesTesting.SummonerID).First(&runesInDb)

	if runesInDb.ID == 0 {
		t.Fatalf("Rune pages are not store on DB")
	}

	if runesInDb.Pages == "" {
		t.Fatalf("Rune pages are not stored correctly on DB got nil")
	}

	if runesTesting.ID == 0 {
		t.Fatalf("Save function is not returning ID")
	}
}

func TestFindRunesPage(t *testing.T) {
	runesFound, err := cache.FindRunesPage(sumTesting.ID)

	if err != nil {
		t.Fatalf("Can't find runes, error: %s", err)
	}

	if !compareRunes(runesTesting, runesFound) {
		t.Fatalf("Runes not match\nexpected: %+v\ngot: %+v", runesTesting, runesFound)
	}
}
