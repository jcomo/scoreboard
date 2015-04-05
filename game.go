package main

import (
	"fmt"
	"time"
)

type Game interface {
	State() string
	HomeTeam() string
	AwayTeam() string
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

func (g UpcomingGame) HomeTeam() string {
	return g.home
}

func (g UpcomingGame) AwayTeam() string {
	return g.away
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

func (g InProgressGame) HomeTeam() string {
	return g.home.teamName()
}

func (g InProgressGame) AwayTeam() string {
	return g.away.teamName()
}

type FinishedGame struct {
	home teamStatus
	away teamStatus
}

func (g FinishedGame) State() string {
	return fmt.Sprintf("%s %d - %d %s F",
		g.away.abbrev, g.away.score, g.home.score, g.home.abbrev)
}

func (g FinishedGame) HomeTeam() string {
	return g.home.teamName()
}

func (g FinishedGame) AwayTeam() string {
	return g.away.teamName()
}

type teamStatus struct {
	abbrev string
	score  int
}

func (ts teamStatus) teamName() string {
	return ts.abbrev
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
