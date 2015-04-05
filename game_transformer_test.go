package main

import (
    "testing"
)

func TestTransformRawUpcomingGame(t *testing.T) {
    rg := RawGame{
        time: "7:05",
        amOrPm: "PM",
        homeName: "NYY",
        awayName: "BOS",
        gameState: gameState{
            state: "Preview",
        },
    }

    got := GameFromRaw(rg)
    want := UpcomingGame{
        home: "NYY",
        away: "BOS",
        time: makeGameTime("7:05PM"),
    }

    assertEqual(t, want.State(), got.State())
}

func TestTransformRawInProgressGameTopInning(t *testing.T) {
    rg := RawGame{
        homeName: "NYY",
        awayName: "BOS",
        lineScore: lineScore{
            runs: runs{
                home: 6,
                away: 5,
            },
        },
        gameState: gameState{
            state: "In Progress",
            inning: 7,
            inningState: "Top",
        },
    }

    got := GameFromRaw(rg)
    want := InProgressGame{
        home:   newTeamStatus("NYY", 6),
        away:   newTeamStatus("BOS", 5),
        inning: topInning(7),
    }

    assertEqual(t, want.State(), got.State())
}

func TestTransformRawInProgressGameBottomInning(t *testing.T) {
    rg := RawGame{
        homeName: "NYY",
        awayName: "BOS",
        lineScore: lineScore{
            runs: runs{
                home: 6,
                away: 5,
            },
        },
        gameState: gameState{
            state: "In Progress",
            inning: 7,
            inningState: "Bottom",
        },
    }

    got := GameFromRaw(rg)
    want := InProgressGame{
        home:   newTeamStatus("NYY", 6),
        away:   newTeamStatus("BOS", 5),
        inning: bottomInning(7),
    }

    assertEqual(t, want.State(), got.State())
}

func TestTransformRawFinishedGame(t *testing.T) {
    rg := RawGame{
        homeName: "NYY",
        awayName: "BOS",
        lineScore: lineScore{
            runs: runs{
                home: 6,
                away: 5,
            },
        },
        gameState: gameState{
            state: "Final",
        },
    }

    got := GameFromRaw(rg)
    want := FinishedGame{
        home:   newTeamStatus("NYY", 6),
        away:   newTeamStatus("BOS", 5),
    }

    assertEqual(t, want.State(), got.State())
}