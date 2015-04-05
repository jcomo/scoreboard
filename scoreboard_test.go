package main

import (
	"testing"
	"time"
)

var testDate = time.Date(2015, 4, 4, 0, 0, 0, 0, time.Local)

func assertEqual(t *testing.T, want, got interface{}) {
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGet(t *testing.T) {
	want := []string{
		"DET 1 - 0 TB F",
		"BOS 4 - 2 MIN F",
		"NYY 4 - 3 WSH F",
		"PIT 4 - 6 PHI F",
		"CIN 1 - 9 TOR F",
		"MIA 8 - 14 JAX F",
		"NYM 4 - 4 TEX F",
		"BAL 3 - 5 ATL F",
		"CLE 3 - 4 MIL F",
		"SEA 6 - 3 COL F",
		"KC 3 - 1 HOU F",
		"SF 2 - 1 OAK F",
		"CHC 2 - 4 ARI F",
		"MEX 0 - 8 SD F",
		"LAA 6 - 5 LAD T8",
	}

	sb := NewScoreboard(DefaultFixtureClient())

	got, err := sb.Get(testDate)

	assertNoError(t, err)
	assertEqual(t, len(want), len(got))
	for i, g := range got {
		assertEqual(t, want[i], g)
	}
}

func TestGetTeam(t *testing.T) {
	sb := NewScoreboard(DefaultFixtureClient())
	got := sb.GetTeam("NYY")

	want := "WSH 3 - 4 NYY F"
	assertEqual(t, want, got)
}
