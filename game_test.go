package main

import (
    "testing"
)

func TestStateForUpcomingGame(t *testing.T) {
    g := Game{
        homeTeam: "NYY",
        awayTeam: "BOS",
        time: "7:05PM",
        state: "Upcoming",
    }

    want := "BOS vs NYY 7:05PM"
    got := g.State()

    assertEqual(t, want, got)
}

func TestStateForGameInProgress(t *testing.T) {
    g := Game{
        homeTeam: "NYY",
        awayTeam: "BOS",
        homeScore: 5,
        awayScore: 2,
        inning: 7,
        inningState: "T",
        state: "In Progress",
    }

    want := "BOS 2 - 5 NYY T7"
    got := g.State()

    assertEqual(t, want, got)
}

func TestStateForFinishedGame(t *testing.T) {
    g := Game{
        homeTeam: "NYY",
        awayTeam: "BOS",
        homeScore: 8,
        awayScore: 3,
        state: "Final",
    }

    want := "BOS 3 - 8 NYY F"
    got := g.State()

    assertEqual(t, want, got)
}