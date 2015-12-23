package main

import (
	"testing"
	"time"

	"github.com/jcomo/scoreboard/assert"
	"github.com/jcomo/scoreboard/mlb"
)

var testDate = time.Date(2015, 4, 4, 0, 0, 0, 0, time.Local)

type fakeClient struct {
}

func (c *fakeClient) FetchGames(day time.Time) ([]mlb.Game, error) {
	games := []mlb.Game{
		mlb.FinishedGame{
			Home: mlb.NewTeamStatus("TB", 0),
			Away: mlb.NewTeamStatus("DET", 1),
		},
		mlb.InProgressGame{
			Home:   mlb.NewTeamStatus("LAD", 5),
			Away:   mlb.NewTeamStatus("LAA", 6),
			Inning: mlb.TopInning(8),
		},
	}

	return games, nil
}

func TestGet(t *testing.T) {
	want := []string{
		"DET 1 • 0 TB F",
		"LAA 6 • 5 LAD ↑8",
	}

	sb := NewScoreboard(&fakeClient{})

	got, err := sb.Get(testDate)

	assert.NoError(t, err)
	assert.Equal(t, len(want), len(got))
	for i, g := range got {
		assert.Equal(t, want[i], g)
	}
}

func TestGetTeam(t *testing.T) {
	sb := NewScoreboard(&fakeClient{})

	want := "DET 1 • 0 TB F"
	got, err := sb.GetTeam(testDate, "DET")
	assert.NoError(t, err)
	assert.Equal(t, want, got)

	want = "DET 1 • 0 TB F"
	got, err = sb.GetTeam(testDate, "TB")
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestGetTeamForNoTeamFound(t *testing.T) {
	sb := NewScoreboard(&fakeClient{})

	want := "No games for BAD"
	got, err := sb.GetTeam(testDate, "BAD")

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
