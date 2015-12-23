package mlb

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
	Home string
	Away string
	Time time.Time
}

func (g UpcomingGame) State() string {
	return fmt.Sprintf("%s vs %s %s", g.Away, g.Home,
		g.Time.Local().Format("3:04PM"))
}

func (g UpcomingGame) HomeTeam() string {
	return g.Home
}

func (g UpcomingGame) AwayTeam() string {
	return g.Away
}

type InProgressGame struct {
	Home   TeamStatus
	Away   TeamStatus
	Inning Inning
}

func (g InProgressGame) State() string {
	return fmt.Sprintf("%s %d • %d %s %s",
		g.Away.Abbrev, g.Away.Score, g.Home.Score, g.Home.Abbrev, g.Inning)
}

func (g InProgressGame) HomeTeam() string {
	return g.Home.teamName()
}

func (g InProgressGame) AwayTeam() string {
	return g.Away.teamName()
}

type FinishedGame struct {
	Home TeamStatus
	Away TeamStatus
}

func (g FinishedGame) State() string {
	return fmt.Sprintf("%s %d • %d %s F",
		g.Away.Abbrev, g.Away.Score, g.Home.Score, g.Home.Abbrev)
}

func (g FinishedGame) HomeTeam() string {
	return g.Home.teamName()
}

func (g FinishedGame) AwayTeam() string {
	return g.Away.teamName()
}

type TeamStatus struct {
	Abbrev string
	Score  int
}

func (ts TeamStatus) teamName() string {
	return ts.Abbrev
}

func NewTeamStatus(abbrev string, score int) TeamStatus {
	return TeamStatus{
		Abbrev: abbrev,
		Score:  score,
	}
}

type Inning struct {
	number int
	top    bool
}

func BottomInning(number int) Inning {
	return Inning{
		number: number,
		top:    false,
	}
}

func TopInning(number int) Inning {
	return Inning{
		number: number,
		top:    true,
	}
}

func (i Inning) String() string {
	if i.top {
		return fmt.Sprintf("↑%d", i.number)
	} else {
		return fmt.Sprintf("↓%d", i.number)
	}
}
