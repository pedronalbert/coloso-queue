package cache_test

import "os"
import "testing"
import "strings"
import "coloso-queue/models"
import "coloso-queue/cache"
import "coloso-queue/clients/mysql"
import "coloso-queue/utils"

var sumTesting models.Summoner

func setup() {
	sumTesting = models.Summoner{
		ID:            "LAN_12345",
		AccountID:     "LAN_12345",
		Name:          "Testing",
		SummonerLevel: 30,
		ProfileIconID: 1,
		Region:        "LAN",
		RevisionDate:  12345,
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestSaveSummoner(t *testing.T) {
	var sumInDb models.Summoner
	var compareKeys = []string{"ID", "AccountID", "Name", "SummonerLevel", "ProfileIconID", "Region", "RevisionDate"}

	sumCreated := cache.SaveSummoner(sumTesting)

	if !utils.CompareStructs(sumCreated, sumTesting, compareKeys) {
		t.Fatalf("Summoner returned by cache not match\nexpected: %+v\ngot: %+v", sumTesting, sumCreated)
	}

	mysql.Client.Where("id = ?", sumTesting.ID).First(&sumInDb)

	if !utils.CompareStructs(sumCreated, sumInDb, compareKeys) {
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

	if sumFound.ID != sumTesting.ID {
		t.Fatalf("Summoner found not match\nexpected ID: %s got ID: %s", sumTesting.ID, sumFound.ID)
	}
}

func TestFindSummonerByName(t *testing.T) {
	sumFound, err := cache.FindSummonerByName(sumTesting.Name, sumTesting.Region)

	if err != nil {
		t.Fatalf("Can't find summoner name: %s, region: %s", sumTesting.Name, sumTesting.Region)
	}

	if sumFound.ID != sumTesting.ID {
		t.Fatalf("Summoner found not match\nexpected ID: %s got ID: %s", sumTesting.ID, sumFound.ID)
	}

	// Test uppercased name
	sumFound, err = cache.FindSummonerByName(sumTesting.Name, sumTesting.Region)

	if err != nil {
		t.Fatalf("Can't find summoner name: %s, region: %s", strings.ToUpper(sumTesting.Name), sumTesting.Region)
	}

	if sumFound.ID != sumTesting.ID {
		t.Fatalf("Summoner found not match\nexpected ID: %s got ID: %s", sumTesting.ID, sumFound.ID)
	}
}
