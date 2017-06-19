package cache_test

import "testing"
import "coloso-queue/models"
import "coloso-queue/cache"
import "coloso-queue/clients/mysql"

func TestSaveSummoner(t *testing.T) {
	var sumInDb models.Summoner

	sumCreated := cache.SaveSummoner(sumTesting)

	if !compareSummoners(sumCreated, sumTesting) {
		t.Fatalf("Summoner returned by cache not match\nexpected: %+v\ngot: %+v", sumTesting, sumCreated)
	}

	mysql.Client.Where("id = ?", sumTesting.ID).First(&sumInDb)

	if !compareSummoners(sumCreated, sumInDb) {
		t.Fatalf("Summoner returned by mysql not match\nexpected: %+v\ngot: %+v", sumTesting, sumCreated)
	}
}

// TestUpdateOnSaveSummoner - Change the name and update the same model
func TestSaveSummonerUpdateData(t *testing.T) {
	sumTesting.Name = "Retesting"

	TestSaveSummoner(t)
}

func TestFindSummonerById(t *testing.T) {
	sumFound, err := cache.FindSummonerByID(sumTesting.ID)

	if err != nil {
		t.Fatalf("Can't find summoner ID: %s in cache\nerror: %s", sumTesting.ID, err)
	}

	if !compareSummoners(sumFound, sumTesting) {
		t.Fatalf("Summoner found in cache not match with testing\nexpected: %+v\ngot: %+v", sumTesting, sumFound)
	}
}

func TestFindSummonerByIdNotFound(t *testing.T) {
	var err error

	_, err = cache.FindSummonerByID("")

	if err != nil && err != cache.ErrNotFound {
		t.Fatalf("Error returned not match\nexpected: %s\ngot: %s", cache.ErrNotFound, err)
	}
}

func TestFindSummonerByName(t *testing.T) {
	sumFound, err := cache.FindSummonerByName(sumTesting.Name, sumTesting.Region)

	if err != nil {
		t.Fatalf("Can't find summoner name: %s in cache\nerror: %s", sumTesting.Name, err)
	}

	if !compareSummoners(sumFound, sumTesting) {
		t.Fatalf("Summoner found in cache not match with testing\nexpected: %+v\ngot: %+v", sumTesting, sumFound)
	}
}

func TestFindSummonerByNameNotFound(t *testing.T) {
	var err error

	_, err = cache.FindSummonerByName("", "")

	if err != nil && err != cache.ErrNotFound {
		t.Fatalf("Error returned not match\nexpected: %s\ngot: %s", cache.ErrNotFound, err)
	}
}
