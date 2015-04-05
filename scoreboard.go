package main

type Scoreboard struct {
}

func NewScoreboard() *Scoreboard {
	return &Scoreboard{}
}

func (sb *Scoreboard) Get() string {
	return "NYY vs BOS 7:05PM"
}

func (sb *Scoreboard) GetTeam(team string) string {
	return "WSH 3 - 4 NYY F"
}
