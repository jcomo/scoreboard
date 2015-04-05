package main

import (
	"encoding/json"
)

type rawScoreboard struct {
	Data scoreboardData
}

type scoreboardData struct {
	Games scoreboardGames
}

type scoreboardGames struct {
	Game json.RawMessage
}

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

// Need a special helper to unmarshal games because if there is a list of
// games, it is returned as an array, but if there is only one game it is
// and object.
func unmarshalGames(gd json.RawMessage) ([]RawGame, error) {
	var gs []RawGame

	e := json.Unmarshal(gd, &gs)
	if e != nil {
		// Only one game -- unmarshal as one game specifically
		var g RawGame

		e = json.Unmarshal(gd, &g)
		if e != nil {
			return nil, e
		}

		gs = make([]RawGame, 1)
		gs[0] = g
	}

	return gs, nil
}

func GamesFromRaw(rgs []RawGame) []Game {
	gs := make([]Game, len(rgs), len(rgs))
	for i, rg := range rgs {
		gs[i] = GameFromRaw(rg)
	}

	return gs
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
