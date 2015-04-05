package main

import (
	"fmt"
	"time"
)

type Game interface {
	State() string
}

type UpcomingGame struct {
	home string
	away string
	time time.Time
}

func (g UpcomingGame) State() string {
	return fmt.Sprintf("%s vs %s %s", g.away, g.home,
		g.time.Local().Format("3:04PM"))
}

type InProgressGame struct {
	home   teamStatus
	away   teamStatus
	inning inning
}

func (g InProgressGame) State() string {
	return fmt.Sprintf("%s %d - %d %s %s",
		g.away.abbrev, g.away.score, g.home.score, g.home.abbrev, g.inning)
}

type FinishedGame struct {
	home teamStatus
	away teamStatus
}

func (g FinishedGame) State() string {
	return fmt.Sprintf("%s %d - %d %s F",
		g.away.abbrev, g.away.score, g.home.score, g.home.abbrev)
}

type teamStatus struct {
	abbrev string
	score  int
}

func newTeamStatus(abbrev string, score int) teamStatus {
	return teamStatus{
		abbrev: abbrev,
		score:  score,
	}
}

type inning struct {
	number int
	top    bool
}

func bottomInning(number int) inning {
	return inning{
		number: number,
		top:    false,
	}
}

func topInning(number int) inning {
	return inning{
		number: number,
		top:    true,
	}
}

func (i inning) String() string {
	if i.top {
		return fmt.Sprintf("T%d", i.number)
	} else {
		return fmt.Sprintf("B%d", i.number)
	}
}
