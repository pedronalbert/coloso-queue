package cache_test

import "os"
import "testing"

import "coloso-queue/models"
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

func compareSummoners(sumA models.Summoner, sumB models.Summoner) bool {
	var compareKeys = []string{"ID", "AccountID", "Name", "SummonerLevel", "ProfileIconID", "Region", "RevisionDate"}

	return utils.CompareStructs(sumA, sumB, compareKeys)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
