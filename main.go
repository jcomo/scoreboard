package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func printCurrentTeamStatus(sb *Scoreboard, team string) {
	game, err := sb.GetTeam(time.Now(), team)
	if err != nil {
		fmt.Println("ERROR!")
	} else {
		fmt.Println(game)
	}
}

func printCurrentStatus(sb *Scoreboard) {
	games, err := sb.Get(time.Now())
	if err != nil {
		fmt.Println("ERROR!")
	} else {
		for _, g := range games {
			fmt.Println(g)
		}
	}
}

func main() {
	teamFlag := flag.String("team", "", "Team to focus on (eg. NYY)")
	flag.Parse()

	sb := NewScoreboard(DefaultHttpClient())
	team := strings.ToUpper(*teamFlag)
	if team != "" {
		printCurrentTeamStatus(sb, team)
	} else {
		printCurrentStatus(sb)
	}
}
