package main

type RawGame struct {
	Time      string
	AmOrPm    string    `json:"ampm"`
	HomeName  string    `json:"home_name_abbrev"`
	AwayName  string    `json:"away_name_abbrev"`
	LineScore lineScore `json:"linescore"`
	GameState gameState `json:"status"`
}

type lineScore struct {
	Runs runs `json:"r"`
}

type runs struct {
	Home string
	Away string
}

type gameState struct {
	State       string `json:"status"`
	Inning      string
	InningState string `json:"inning_state"`
}

func GameFromRaw(rg RawGame) Game {
	if rg.GameState.State == "Preview" {
		return upcomingGameFromRaw(rg)
	} else if rg.GameState.State == "Final" {
		return finishedGameFromRaw(rg)
	} else {
		return inProgressGameFromRaw(rg)
	}
}

func upcomingGameFromRaw(rg RawGame) Game {
	return UpcomingGame{
		home: rg.HomeName,
		away: rg.AwayName,
		time: parseGameTime(rg.Time + rg.AmOrPm),
	}
}

func inProgressGameFromRaw(rg RawGame) Game {
	return InProgressGame{
		home: homeTeamStatus(rg),
		away: awayTeamStatus(rg),
		inning: inning{
			number: intFromStr(rg.GameState.Inning),
			top:    rg.GameState.InningState == "Top",
		},
	}
}

func finishedGameFromRaw(rg RawGame) Game {
	return FinishedGame{
		home: homeTeamStatus(rg),
		away: awayTeamStatus(rg),
	}
}

func homeTeamStatus(rg RawGame) teamStatus {
	return newTeamStatus(rg.HomeName, intFromStr(rg.LineScore.Runs.Home))
}

func awayTeamStatus(rg RawGame) teamStatus {
	return newTeamStatus(rg.AwayName, intFromStr(rg.LineScore.Runs.Away))
}
