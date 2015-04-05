package main

import (
	"testing"
	"time"
)

func makeGameTime(gt string) time.Time {
	easternLoc, _ := time.LoadLocation("America/New_York")
	t, _ := time.ParseInLocation("3:04PM", gt, easternLoc)
	return t
}

func TestParseGameTime(t *testing.T) {
	localTime := makeGameTime("7:05PM").Local()
	gameTime := parseGameTime("7:05PM")

	assertEqual(t, localTime.Hour(), gameTime.Hour())
	assertEqual(t, localTime.Minute(), gameTime.Minute())
}

func TestIntFromString(t *testing.T) {
	assertEqual(t, 0, intFromStr(""))
	assertEqual(t, 0, intFromStr("bogus"))
	assertEqual(t, 8, intFromStr("8"))
}
