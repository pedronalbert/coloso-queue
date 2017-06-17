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

func fetch(url string, response interface{}) error {
  log.Debugf("Obteniendo datos de Riot API %s", url)

  res, err := goreq.Request{ Uri: url }.WithHeader("X-Riot-Token", riotKey).Do()

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

  res.Body.FromJsonTo(&response)

  log.Debugf("Datos obtenidos correctamente: %s", url)

  return nil
}

// FetchSummonerByName - Devuelve la informacion del summoner
func FetchSummonerByName(name string, region string) (responses.Summoner, error) {
  var summoner responses.Summoner

  url := createURL(region, "summoner/v3/summoners/by-name/" + name)
  err := fetch(url, &summoner)

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
  err := fetch(url, &summoner)

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
  err := fetch(url, &runesPages)

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
  err := fetch(url, &masteriesPages)

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
  err := fetch(url, &championsMasteries)

  if err != nil {
    return championsMasteries, err
  }

  return championsMasteries, nil
}
