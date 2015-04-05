package main

import (
    "fmt"
    "time"
    "net/http"
    "encoding/json"
)

type rawScoreboard struct {
    Data scoreboardData
}

type scoreboardData struct {
    Games scoreboardGames
}

type scoreboardGames struct {
    Game []RawGame
}

type HttpClient struct {
    BaseURL string
}

func DefaultHttpClient() *HttpClient {
    return &HttpClient{
        BaseURL: "http://mlb.mlb.com/gdcross/components/game/mlb",
    }
}

func (c *HttpClient) FetchGames(day time.Time) []Game {
    resp, err := http.Get(c.urlForDay(day))
    if err != nil {
        // TODO: handle this instead
        panic("Could not reach mlb server!")
    }

    defer resp.Body.Close()

    var scoreboard rawScoreboard
    err = json.NewDecoder(resp.Body).Decode(&scoreboard)
    if err != nil {
        // TODO: handle this instead
        panic("Error fetching games!")
    }

    return GamesFromRaw(scoreboard.Data.Games.Game)
}

func (c *HttpClient) urlForDay(day time.Time) string {
    pathForDay := fmt.Sprintf("/year_%d/month_%02d/day_%02d",
                              day.Year(), day.Month(), day.Day())
    return c.BaseURL + pathForDay + "/master_scoreboard.json"
}