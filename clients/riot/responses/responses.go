package responses

// Summoner - Esquema de respuesta de un invocador
type Summoner struct {
  ID int `json:"id"`
  ProfileIconID int `json:"profileIconId"`
  Name string `json:"name"`
  SummonerLevel int `json:"summonerLevel"`
  AccountID int `json:"accountId"`
  RevisionDate int `json:"revisionDate"`
}

// RunesPages - Esquema de respuesta de runas
type RunesPages struct {
	SummonerID int `json:"summonerId"`
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
  ChampionLevel int `json:"championLevel"`
  ChampionPoints int `json:"championPoints"`
  ChampionID int `json:"championId"`
  ChampionPointsUntilNextLevel int `json:"championPointsUntilNextLevel"`
  ChampionPointsSinceLastLevel int `json:"championPointsSinceLastLevel"`
  LastPlayTime int `json:"lastPlayTime"`
}

// MasteriesPages = Esquema de respuesta de masteries
type MasteriesPages struct {
	SummonerID int `json:"summonerId"`
	Pages []struct {
		Current bool `json:"current"`
		Masteries []struct {
			ID int `json:"id"`
			Rank int `json:"rank"`
		} `json:"masteries"`
		ID int `json:"id"`
		Name string `json:"name"`
	} `json:"pages"`
}

