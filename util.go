package main

import (
	"strconv"
	"time"
)

func intFromStr(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	} else {
		return i
	}
}

func parseGameTime(gt string) time.Time {
	estLoc, _ := time.LoadLocation("America/New_York")
	t, _ := time.ParseInLocation("3:04PM", gt, estLoc)
	return t.Local()
}
