package main

import (
	"testing"
	"time"
)

var upcomingGame = UpcomingGame{
	home: "NYY",
	away: "BOS",
	time: time.Date(2015, 4, 4, 19, 5, 0, 0, time.Local),
}

var inProgressGame = InProgressGame{
	home:   newTeamStatus("WSH", 5),
	away:   newTeamStatus("PHI", 2),
	inning: topInning(7),
}

var finishedGame = FinishedGame{
	home: newTeamStatus("OAK", 8),
	away: newTeamStatus("LAA", 3),
}

func TestStateForUpcomingGame(t *testing.T) {
	assertEqual(t, "BOS vs NYY 7:05PM", upcomingGame.State())
}

func TestHomeTeamForUpcomingGame(t *testing.T) {
	assertEqual(t, "NYY", upcomingGame.HomeTeam())
}

func TestAwayTeamForUpcomingGame(t *testing.T) {
	assertEqual(t, "BOS", upcomingGame.AwayTeam())
}

func TestStateForGameInProgress(t *testing.T) {
	assertEqual(t, "PHI 2 - 5 WSH T7", inProgressGame.State())
}

func TestHomeTeamForInProgressGame(t *testing.T) {
	assertEqual(t, "WSH", inProgressGame.HomeTeam())
}

func TestAwayTeamForInProgressGame(t *testing.T) {
	assertEqual(t, "PHI", inProgressGame.AwayTeam())
}

func TestStateForFinishedGame(t *testing.T) {
	assertEqual(t, "LAA 3 - 8 OAK F", finishedGame.State())
}

func TestHomeTeamForFinishedGame(t *testing.T) {
	assertEqual(t, "OAK", finishedGame.HomeTeam())
}

func TestAwayTeamForFinishedGame(t *testing.T) {
	assertEqual(t, "LAA", finishedGame.AwayTeam())
}

func TestInningString(t *testing.T) {
	bi := bottomInning(5)
	assertEqual(t, "B5", bi.String())

	ti := topInning(4)
	assertEqual(t, "T4", ti.String())
}
