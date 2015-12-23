package mlb

import (
	"testing"
	"time"

	"github.com/jcomo/scoreboard/assert"
)

var testDate = time.Date(2015, 4, 4, 0, 0, 0, 0, time.Local)

// Spot check using the data from 4/4/2015. There were 15 games played
// on that day. One game result was WSH 3 - 4 NYY
func TestHttpClient(t *testing.T) {
	c := DefaultHttpClient()
	games, err := c.FetchGames(testDate)

	assert.NoError(t, err)
	assert.Equal(t, 15, len(games))

	want := FinishedGame{
		Home: NewTeamStatus("WSH", 3),
		Away: NewTeamStatus("NYY", 4),
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

func TestFixtureClient(t *testing.T) {
	c := DefaultFixtureClient()
	games, err := c.FetchGames(testDate)
	assert.NoError(t, err)

	want := []Game{
		FinishedGame{
			Home: NewTeamStatus("TB", 0),
			Away: NewTeamStatus("DET", 1),
		},
		FinishedGame{
			Home: NewTeamStatus("MIN", 2),
			Away: NewTeamStatus("BOS", 4),
		},
		FinishedGame{
			Home: NewTeamStatus("WSH", 3),
			Away: NewTeamStatus("NYY", 4),
		},
		FinishedGame{
			Home: NewTeamStatus("PHI", 6),
			Away: NewTeamStatus("PIT", 4),
		},
		FinishedGame{
			Home: NewTeamStatus("TOR", 9),
			Away: NewTeamStatus("CIN", 1),
		},
		FinishedGame{
			Home: NewTeamStatus("JAX", 14),
			Away: NewTeamStatus("MIA", 8),
		},
		FinishedGame{
			Home: NewTeamStatus("TEX", 4),
			Away: NewTeamStatus("NYM", 4),
		},
		FinishedGame{
			Home: NewTeamStatus("ATL", 5),
			Away: NewTeamStatus("BAL", 3),
		},
		FinishedGame{
			Home: NewTeamStatus("MIL", 4),
			Away: NewTeamStatus("CLE", 3),
		},
		FinishedGame{
			Home: NewTeamStatus("COL", 3),
			Away: NewTeamStatus("SEA", 6),
		},
		FinishedGame{
			Home: NewTeamStatus("HOU", 1),
			Away: NewTeamStatus("KC", 3),
		},
		FinishedGame{
			Home: NewTeamStatus("OAK", 1),
			Away: NewTeamStatus("SF", 2),
		},
		FinishedGame{
			Home: NewTeamStatus("ARI", 4),
			Away: NewTeamStatus("CHC", 2),
		},
		FinishedGame{
			Home: NewTeamStatus("SD", 8),
			Away: NewTeamStatus("MEX", 0),
		},
		InProgressGame{
			Home:   NewTeamStatus("LAD", 5),
			Away:   NewTeamStatus("LAA", 6),
			Inning: TopInning(8),
		},
	}

	assert.ArrayEqual(t, want, games)
}
