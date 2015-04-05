package main

type RawGame struct {
    time string
    amOrPm string
    homeName string
    awayName string
    lineScore lineScore
    gameState gameState
}

type lineScore struct {
    runs runs
}

type runs struct {
    home int
    away int
}

type gameState struct {
    state string
    inning int
    inningState string
}


func GameFromRaw(rg RawGame) Game {
    if rg.gameState.state == "Preview" {
        return upcomingGameFromRaw(rg)
    } else if rg.gameState.state == "Final" {
        return finishedGameFromRaw(rg)
    } else {
        return inProgressGameFromRaw(rg)
    }
}

func upcomingGameFromRaw(rg RawGame) Game {
    return UpcomingGame{
        home: rg.homeName,
        away: rg.awayName,
        time: parseGameTime(rg.time + rg.amOrPm),
    }
}

func inProgressGameFromRaw(rg RawGame) Game {
    return InProgressGame{
        home: homeTeamStatus(rg),
        away: awayTeamStatus(rg),
        inning: inning{
            number: rg.gameState.inning,
            top: rg.gameState.inningState == "Top",
        },
    }
}

func finishedGameFromRaw(rg RawGame) Game {
    return FinishedGame{
        home: homeTeamStatus(rg),
        away: awayTeamStatus(rg),
    }
}

func homeTeamStatus(rg RawGame) teamStatus {
    return newTeamStatus(rg.homeName, rg.lineScore.runs.home)
}

func awayTeamStatus(rg RawGame) teamStatus {
    return newTeamStatus(rg.awayName, rg.lineScore.runs.away)
}