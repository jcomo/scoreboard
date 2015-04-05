package main

import (
    "testing"
)

func assertEqual(t *testing.T, want, got interface{}) {
    if got != want {
        t.Errorf("Expected %v, got %v", want, got)
    }
}

// TODO: flesh out this test
func TestGet(t *testing.T) {
    sb := NewScoreboard()
    got := sb.Get()

    want := "NYY vs BOS 7:05PM"
    assertEqual(t, want, got)
}

// TODO: flesh out this test
func TestGetTeam(t *testing.T) {
    sb := NewScoreboard()
    got := sb.GetTeam("NYY")

    want := "WSH 3 - 4 NYY F"
    assertEqual(t, want, got)
}