package cache

import "errors"
import "coloso-queue/clients/mysql"
import "coloso-queue/models"
import "github.com/op/go-logging"

var log = logging.MustGetLogger("queue")

// ErrNotFound - Entity not found on DB
var ErrNotFound = errors.New("Entity not found on DB")

// SaveSummoner - save summoner in cache
func SaveSummoner(newSum models.Summoner) models.Summoner {
  var sumInDb models.Summoner

  mysql.Client.First(&sumInDb)

  if sumInDb.ID == "" {
    log.Debugf("Summoner ID: %s not found in cache, creating new", newSum.ID)

    mysql.Client.Create(&newSum)

    return newSum
  }

  log.Debugf("Summoner ID: %s already exist in cache, updating", newSum.ID)

  mysql.Client.Model(&sumInDb).Updates(newSum)

  return sumInDb
}

// FindSummonerByID - find summone rin cache by URID
func FindSummonerByID(sumID string) (models.Summoner, error) {
  var sumFound models.Summoner
  var err error

  mysql.Client.Where("id = ?", sumID).First(&sumFound)

  if sumFound.ID != sumID {
    err = ErrNotFound
  }

  return sumFound, err
}
