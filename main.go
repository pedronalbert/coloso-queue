package main

import (
  "fmt"
  "coloso-queue/clients/riot"
)
func main() {
  summoner, err := riot.FetchSummonerByID("LAN_754512")

  if err == nil {
    fmt.Printf("%+v\n", summoner)
  }

  fmt.Printf("%s", err)
}
