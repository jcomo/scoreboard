package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

type HttpClient struct {
	BaseURL string
}

func DefaultHttpClient() *HttpClient {
	return &HttpClient{
		BaseURL: "http://mlb.mlb.com/gdcross/components/game/mlb",
	}
}

func (c *HttpClient) FetchGames(day time.Time) ([]Game, error) {
	resp, err := http.Get(c.urlForDay(day))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var scoreboard rawScoreboard
	err = json.NewDecoder(resp.Body).Decode(&scoreboard)
	if err != nil {
		return nil, err
	}

	rgs, err := unmarshalGames(scoreboard.Data.Games.Game)
	if err != nil {
		return nil, err
	}

	return GamesFromRaw(rgs), nil
}

func (c *HttpClient) urlForDay(day time.Time) string {
	pathForDay := fmt.Sprintf("/year_%d/month_%02d/day_%02d",
		day.Year(), day.Month(), day.Day())
	return c.BaseURL + pathForDay + "/master_scoreboard.json"
}
