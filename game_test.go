package main

import (
	"testing"
	"time"
)

func TestStateForUpcomingGame(t *testing.T) {
	g := UpcomingGame{
		home: "NYY",
		away: "BOS",
		time: time.Date(2015, time.January, 1, 19, 5, 0, 0, time.Local),
	}

	want := "BOS vs NYY 7:05PM"
	got := g.State()

	assertEqual(t, want, got)
}

func TestStateForGameInProgress(t *testing.T) {
	g := InProgressGame{
		home:   newTeamStatus("NYY", 5),
		away:   newTeamStatus("BOS", 2),
		inning: topInning(7),
	}

	want := "BOS 2 - 5 NYY T7"
	got := g.State()

	assertEqual(t, want, got)
}

func TestStateForFinishedGame(t *testing.T) {
	g := FinishedGame{
		home: newTeamStatus("NYY", 8),
		away: newTeamStatus("BOS", 3),
	}

	want := "BOS 3 - 8 NYY F"
	got := g.State()

	assertEqual(t, want, got)
}

func TestInningString(t *testing.T) {
	bi := bottomInning(5)
	assertEqual(t, "B5", bi.String())

	ti := topInning(4)
	assertEqual(t, "T4", ti.String())
}
