# nba2k20_simulator_go
This is a simple go project that simulates nba games from the NBA 19-20 season using stats from basketballreference.com

Implementation:

Given 2 teams, the simulator will read in the season box score for the entire NBA team.
The simulator then estimates each players score using a normal distribution of their average points per game

Details:
The standard deviation of the score distribution is set to score/3, potentially a player could score well below their typical average.
The simulator accounts for minutes played by taking total game time (240 minutes) and dividing by the total minutes played from the team's boxscore. This modified is used to downgrade the total average points of each player so that game points are not artificially inflated.

## Things I found helpful
https://gobyexample.com/

https://learn.go.dev/

https://medium.com/rungo/how-to-write-a-simple-go-program-13fd104f3018

https://stackoverflow.com/questions/66894200/go-go-mod-file-not-found-in-current-directory-or-any-parent-directory-see-go
