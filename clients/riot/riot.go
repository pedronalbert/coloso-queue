package riot

import (
  "time"
  "strings"
  "coloso-queue/utils"
  "coloso-queue/utils/urid"
  "coloso-queue/clients/riot/entities"
  "coloso-queue/clients/riot/responses"
  "github.com/op/go-logging"
  "github.com/franela/goreq"
)

var riotKey = "RGAPI-9938cbc8-8a46-43f2-babe-784c4ffe6814"
var log = logging.MustGetLogger("riot_client")

func createURL(region string, url string) string {
  platform := utils.RegionToPlatform(region)
  platform = strings.ToLower(platform)

  return "https://" + platform + ".api.riotgames.com/lol/" + url
}

func fetch(url string, response interface{}) error {
  log.Debugf("Obteniendo datos de Riot API %s", url)

  res, err := goreq.Request{ Uri: url }.WithHeader("X-Riot-Token", riotKey).Do()

  if err != nil {
    return err
  }

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

  res.Body.FromJsonTo(&response)

  log.Debugf("Datos obtenidos correctamente: %s", url)

  return nil
}

func init() {
  goreq.SetConnectTimeout(30000 * time.Millisecond)
}

// FetchSummonerByName - Devuelve la informacion del summoner
func FetchSummonerByName(name string, region string) (entities.Summoner, error) {
  var response responses.Summoner
  var summoner entities.Summoner

  url := createURL(region, "summoner/v3/summoners/by-name/" + name)
  err := fetch(url, &response)

  if err != nil {
    return summoner, err
  }

  summoner = entities.Summoner{
    ID: urid.Generate(region, response.ID),
    ProfileIconID: response.ProfileIconID,
    Name: response.Name,
    SummonerLevel: response.SummonerLevel,
    AccountID: response.AccountID,
    RevisionDate: response.RevisionDate,
  }

  return summoner, nil
}

// FetchSummonerByID - Devuelve la informacion del summoner por ID tipo URID
func FetchSummonerByID(sumUrid string) (entities.Summoner, error) {
  var response responses.Summoner
  var summoner entities.Summoner
  ID := urid.GetID(sumUrid)
  region := urid.GetRegion(sumUrid)

  url := createURL(region, "summoner/v3/summoners/" + ID)
  err := fetch(url, &response)

  if err != nil {
    return summoner, err
  }

  summoner = entities.Summoner{
    ID: urid.Generate(region, response.ID),
    ProfileIconID: response.ProfileIconID,
    Name: response.Name,
    SummonerLevel: response.SummonerLevel,
    AccountID: response.AccountID,
    RevisionDate: response.RevisionDate,
  }

  return summoner, nil
}
