package main

import (
	"time"
)

func parseGameTime(gt string) time.Time {
	estLoc, _ := time.LoadLocation("America/New_York")
	t, _ := time.ParseInLocation("3:04PM", gt, estLoc)
	return t.Local()
}
