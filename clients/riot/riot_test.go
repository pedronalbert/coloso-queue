package riot_test

import "testing"
import "coloso-queue/clients/riot"

type account struct {
  AccountID int
  AccountURID string
  SummonerID int
  SummonerName string
  Region string
  URID string
  GameURID string
  GameID int
}

var testAccount = account{
  AccountID: 200050594,
  AccountURID: "LAN_200050594",
  SummonerID: 75453,
  SummonerName: "armaghyon",
  Region: "lan",
  URID: "LAN_75453",
  GameURID: "LAN_423649469",
  GameID: 423649469,
}

func TestFetchSummonerByName(t *testing.T) {
  summoner, err := riot.FetchSummonerByName(testAccount.SummonerName, testAccount.Region)

  if err != nil {
    t.Fatalf("%s", err)
  }

  if testAccount.AccountID != summoner.AccountID {
    t.Fatalf("AccountID not match \n" +
      "expected AccountID: %d got AccountID: %d",
    testAccount.AccountID, summoner.AccountID)
  }
}

func TestFetchSummonerByID(t *testing.T) {
  summoner, err := riot.FetchSummonerByID(testAccount.URID)

  if err != nil {
    t.Fatalf("%s", err)
  }

  if testAccount.AccountID != summoner.AccountID {
    t.Fatalf("AccountID not match \n" +
      "expected AccountID: %d got AccountID: %d",
    testAccount.AccountID, summoner.AccountID)
  }
}

func TestFetchRunesPages(t *testing.T) {
  runesPages, err := riot.FetchRunesPages(testAccount.URID)

  if err != nil {
    t.Fatalf("%s", err)
  }

  if testAccount.SummonerID != runesPages.SummonerID {
    t.Fatalf("SummonerIDs not match \n" +
      "expected SummonerID: %d got SummonerID: %d",
    testAccount.SummonerID, runesPages.SummonerID)
  }

  if len(runesPages.Pages) == 0 {
    t.Error("Runes pages are empty")
  }
}

func TestFetchMasteriesPages(t *testing.T) {
  masteriesPages, err := riot.FetchMasteriesPages(testAccount.URID)

  if err != nil {
    t.Fatalf("%s", err)
  }

  if testAccount.SummonerID != masteriesPages.SummonerID {
    t.Fatalf("SummonerIDs not match \n" +
      "expected SummonerID: %d got SummonerID: %d",
    testAccount.SummonerID, masteriesPages.SummonerID)
  }

  if len(masteriesPages.Pages) == 0 {
    t.Error("Masteries pages are empty")
  }
}

func TestFetchChampionsMasteries(t *testing.T) {
  championsMasteries, err := riot.FetchChampionsMasteries(testAccount.URID)

  if err != nil {
    t.Fatalf("%s", err)
  }

  if len(championsMasteries) == 0 {
    t.Fatalf("Champions masteries are empty")
  }
}

func TestFetchGame(t *testing.T) {
  game, err := riot.FetchGame(testAccount.GameURID)

  if err != nil {
    t.Fatalf("%s", err)
  }

  if game.GameID != testAccount.GameID {
    t.Fatalf("GameID not match \nexpected GameID: %d got GameID: %d", testAccount.GameID, game.GameID)
  }
}

func TestFetchGameTimelines(t *testing.T) {
  timelines, err := riot.FetchGameTimelines(testAccount.GameURID)

  if err != nil {
    t.Fatalf("%s", err)
  }

  if len(timelines.Frames) == 0 {
    t.Fatalf("Frames are empty")
  }
}

func TestFetchGamesList(t *testing.T) {
  games, err := riot.FetchGamesList(testAccount.AccountURID, nil)

  if err != nil {
    t.Fatalf("%s", err)
  }

  if len(games.Matches) == 0 {
    t.Fatalf("Game list is empty")
  }
}

func TestFetchLeaguePosition(t *testing.T) {
  leaguePos, err := riot.FetchLeaguePosition(testAccount.URID)

  if err != nil {
    t.Fatalf("%s", err)
  }

  if len(leaguePos) == 0 {
    t.Fatalf("League position is empty")
  }
}
