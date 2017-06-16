package utils

import "strings"

var regionAndPlatform = [string]string{
  "BR": "BR1",
  "EUNE": "EUN1",
  "EUW": "EUW1",
  "JP": "JP1",
  "KR": "KR",
  "LAN": "LA1",
  "LAS": "LA2",
  "NA": "NA1",
  "OCE": "OC1",
  "RU": "RU",
  "TR": "TR1",
}

// RegionToPlatform - Devuelve la plataforma correspondiente a la region dada
func RegionToPlatform(string region) string {
  return regionAndPlatform[strings.ToUpper(region)]
}

// PlatformToRegion - Devuelve la region correspondiente a la plataforma dada
func PlatformToRegion(string platformToFind) string {
  platformToFind = strings.ToUpper(platformToFind)

  for _, platform := range regionAndPlatform {
    if platform == platformToFind {
      return platform
    }
  }
}
