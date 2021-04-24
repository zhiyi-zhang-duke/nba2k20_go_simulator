# nba2k20_simulator_go
This is a simple go project that simulates nba games from the NBA 19-20 season using stats from basketballreference.com

Implementation:

Given 2 teams, the simulator will read in the season box score for the entire NBA team.
The simulator then estimates each players score using a normal distribution of their average points per game

Details:
The standard deviation of the score distribution is set to score/3, potentially a player could score well below their typical average.
The simulator accounts for minutes played by taking total game time (240 minutes) and dividing by the total minutes played from the team's boxscore. This modified is used to downgrade the total average points of each player so that game points are not artificially inflated.

## Example output
The lakers beat the clippers.The score was 101 to 86.
lakers Score Breakdown:
LeBron James scored 20
Kentavious Caldwell-Pope scored 6
J.R. Smith scored 2
Zach Norvell scored 0
Anthony Davis scored 12
JaVale McGee scored 4
Quinn Cook scored 2
Danny Green scored 5
Dion Waiters scored 9
Kostas Antetokounmpo scored 1
Alex Caruso scored 4
Talen Horton-Tucker scored 4
Troy Daniels scored 3
Devontae Cacok scored 2
Kyle Kuzma scored 9
Avery Bradley scored 4
Rajon Rondo scored 4
Dwight Howard scored 6
Jared Dudley scored 0

clippers Score Breakdown:
JaMychal Green scored 5
Ivica Zubac scored 4
Joakim Noah scored 1
Derrick Walton scored 1
Johnathan Motley scored 1
Paul George scored 11
Lou Williams scored 13
Montrezl Harrell scored 10
Maurice Harkless scored 1
Reggie Jackson scored 4
Jerome Robinson scored 1
Amir Coffey scored 1
Marcus Morris scored 6
Rodney McGruder scored 1
Patrick Patterson scored 3
Mfiondu Kabengele scored 2
Kawhi Leonard scored 10
Landry Shamet scored 7
Patrick Beverley scored 3
Terance Mann scored 1

PS C:\Users\Jack\Desktop\Jack\Code Projects\nba2k20_simulator_go> go run .\nba2k20.go
The clippers beat the lakers.The score was 102 to 99.
lakers Score Breakdown:
Alex Caruso scored 4
Markieff Morris scored 3
Zach Norvell scored 0
Danny Green scored 2
Avery Bradley scored 4
Dion Waiters scored 8
Dwight Howard scored 4
Quinn Cook scored 3
Troy Daniels scored 3
Devontae Cacok scored 4
Kentavious Caldwell-Pope scored 10
Kyle Kuzma scored 10
Rajon Rondo scored 6
J.R. Smith scored 2
Jared Dudley scored 0
Kostas Antetokounmpo scored 1
LeBron James scored 12
Anthony Davis scored 16
JaVale McGee scored 5
Talen Horton-Tucker scored 2

clippers Score Breakdown:
Kawhi Leonard scored 23
Paul George scored 11
Maurice Harkless scored 1
Reggie Jackson scored 8
Marcus Morris scored 6
Amir Coffey scored 1
Montrezl Harrell scored 14
Landry Shamet scored 5
Ivica Zubac scored 2
Patrick Patterson scored 2
Joakim Noah scored 1
Derrick Walton scored 1
Terance Mann scored 1
Johnathan Motley scored 1
Lou Williams scored 14
Patrick Beverley scored 5
JaMychal Green scored 2
Rodney McGruder scored 2
Jerome Robinson scored 1
Mfiondu Kabengele scored 1

## Things I found helpful
https://gobyexample.com/

https://learn.go.dev/

https://medium.com/rungo/how-to-write-a-simple-go-program-13fd104f3018

https://stackoverflow.com/questions/66894200/go-go-mod-file-not-found-in-current-directory-or-any-parent-directory-see-go
