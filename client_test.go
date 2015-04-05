package main

import (
	"testing"
	"time"
)

func TestHttpClient(t *testing.T) {
	// Spot check using the data from 4/4/2015
	// There were 15 games played on that day
	// One game result was WSH 3 - 4 NYY
	d := time.Date(2015, 4, 4, 0, 0, 0, 0, time.UTC)

	c := DefaultHttpClient()
	games, err := c.FetchGames(d)

	assertNoError(t, err)
	assertEqual(t, 15, len(games))

	want := FinishedGame{
		home: newTeamStatus("WSH", 3),
		away: newTeamStatus("NYY", 4),
	}

	found := false
	for _, game := range games {
		if game.State() == want.State() {
			found = true
		}
	}

	if !found {
		t.Errorf("Did not find expected entry for game: %s", want.State())
	}
}
