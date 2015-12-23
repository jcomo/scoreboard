package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/jcomo/scoreboard/mlb"
)

func printCurrentTeamStatus(sb *Scoreboard, team string) {
	game, err := sb.GetTeam(time.Now(), team)
	if err != nil {
		fmt.Println("Error!")
	} else {
		fmt.Println(game)
	}
}

func printCurrentStatus(sb *Scoreboard) {
	games, err := sb.Get(time.Now())
	if err != nil {
		fmt.Println("Error!")
	} else {
		if len(games) == 0 {
			fmt.Println("No games")
		} else {
			for _, g := range games {
				fmt.Println(g)
			}
		}
	}
}

func main() {
	teamFlag := flag.String("team", "", "Team to focus on (eg. NYY)")
	flag.Parse()

	sb := NewScoreboard(mlb.DefaultHttpClient())
	team := strings.ToUpper(*teamFlag)
	if team != "" {
		printCurrentTeamStatus(sb, team)
	} else {
		printCurrentStatus(sb)
	}
}
