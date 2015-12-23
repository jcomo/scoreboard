package main

import (
	"time"

	"github.com/jcomo/scoreboard/mlb"
)

type Scoreboard struct {
	Client mlb.Client
}

// TODO: this shouldnt take a client
func NewScoreboard(c mlb.Client) *Scoreboard {
	return &Scoreboard{
		Client: c,
	}
}

func (sb *Scoreboard) Get(day time.Time) ([]string, error) {
	games, err := sb.Client.FetchGames(day)
	if err != nil {
		return nil, err
	}

	gameStates := make([]string, len(games))
	for i, g := range games {
		gameStates[i] = g.State()
	}

	return gameStates, nil
}

func (sb *Scoreboard) GetTeam(day time.Time, team string) (string, error) {
	games, err := sb.Client.FetchGames(day)
	if err != nil {
		return "", err
	}

	for _, g := range games {
		if g.HomeTeam() == team || g.AwayTeam() == team {
			return g.State(), nil
		}
	}

	return "No games for " + team, nil
}
