package responses

// Summoner - Esquema de respuesta de un invocador
type Summoner struct {
  ID int64 `json:"id"`
  ProfileIconID int64 `json:"profileIconId"`
  Name string `json:"name"`
  SummonerLevel int64 `json:"summonerLevel"`
  AccountID int64 `json:"accountId"`
  RevisionDate int64 `json:"revisionDate"`
}

// RunesPages - Esquema de respuesta de runas
type RunesPages struct {
	SummonerID int64 `json:"summonerId"`
	Pages []struct {
    ID int `json:"id"`
    Name string `json:"name"`
		Current bool `json:"current"`
		Slots []struct {
			RuneSlotID int `json:"runeSlotId"`
			RuneID int `json:"runeId"`
		} `json:"slots"`
	} `json:"pages"`
}

// ChampionMastery - Esquema de la respuesta de champion-mastery
type ChampionMastery struct {
  ChestGranted bool `json:"ChestGranted"`
  ChampionLevel int64 `json:"championLevel"`
  ChampionPoints int64 `json:"championPoints"`
  ChampionID int64 `json:"championId"`
  ChampionPointsUntilNextLevel int64 `json:"championPointsUntilNextLevel"`
  ChampionPointsSinceLastLevel int64 `json:"championPointsSinceLastLevel"`
  LastPlayTime int64 `json:"lastPlayTime"`
}
