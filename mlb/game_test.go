package mlb

import (
	"testing"
	"time"

	"github.com/jcomo/scoreboard/assert"
)

var upcomingGame = UpcomingGame{
	Home: "NYY",
	Away: "BOS",
	Time: time.Date(2015, 4, 4, 19, 5, 0, 0, time.Local),
}

func TestStateForUpcomingGame(t *testing.T) {
	assert.Equal(t, "BOS vs NYY 7:05PM", upcomingGame.State())
}

func TestHomeTeamForUpcomingGame(t *testing.T) {
	assert.Equal(t, "NYY", upcomingGame.HomeTeam())
}

func TestAwayTeamForUpcomingGame(t *testing.T) {
	assert.Equal(t, "BOS", upcomingGame.AwayTeam())
}

var inProgressGame = InProgressGame{
	Home:   NewTeamStatus("WSH", 5),
	Away:   NewTeamStatus("PHI", 2),
	Inning: TopInning(7),
}

func TestStateForGameInProgress(t *testing.T) {
	assert.Equal(t, "PHI 2 • 5 WSH ↑7", inProgressGame.State())
}

func TestHomeTeamForInProgressGame(t *testing.T) {
	assert.Equal(t, "WSH", inProgressGame.HomeTeam())
}

func TestAwayTeamForInProgressGame(t *testing.T) {
	assert.Equal(t, "PHI", inProgressGame.AwayTeam())
}

var finishedGame = FinishedGame{
	Home: NewTeamStatus("OAK", 8),
	Away: NewTeamStatus("LAA", 3),
}

func TestStateForFinishedGame(t *testing.T) {
	assert.Equal(t, "LAA 3 • 8 OAK F", finishedGame.State())
}

func TestHomeTeamForFinishedGame(t *testing.T) {
	assert.Equal(t, "OAK", finishedGame.HomeTeam())
}

func TestAwayTeamForFinishedGame(t *testing.T) {
	assert.Equal(t, "LAA", finishedGame.AwayTeam())
}

func TestInningString(t *testing.T) {
	bi := BottomInning(5)
	assert.Equal(t, "↓5", bi.String())

	ti := TopInning(4)
	assert.Equal(t, "↑4", ti.String())
}
