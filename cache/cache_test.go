package cache_test

import "os"
import "testing"

import "coloso-queue/models"
import "coloso-queue/utils"
import "coloso-queue/clients/mysql"

var sumTesting models.Summoner
var runesTesting models.RunesPage
var masteriesTesting models.MasteriesPages

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

	runesTesting = models.RunesPage{
		SummonerID: sumTesting.ID,
		Pages: []models.RunePage{
			{
				ID:      1,
				Name:    "Testing",
				Current: false,
				Slots: []models.RunePageSlot{
					{RuneSlotID: 1, RuneID: 1},
					{RuneSlotID: 2, RuneID: 1},
				},
			},
		},
	}

	masteriesTesting = models.MasteriesPages{
		SummonerID: sumTesting.ID,
		Pages: []models.MasteryPage{
			{
				ID:      1,
				Name:    "Testing",
				Current: false,
				Masteries: []models.Mastery{
					{ID: 1, Rank: 1},
					{ID: 2, Rank: 1},
				},
			},
		},
	}
}

func shutdown() {
	mysql.Client.Delete(sumTesting)
}

func compareSummoners(sumA models.Summoner, sumB models.Summoner) bool {
	var compareKeys = []string{"ID", "AccountID", "Name", "SummonerLevel", "ProfileIconID", "Region", "RevisionDate"}

	return utils.CompareStructs(sumA, sumB, compareKeys)
}

func compareRunes(runA models.RunesPage, runB models.RunesPage) bool {
	var compareKeys = []string{"SummonerID", "Pages"}

	return utils.CompareStructs(runA, runB, compareKeys)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
