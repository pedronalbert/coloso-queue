package riot

import (
  "time"
  "strings"
  "coloso-queue/utils"
  "coloso-queue/utils/urid"
  "coloso-queue/clients/riot/responses"
  "github.com/op/go-logging"
  "github.com/franela/goreq"
)

var riotKey = "RGAPI-9938cbc8-8a46-43f2-babe-784c4ffe6814"
var log = logging.MustGetLogger("riot_client")

func init() {
  goreq.SetConnectTimeout(30000 * time.Millisecond)
}

func createURL(region string, url string) string {
  platform := utils.RegionToPlatform(region)
  platform = strings.ToLower(platform)

  return "https://" + platform + ".api.riotgames.com/lol/" + url
}

func fetch(response interface{}, url string, query interface {}) error {
  log.Debugf("Obteniendo datos de Riot API %s", url)

  res, err := goreq.Request{ Uri: url, QueryString: query }.WithHeader("X-Riot-Token", riotKey).Do()

  if err != nil {
    switch res.StatusCode {
    case 400:
      return ErrBadRequest
    case 401:
      return ErrUnauthorized
    case 403:
      return ErrForbidden
    case 404:
      return ErrNotFound
    default:
      return ErrServerError
    }

    return err
  }

  res.Body.FromJsonTo(response)

  log.Debugf("Datos obtenidos correctamente: %s", url)

  return nil
}

// FetchSummonerByName - Devuelve la informacion del summoner
func FetchSummonerByName(name string, region string) (responses.Summoner, error) {
  var summoner responses.Summoner

  url := createURL(region, "summoner/v3/summoners/by-name/" + name)
  err := fetch(&summoner, url, nil)

  if err != nil {
    return summoner, err
  }

  return summoner, nil
}

// FetchSummonerByID - Devuelve la informacion del summoner por ID tipo URID
func FetchSummonerByID(sumURID string) (responses.Summoner, error) {
  var summoner responses.Summoner

  ID := urid.GetID(sumURID)
  region := urid.GetRegion(sumURID)

  url := createURL(region, "summoner/v3/summoners/" + ID)
  err := fetch(&summoner, url, nil)

  if err != nil {
    return summoner, err
  }

  return summoner, nil
}

// FetchRunesPages - Devuelve las runas de un invocador
func FetchRunesPages(sumURID string) (responses.RunesPages, error) {
  var runesPages responses.RunesPages

  ID := urid.GetID(sumURID)
  region := urid.GetRegion(sumURID)

  url := createURL(region, "platform/v3/runes/by-summoner/" + ID)
  err := fetch(&runesPages, url, nil)

  if err != nil {
    return runesPages, err
  }

  return runesPages, nil
}

// FetchMasteriesPages - Devuelve las runas de un invocador
func FetchMasteriesPages(sumURID string) (responses.MasteriesPages, error) {
  var masteriesPages responses.MasteriesPages

  ID := urid.GetID(sumURID)
  region := urid.GetRegion(sumURID)

  url := createURL(region, "platform/v3/masteries/by-summoner/" + ID)
  err := fetch(&masteriesPages, url, nil)

  if err != nil {
    return masteriesPages, err
  }

  return masteriesPages, nil
}

// FetchChampionsMasteries - Devuelve la lista de maestrías de un invocador
func FetchChampionsMasteries(sumURID string) ([]responses.ChampionMastery, error) {
  var championsMasteries []responses.ChampionMastery

  ID := urid.GetID(sumURID)
  region := urid.GetRegion(sumURID)

  url := createURL(region, "champion-mastery/v3/champion-masteries/by-summoner/" + ID)
  err := fetch(&championsMasteries, url, nil)

  if err != nil {
    return championsMasteries, err
  }

  return championsMasteries, nil
}

// FetchGame - Devuelve el juego
func FetchGame(gameURID string) (responses.Game, error) {
  var game responses.Game

  ID := urid.GetID(gameURID)
  region := urid.GetRegion(gameURID)

  url := createURL(region, "match/v3/matches/" + ID)
  err := fetch(&game, url, nil)

  if err != nil {
    return game, err
  }

  return game, nil
}

// FetchGameTimelines - Devuelve el juego
func FetchGameTimelines(gameURID string) (responses.GameTimelines, error) {
  var timelines responses.GameTimelines

  ID := urid.GetID(gameURID)
  region := urid.GetRegion(gameURID)

  url := createURL(region, "match/v3/timelines/by-match/" + ID)
  err := fetch(&timelines, url, nil)

  if err != nil {
    return timelines, err
  }

  return timelines, nil
}

// FetchGamesList - Hace un fetch de los matchsList de la API
func FetchGamesList(accountURID string, query interface{}) (responses.GamesList, error) {
  var gamesList responses.GamesList

  ID := urid.GetID(accountURID)
  region := urid.GetRegion(accountURID)

  url := createURL(region, "match/v3/matchlists/by-account/" + ID)
  err := fetch(&gamesList, url, query)

  if err != nil {
    return gamesList, err
  }

  return gamesList, nil
}

// FetchLeaguePosition - Hace un fetch de la posiciones en la liga
func FetchLeaguePosition(sumURID string) ([]responses.LeaguePostition, error) {
  var leaguePosition []responses.LeaguePostition

  ID := urid.GetID(sumURID)
  region := urid.GetRegion(sumURID)

  url := createURL(region, "league/v3/positions/by-summoner/" + ID)
  err := fetch(&leaguePosition, url, nil)

  if err != nil {
    return leaguePosition, err
  }

  return leaguePosition, nil
}
