package main

import (
	"testing"
	"time"
)

func TestParseGameTime(t *testing.T) {
	easternLoc, _ := time.LoadLocation("America/New_York")

	localTime := time.Date(2015, 1, 1, 19, 5, 0, 0, easternLoc).Local()
	gameTime := parseGameTime("7:05PM")

	assertEqual(t, localTime.Hour(), gameTime.Hour())
	assertEqual(t, localTime.Minute(), gameTime.Minute())
}
