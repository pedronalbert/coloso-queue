package cache

import "os"
import "testing"
import "coloso-queue/models"
import "coloso-queue/cache"
import "coloso-queue/clients/mysql"

var sumTesting models.Summoner

func setup() {
  sumTesting = models.Summoner{
    ID: "LAN_12345",
    AccountID: "LAN_12345",
    Name: "Testing",
    SummonerLevel: 30,
    ProfileIconID: 1,
    Region: "LAN",
    RevisionDate: 12345,
  }
}

func TestMain(m *testing.M) {
  setup()
  code := m.Run();
  os.Exit(code)
}

func TestSaveSummoner(t *testing.T) {
  var sumInDb models.Summoner

  sumCreated := cache.SaveSummoner(sumTesting)

  mysql.Client.Where("id = ?", sumTesting.ID).First(&sumInDb)

  if sumInDb.ID != sumCreated.ID {
    t.Fatalf("Summoner saved not match\nexpected: %+v\ngot: %+v", sumCreated, sumInDb)
  }
}

// TestUpdateOnSaveSummoner - Change the name and update the same model
func TestUpdateOnSaveSummoner(t *testing.T) {
  var sumInDb models.Summoner

  sumTesting.Name = "Retesting"

  sumUpdated := cache.SaveSummoner(sumTesting)

  mysql.Client.Where("id = ?", sumTesting.ID).First(&sumInDb)

  if (sumInDb.Name != sumTesting.Name) || (sumUpdated.Name != sumInDb.Name) {
    t.Fatalf("Summoner is not updating on save\nexpected: %+v\ngot: %+v", sumTesting, sumInDb)
  }
}
