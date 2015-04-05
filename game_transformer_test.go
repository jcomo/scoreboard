package main

import (
	"encoding/json"
	"testing"
)

func TestRawGameUnmarshalling(t *testing.T) {
	rgData := []byte(`
    {
        "time": "7:05",
        "ampm": "PM",
        "home_name_abbrev": "NYY",
        "away_name_abbrev": "BOS",
        "linescore": {
            "r": {
                "home": "10",
                "away": "1"
            }
        },
        "status": {
            "status": "In Progress",
            "inning": "8",
            "inning_state": "Top"
        }
    }`)

	got := RawGame{}
	err := json.Unmarshal(rgData, &got)
	if err != nil {
		panic("Could not unmarshal raw game data!")
	}

	want := RawGame{
		Time:     "7:05",
		AmOrPm:   "PM",
		HomeName: "NYY",
		AwayName: "BOS",
		LineScore: lineScore{
			Runs: runs{
				Home: "10",
				Away: "1",
			},
		},
		GameState: gameState{
			State:       "In Progress",
			Inning:      "8",
			InningState: "Top",
		},
	}

	assertEqual(t, want, got)
}

func TestTransformRawUpcomingGame(t *testing.T) {
	rg := RawGame{
		Time:     "7:05",
		AmOrPm:   "PM",
		HomeName: "NYY",
		AwayName: "BOS",
		GameState: gameState{
			State: "Preview",
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
		HomeName: "NYY",
		AwayName: "BOS",
		LineScore: lineScore{
			Runs: runs{
				Home: "6",
				Away: "5",
			},
		},
		GameState: gameState{
			State:       "In Progress",
			Inning:      "7",
			InningState: "Top",
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
		HomeName: "NYY",
		AwayName: "BOS",
		LineScore: lineScore{
			Runs: runs{
				Home: "6",
				Away: "5",
			},
		},
		GameState: gameState{
			State:       "In Progress",
			Inning:      "7",
			InningState: "Bottom",
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
		HomeName: "NYY",
		AwayName: "BOS",
		LineScore: lineScore{
			Runs: runs{
				Home: "6",
				Away: "5",
			},
		},
		GameState: gameState{
			State: "Final",
		},
	}

	got := GameFromRaw(rg)
	want := FinishedGame{
		home: newTeamStatus("NYY", 6),
		away: newTeamStatus("BOS", 5),
	}

	assertEqual(t, want.State(), got.State())
}
