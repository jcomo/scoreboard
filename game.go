package main

import (
    "fmt"
)

type Game struct {
    homeTeam string
    awayTeam string
    homeScore int
    awayScore int
    inning int
    inningState string
    time string
    state string
}

func (g Game) State() string {
    if g.state == "Upcoming" {
        return fmt.Sprintf("%s vs %s %s", g.awayTeam, g.homeTeam, g.time)    
    } else if g.state == "In Progress" {
        return fmt.Sprintf("%s %d - %d %s %s%d",
                           g.awayTeam, g.awayScore, g.homeScore, g.homeTeam,
                           g.inningState, g.inning)
    } else {
        return fmt.Sprintf("%s %d - %d %s F",
                           g.awayTeam, g.awayScore, g.homeScore, g.homeTeam)
    }
    
}