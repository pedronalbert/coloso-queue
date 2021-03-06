package utils

import "strings"
import "reflect"
import "encoding/json"

var regionAndPlatform = map[string]string{
	"BR":   "BR1",
	"EUNE": "EUN1",
	"EUW":  "EUW1",
	"JP":   "JP1",
	"KR":   "KR",
	"LAN":  "LA1",
	"LAS":  "LA2",
	"NA":   "NA1",
	"OCE":  "OC1",
	"RU":   "RU",
	"TR":   "TR1",
}

// RegionToPlatform - Devuelve la plataforma correspondiente a la region dada
func RegionToPlatform(region string) string {
	return regionAndPlatform[strings.ToUpper(region)]
}

// PlatformToRegion - Devuelve la region correspondiente a la plataforma dada
func PlatformToRegion(platformToFind string) string {
	platformToFind = strings.ToUpper(platformToFind)

	for _, platform := range regionAndPlatform {
		if platform == platformToFind {
			return platform
		}
	}

	return "NA"
}

// CompareStructs - Compare structs with given keys
func CompareStructs(sta interface{}, stb interface{}, keys []string) bool {
	var refValueA, refValueB reflect.Value
	var equals = true

	refValueA = reflect.ValueOf(sta)
	refValueB = reflect.ValueOf(stb)

	for _, key := range keys {
		if !reflect.DeepEqual(refValueA.FieldByName(key).Interface(), refValueB.FieldByName(key).Interface()) {
			equals = false
		}
	}
	return equals
}

// StructToString - Parse json struct to string
func StructToString(st interface{}) string {
	var resultBytes []byte

	resultBytes, _ = json.Marshal(st)

	return string(resultBytes)
}
