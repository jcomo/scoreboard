package main

import (
	"fmt"
	"time"
)

func main() {
	sb := NewScoreboard(DefaultHttpClient())
	games, err := sb.Get(time.Now())
	if err != nil {
		fmt.Println("ERROR!")
	} else {
		for _, g := range games {
			fmt.Println(g)
		}
	}
}
