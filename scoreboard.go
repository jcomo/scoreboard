package main

import (
	"time"
)

type Scoreboard struct {
	Client Client
}

func NewScoreboard(c Client) *Scoreboard {
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

func (sb *Scoreboard) GetTeam(team string) string {
	return "WSH 3 - 4 NYY F"
}
