# Scoreboard
A simple CLI to retrieve MLB scores and upcoming games.

## Usage
The basic usage of scoreboard is to run the program with no arguments

```sh
$ scoreboard

NYY vs BOS 7:05PM
SF 4 - 1 OAK T9 
# ...
```

You may also specify a team using the team flag to get information about one
team only.

The specified team should be the abbreviation (eg. NYY). This is case
insensitive.

```sh
$ scoreboard -team nyy

NYY 5 - 3 BOS B4
```


## Behavior

Scoreboard will display the score and state of the game if it is currently in
play. If there is no game in play for a team, scoreboard will show the upcoming
game in the schedule instead.

When a team is specified, the upcoming game will be displayed with respect to
the team.

For example,

```
BOS vs NYY 7:05PM
NYY @ OAK 1:05PM
```