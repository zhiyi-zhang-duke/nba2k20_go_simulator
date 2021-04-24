package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

type team struct {
	name           string
	score          int
	scoreBreakdown string
}

// reads csv of team and generates a random set of scores
// for each player based on a normal distribution of their average
func genEducatedScore(team string) (score int, scoreBreakdown string) {
	dat, err := ioutil.ReadFile(team + ".csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(strings.NewReader(string(dat)))

	var totalMinutesPlayed float64 = 0
	playerMap := make(map[string]float64)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(record)

		playerAverage, floaterr1 := strconv.ParseFloat(record[27], 32)
		playerName := strings.Split(record[1], "\\")[0]
		if floaterr1 == nil {
			playerMap[playerName] = playerAverage
		}
		playerMinutesPlayed, floaterr2 := strconv.ParseFloat(record[5], 32)
		if floaterr2 == nil {
			totalMinutesPlayed += playerMinutesPlayed
		}
	}

	//We create a modifier to account for score inflation from play time
	//(Some players sit out certain games inflating the average points per game)
	modifier := 240 / totalMinutesPlayed
	//Debug messages
	//fmt.Printf("The %s have %d players\n", team, len(playerMap))
	//fmt.Println("PlayerMap:", playerMap)

	total := 0
	scoreBreakdown = ""
	for player, score := range playerMap {
		dist := distuv.Normal{
			Mu:    score,
			Sigma: score / 3,
		}
		randomScore := int(math.Round(dist.Rand()) * modifier)
		total += randomScore
		scoreBreakdown += fmt.Sprintf("%s scored %d\n", player, randomScore)
		// fmt.Printf("%s scored %d\n", player, randomScore)
	}
	return total, scoreBreakdown
}

//reference function from distuv
// func exampleDistribution() {
// 	dist := distuv.Normal{
// 		Mu:    2,
// 		Sigma: 5,
// 	}

// 	data := make([]float64, 10)

// 	for i := range data {
// 		data[i] = dist.Rand()
// 	}
// 	fmt.Println(data)
// 	mean, std := stat.MeanStdDev(data, nil)
// 	meanErr := stat.StdErr(std, float64(len(data)))

// 	fmt.Printf("mean= %1.1f ± %0.1v\n", mean, meanErr)
// }

func main() {
	//Generate a random seed so that scores are truly random
	rand.Seed(time.Now().UnixNano())

	//Create two teams with completely random score totals from 0 to 150
	var teamA_name = "lakers"
	teamA_score, teamA_scorebreakdown := genEducatedScore(teamA_name)
	teamA := team{teamA_name, teamA_score, teamA_scorebreakdown}

	var teamB_name = "clippers"
	teamB_score, teamB_scorebreakdown := genEducatedScore(teamB_name)
	teamB := team{teamB_name, teamB_score, teamB_scorebreakdown}

	//Print the result
	if teamA.score > teamB.score {
		fmt.Printf("The %s beat the %s.", teamA.name, teamB.name)
		fmt.Printf("The score was %d to %d.\n", teamA.score, teamB.score)
		fmt.Printf("%s Score Breakdown:\n%s \n", teamA.name, teamA.scoreBreakdown)
		fmt.Printf("%s Score Breakdown:\n%s \n", teamB.name, teamB.scoreBreakdown)
	} else {
		fmt.Printf("The %s beat the %s.", teamB.name, teamA.name)
		fmt.Printf("The score was %d to %d.\n", teamB.score, teamA.score)
		fmt.Printf("%s Score Breakdown:\n%s \n", teamA.name, teamA.scoreBreakdown)
		fmt.Printf("%s Score Breakdown:\n%s \n", teamB.name, teamB.scoreBreakdown)
	}
}
