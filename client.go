package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Client interface {
	FetchGames(time.Time) ([]Game, error)
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

	return gamesFromReader(resp.Body)
}

func (c *HttpClient) urlForDay(day time.Time) string {
	pathForDay := fmt.Sprintf("/year_%d/month_%02d/day_%02d",
		day.Year(), day.Month(), day.Day())
	return c.BaseURL + pathForDay + "/master_scoreboard.json"
}

type FixtureClient struct {
	BaseDir string
}

func DefaultFixtureClient() *FixtureClient {
	return &FixtureClient{
		BaseDir: "fixtures",
	}
}

func (c *FixtureClient) FetchGames(day time.Time) ([]Game, error) {
	f, err := os.Open(c.filePathForDay(day))
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return gamesFromReader(f)
}

func (c *FixtureClient) filePathForDay(day time.Time) string {
	pathForDay := fmt.Sprintf("%d_%02d_%02d.json",
		day.Year(), day.Month(), day.Day())

	return filepath.Join(c.BaseDir, pathForDay)
}

func gamesFromReader(f io.Reader) ([]Game, error) {
	var scoreboard rawScoreboard
	err := json.NewDecoder(f).Decode(&scoreboard)
	if err != nil {
		return nil, err
	}

	rgs, err := unmarshalGames(scoreboard.Data.Games.Game)
	if err != nil {
		return nil, err
	}

	return GamesFromRaw(rgs), nil
}
