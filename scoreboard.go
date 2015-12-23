package main

import (
	"log"
	"os"
	"time"

	"github.com/jcomo/scoreboard/mlb"
)

type Scoreboard struct {
	Client mlb.Client
	Logger *log.Logger
}

func NewScoreboard() *Scoreboard {
	return &Scoreboard{
		Client: mlb.DefaultHttpClient(),
		Logger: log.New(os.Stdout, "[SCOREBOARD] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (sb *Scoreboard) Get(day time.Time) ([]string, error) {
	games, err := sb.Client.FetchGames(day)
	if err != nil {
		sb.Logger.Println(err)
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
		sb.Logger.Println(err)
		return "", err
	}

	for _, g := range games {
		if g.HomeTeam() == team || g.AwayTeam() == team {
			return g.State(), nil
		}
	}

	return "No games for " + team, nil
}
