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
