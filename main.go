package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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
	outFlag := flag.String("log-file", "", "File to send log output")
	prefixFlag := flag.String("log-prefix", "[SCOREBOARD]", "Prefix to apply to logs")
	flag.Parse()

	sb := NewScoreboard()
	if *outFlag != "" {
		f, err := os.Open(*outFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sb.Logger = log.New(f, *prefixFlag+" ", log.LstdFlags)
	}

	team := strings.ToUpper(*teamFlag)
	if team != "" {
		printCurrentTeamStatus(sb, team)
	} else {
		printCurrentStatus(sb)
	}
}