// Game - Esquema de un juego
type Game struct {
	SeasonID int `json:"seasonId"`
	QueueID int `json:"queueId"`
	GameID int `json:"gameId"`
	ParticipantIdentities []struct {
		Player struct {
			CurrentPlatformID string `json:"currentPlatformId"`
			SummonerName string `json:"summonerName"`
			MatchHistoryURI string `json:"matchHistoryUri"`
			PlatformID string `json:"platformId"`
			CurrentAccountID int `json:"currentAccountId"`
			ProfileIcon int `json:"profileIcon"`
			SummonerID int `json:"summonerId"`
			AccountID int `json:"accountId"`
		} `json:"player"`
		ParticipantID int `json:"participantId"`
	} `json:"participantIdentities"`
	GameVersion string `json:"gameVersion"`
	PlatformID string `json:"platformId"`
	GameMode string `json:"gameMode"`
	MapID int `json:"mapId"`
	GameType string `json:"gameType"`
	Teams []struct {
		FirstDragon bool `json:"firstDragon"`
		FirstInhibitor bool `json:"firstInhibitor"`
		Bans []struct {
			PickTurn int `json:"pickTurn"`
			ChampionID int `json:"championId"`
		} `json:"bans"`
    BaronKills int `json:"baronKills"`
    FirstRiftHerald bool `json:"firstRiftHerald"`
    FirstBaron bool `json:"firstBaron"`
		RiftHeraldKills int `json:"riftHeraldKills"`
    FirstBlood bool `json:"firstBlood"`
    TeamID int `json:"teamId"`
    FirstTower bool `json:"firstTower"`
    VilemawKills int `json:"vilemawKills"`
    InhibitorKills int `json:"inhibitorKills"`
    TowerKills int `json:"towerKills"`
    DominionVictoryScore int `json:"dominionVictoryScore"`
		Win string `json:"win"`
    DragonKills int `json:"dragonKills"`
	} `json:"teams"`
	Participants []struct {
		Stats struct {
      PhysicalDamageDealt int `json:"physicalDamageDealt"`
      NeutralMinionsKilledTeamJungle int `json:"neutralMinionsKilledTeamJungle"`
      MagicDamageDealt int `json:"magicDamageDealt"`
			TotalPlayerScore int `json:"totalPlayerScore"`
      Deaths int `json:"deaths"`
      Win bool `json:"win"`
      NeutralMinionsKilledEnemyJungle int `json:"neutralMinionsKilledEnemyJungle"`
      LargestCriticalStrike int `json:"largestCriticalStrike"`
      TotalDamageDealt int `json:"totalDamageDealt"`
      MagicDamageDealtToChampions int `json:"magicDamageDealtToChampions"`
      AltarsCaptured int `json:"altarsCaptured"`
			VisionScore int `json:"visionScore"`
			UnrealKills int `json:"unrealKills"`
			ObjectivePlayerScore int `json:"objectivePlayerScore"`
			LargestMultiKill int `json:"largestMultiKill"`
			LargestKillingSpree int `json:"largestKillingSpree"`
			QuadraKills int `json:"quadraKills"`
			TotalTimeCrowdControlDealt int `json:"totalTimeCrowdControlDealt"`
			MagicalDamageTaken int `json:"magicalDamageTaken"`
			LongestTimeSpentLiving int `json:"longestTimeSpentLiving"`
			FirstTowerAssist bool `json:"firstTowerAssist"`
			GoldEarned int `json:"goldEarned"`
      Item0 int `json:"item0"`
      Item1 int `json:"item1"`
			Item2 int `json:"item2"`
			Item3 int `json:"item3"`
      Item4 int `json:"item4"`
      Item5 int `json:"item5"`
      Item6 int `json:"item6"`
      Kills int `json:"kills"`
			WardsPlaced int `json:"wardsPlaced"`
			TurretKills int `json:"turretKills"`
			TripleKills int `json:"tripleKills"`
			DamageSelfMitigated int `json:"damageSelfMitigated"`
			GoldSpent int `json:"goldSpent"`
			DoubleKills int `json:"doubleKills"`
			FirstInhibitorKill bool `json:"firstInhibitorKill"`
			TrueDamageTaken int `json:"trueDamageTaken"`
			FirstBloodAssist bool `json:"firstBloodAssist"`
			FirstBloodKill bool `json:"firstBloodKill"`
			Assists int `json:"assists"`
			TotalScoreRank int `json:"totalScoreRank"`
			NeutralMinionsKilled int `json:"neutralMinionsKilled"`
			CombatPlayerScore int `json:"combatPlayerScore"`
			VisionWardsBoughtInGame int `json:"visionWardsBoughtInGame"`
			DamageDealtToTurrets int `json:"damageDealtToTurrets"`
			PhysicalDamageDealtToChampions int `json:"physicalDamageDealtToChampions"`
			PentaKills int `json:"pentaKills"`
			TrueDamageDealt int `json:"trueDamageDealt"`
			TrueDamageDealtToChampions int `json:"trueDamageDealtToChampions"`
			ChampLevel int `json:"champLevel"`
			ParticipantID int `json:"participantId"`
			FirstInhibitorAssist bool `json:"firstInhibitorAssist"`
			WardsKilled int `json:"wardsKilled"`
			FirstTowerKill bool `json:"firstTowerKill"`
			TotalHeal int `json:"totalHeal"`
			TotalMinionsKilled int `json:"totalMinionsKilled"`
			DamageDealtToObjectives int `json:"damageDealtToObjectives"`
			SightWardsBoughtInGame int `json:"sightWardsBoughtInGame"`
			TotalDamageDealtToChampions int `json:"totalDamageDealtToChampions"`
			TotalUnitsHealed int `json:"totalUnitsHealed"`
			InhibitorKills int `json:"inhibitorKills"`
			TotalDamageTaken int `json:"totalDamageTaken"`
			KillingSprees int `json:"killingSprees"`
			TimeCCingOthers int `json:"timeCCingOthers"`
			PhysicalDamageTaken int `json:"physicalDamageTaken"`
			TeamObjective int `json:"teamObjective"`
			NodeNeutralizeAssist int `json:"nodeNeutralizeAssist"`
			NodeNeutralize int `json:"nodeNeutralize"`
			NodeCaptureAssist int `json:"nodeCaptureAssist"`
			AltarsNeutralized int `json:"altarsNeutralized"`
			NodeCapture int `json:"nodeCapture"`
		} `json:"stats"`
    ParticipantID int `json:"participantId"`
    Runes []struct {
      RuneID int `json:"runeId"`
      Rank int `json:"rank"`
    } `json:"runes"`
		Spell1ID int `json:"spell1Id"`
    Spell2ID int `json:"spell2Id"`
		HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
		Masteries []struct {
			MasteryID int `json:"masteryId"`
			Rank int `json:"rank"`
		} `json:"masteries"`
		TeamID int `json:"teamId"`
		Timeline struct {
			Lane string `json:"lane"`
			ParticipantID int `json:"participantId"`
			CsDiffPerMinDeltas struct {
				Two030 float64 `json:"20-30"`
				Zero10 float64 `json:"0-10"`
				One020 int `json:"10-20"`
			} `json:"csDiffPerMinDeltas"`
			GoldPerMinDeltas struct {
				Two030 float64 `json:"20-30"`
				Zero10 float64 `json:"0-10"`
				One020 float64 `json:"10-20"`
			} `json:"goldPerMinDeltas"`
			XpDiffPerMinDeltas struct {
				Two030 float64 `json:"20-30"`
				Zero10 float64 `json:"0-10"`
				One020 float64 `json:"10-20"`
			} `json:"xpDiffPerMinDeltas"`
			CreepsPerMinDeltas struct {
				Two030 float64 `json:"20-30"`
				Zero10 float64 `json:"0-10"`
				One020 float64 `json:"10-20"`
			} `json:"creepsPerMinDeltas"`
			XpPerMinDeltas struct {
				Two030 float64 `json:"20-30"`
				Zero10 float64 `json:"0-10"`
				One020 float64 `json:"10-20"`
			} `json:"xpPerMinDeltas"`
			Role string `json:"role"`
			DamageTakenDiffPerMinDeltas struct {
				Two030 int `json:"20-30"`
				Zero10 float64 `json:"0-10"`
				One020 float64 `json:"10-20"`
			} `json:"damageTakenDiffPerMinDeltas"`
			DamageTakenPerMinDeltas struct {
				Two030 float64 `json:"20-30"`
				Zero10 float64 `json:"0-10"`
				One020 int `json:"10-20"`
			} `json:"damageTakenPerMinDeltas"`
		} `json:"timeline"`
		ChampionID int `json:"championId"`
	} `json:"participants"`
	GameDuration int `json:"gameDuration"`
	GameCreation int `json:"gameCreation"`
}

