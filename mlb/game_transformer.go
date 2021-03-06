package mlb

import (
	"encoding/json"
	"strconv"
	"time"
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
// an object.
func unmarshalGames(gd json.RawMessage) ([]RawGame, error) {
	var gs []RawGame

	e := json.Unmarshal(gd, &gs)
	if e != nil {
		// Only one game -- unmarshal as one game specifically
		var g RawGame

		e = json.Unmarshal(gd, &g)
		if e != nil {
			// No games at all
			return []RawGame{}, nil
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
	state := rg.GameState.State

	if state == "Final" || state == "Game Over" {
		return finishedGameFromRaw(rg)
	} else if state == "In Progress" {
		return inProgressGameFromRaw(rg)
	} else {
		return upcomingGameFromRaw(rg)
	}
}

func upcomingGameFromRaw(rg RawGame) Game {
	return UpcomingGame{
		Home: rg.HomeName,
		Away: rg.AwayName,
		Time: parseGameTime(rg.Time + rg.AmOrPm),
	}
}

func inProgressGameFromRaw(rg RawGame) Game {
	return InProgressGame{
		Home: homeTeamStatus(rg),
		Away: awayTeamStatus(rg),
		Inning: Inning{
			number: intFromStr(rg.GameState.Inning),
			top:    rg.GameState.InningState == "Top",
		},
	}
}

func finishedGameFromRaw(rg RawGame) Game {
	return FinishedGame{
		Home: homeTeamStatus(rg),
		Away: awayTeamStatus(rg),
	}
}

func homeTeamStatus(rg RawGame) TeamStatus {
	return NewTeamStatus(rg.HomeName, intFromStr(rg.LineScore.Runs.Home))
}

func awayTeamStatus(rg RawGame) TeamStatus {
	return NewTeamStatus(rg.AwayName, intFromStr(rg.LineScore.Runs.Away))
}

// MLB sends everything as a string. We want a simple helper function to
// parse an int but parse the int and return 0 on failure instead of an error.
func intFromStr(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	} else {
		return i
	}
}

// MLB sends game times in EST. We want to localize these to the user's
// machine's timezone.
func parseGameTime(gt string) time.Time {
	estLoc, _ := time.LoadLocation("America/New_York")
	t, _ := time.ParseInLocation("3:04PM", gt, estLoc)
	return t.Local()
}
