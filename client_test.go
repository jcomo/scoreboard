package main

import (
	"testing"
)

func TestHttpClient(t *testing.T) {
	assertGamesOnTestDate(t, DefaultHttpClient())
}

func TestFixtureClient(t *testing.T) {
	assertGamesOnTestDate(t, DefaultFixtureClient())
}

func assertGamesOnTestDate(t *testing.T, c Client) {
	// Spot check using the data from 4/4/2015
	// There were 15 games played on that day
	// One game result was WSH 3 - 4 NYY
	games, err := c.FetchGames(testDate)

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