type participantFrame struct {
  ParticipantID int `json:"participantId"`
  Position struct {
    X int `json:"x"`
    Y int `json:"y"`
  } `json:"position"`
  CurrentGold int `json:"currentGold"`
  TotalGold int `json:"totalGold"`
  Level int `json:"level"`
  Xp int `json:"xp"`
  MinionsKilled int `json:"minionsKilled"`
  JungleMinionsKilled int `json:"jungleMinionsKilled"`
  DominionScore int `json:"dominionScore"`
  TeamScore int `json:"teamScore"`
}

// GameTimelines - Timelines de un game
type GameTimelines struct {
	Frames []struct {
		ParticipantFrames struct {
			Num1 participantFrame `json:"1"`
			Num2 participantFrame `json:"2"`
			Num3 participantFrame `json:"3"`
			Num4 participantFrame `json:"4"`
			Num5 participantFrame `json:"5"`
			Num6 participantFrame `json:"6"`
			Num7 participantFrame `json:"7"`
			Num8 participantFrame `json:"8"`
			Num9 participantFrame `json:"9"`
			Num10 participantFrame `json:"10"`
		} `json:"participantFrames"`
		Events []interface{} `json:"events"`
		Timestamp int `json:"timestamp"`
	} `json:"frames"`
	FrameInterval int `json:"frameInterval"`
}

// GamesList - Esquema de respuesta de matchlists
type GamesList struct {
	Matches []struct {
		PlatformID string `json:"platformId"`
		GameID int `json:"gameId"`
		Champion int `json:"champion"`
		Queue int `json:"queue"`
		Season int `json:"season"`
		Timestamp int `json:"timestamp"`
		Role string `json:"role"`
		Lane string `json:"lane"`
	} `json:"matches"`
	StartIndex int `json:"startIndex"`
	EndIndex int `json:"endIndex"`
	TotalGames int `json:"totalGames"`
}

// LeaguePostition - Esquema de posicion en la liga
type LeaguePostition struct {
	QueueType string `json:"queueType"`
	HotStreak bool `json:"hotStreak"`
	Wins int `json:"wins"`
	Veteran bool `json:"veteran"`
	Losses int `json:"losses"`
	PlayerOrTeamID string `json:"playerOrTeamId"`
	Tier string `json:"tier"`
	PlayerOrTeamName string `json:"playerOrTeamName"`
	Inactive bool `json:"inactive"`
	Rank string `json:"rank"`
	FreshBlood bool `json:"freshBlood"`
	LeagueName string `json:"leagueName"`
	LeaguePoints int `json:"leaguePoints"`
  MiniSeries struct {
    Wins int `json:"wins"`
    Losses int `json:"losses"`
    Target int `json:"target"`
    Progress string `json:"progress"`
  } `json:"miniSeries"`
}
