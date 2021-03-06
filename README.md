# Scoreboard
A simple CLI to retrieve MLB scores and upcoming games.

## Installation

To build and install scoreboard, you must have Go version >= 1.3.2 on your system.

```sh
$ go get github.com/jcomo/scoreboard
$ scoreboard
```

It may be necessary to run `go install` from within the directory depending on your PATH
configuration.


## Usage
The basic usage of scoreboard is to run the program with no arguments

```sh
$ scoreboard

NYY vs BOS 7:05PM
SF 4 • 1 OAK ↑9 
# ...
```

You may also specify a team using the team flag to get information about one
team only.

The specified team should be the abbreviation (eg. NYY). This is case
insensitive.

```sh
$ scoreboard -team=nyy

NYY 5 • 3 BOS ↓4
```


## Behavior

Scoreboard will display the score and state of the game if it is currently in
play. If there is no game in play for a team, scoreboard will show the upcoming
game in the schedule instead.

