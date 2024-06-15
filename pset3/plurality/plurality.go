package main

import (
	"fmt"
	"os"

	"example.com/lib50"
)

func plurality(candidates []string, numVoters int) []string {

	resultsTable := make(map[string]int)
	for _, cand := range candidates {
		resultsTable[cand] = 0
	}

	for voterId := 0; voterId < numVoters; voterId++ {
		voteFor := lib50.GetString("Vote: ")
		_, ok := resultsTable[voteFor]
		if ok {
			resultsTable[voteFor]++
		} else {
			fmt.Println("Invalid vote.")
		}
	}

	var mostVotes int
	for _, votes := range resultsTable {
		if votes > mostVotes {
			mostVotes = votes
		}
	}

	winners := make([]string, 0)

	for name, votes := range resultsTable {
		if votes == mostVotes {
			winners = append(winners, name)
		}
	}
	return winners

}

func main() {
	candidates := os.Args[1:]
	if len(candidates) == 0 || len(candidates) > 9 {
		fmt.Println("Invalid number of candidates")
		return
	}
	numVoters := lib50.GetInt("Number of voters: ")

	winners := plurality(candidates, numVoters)
	for _, winner := range winners {
		fmt.Println(winner)
	}

}
