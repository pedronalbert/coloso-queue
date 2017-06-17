package urid

import (
  "regexp"
  "strings"
)

// GetID - Devuelve el ID de una URID
func GetID(urid string) string {
  regex, _ := regexp.Compile("[0-9]+")

  id := regex.FindString(urid)

  return id
}

// GetRegion - Devuelve la region de una urid
func GetRegion(urid string) string {
  regex, _ := regexp.Compile("[A-Z]+")

  region := regex.FindString(urid)

  return region
}

// Generate - Devuelve un nuevo URID
func Generate(region string, ID string) string {
  region = strings.ToUpper(region)

  return region + "_" + ID
}
