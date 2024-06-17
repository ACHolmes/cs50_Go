package main

import (
	"fmt"
	"os"

	"example.com/lib50"
)

type pair struct {
	winner int
	loser  int
}

const (
	maxCandidates = 9
)

var (
	candidateCount = 0
	candidates     [maxCandidates]string
	preferences    = make(map[int]map[int]int, 1)
	pairs          = make([]pair, 0, 10)
	edges          [maxCandidates][maxCandidates]bool
)

func vote(rank int, name string, ranks []int) bool {
	var candIdx int
	ok := false
	for i, testName := range candidates {
		if name == testName {
			candIdx = i
			ok = true
			break
		}
	}
	if !ok {
		return false
	}
	ranks[rank] = candIdx
	return true
}

func recordPreferences(ranks []int) {
	for i, preference := range ranks {
		for _, loser := range ranks[i+1:] {
			_, ok := preferences[preference]
			if !ok {
				preferences[preference] = make(map[int]int)
			}
			preferences[preference][loser]++
		}
	}
}

func addPairs() {
	for i, _ := range preferences {
		for j, _ := range preferences[i] {
			if preferences[i][j] > preferences[j][i] {
				pairs = append(pairs, pair{
					winner: i,
					loser:  j,
				})
			}
		}
	}
}

func sortPairs() {
	for i, _ := range pairs {
		maxIdx := i
		for j := i + 1; j < len(pairs); j++ {
			topPair := pairs[j]
			if preferences[pairs[maxIdx].winner][pairs[maxIdx].loser] <
				preferences[topPair.winner][topPair.loser] {
				maxIdx = j
			}
		}

		tmpPair := pairs[maxIdx]
		pairs[maxIdx] = pairs[i]
		pairs[i] = tmpPair
	}
}

func lockPairs() {
	wouldCreateCycle := func(pair pair) bool {
		visited := make([]bool, len(candidates))

		var dfsHelper func(int, int) bool

		dfsHelper = func(loser, winner int) bool {
			visited[loser] = true
			for j, isEdge := range edges[loser] {
				if isEdge && !visited[j] {
					if j == winner {
						return true
					}
					visited[j] = true
					if dfsHelper(j, winner) {
						return true
					}

				}
			}
			return false
		}
		return dfsHelper(pair.loser, pair.winner)
	}

	for _, pair := range pairs {
		if !wouldCreateCycle(pair) {
			edges[pair.winner][pair.loser] = true
		}
	}
}

func printWinner() string {
	for src := range candidateCount {
		trueSource := true
		for dst := range candidateCount {
			if edges[dst][src] {
				trueSource = false
			}
		}
		if trueSource {
			return candidates[src]
		}
	}
	return "No winner"
}

func main() {

	argv := os.Args[1:]
	candidateCount = len(argv)
	if candidateCount < 1 {
		panic("Usage: tideman [candidate ...]\n")
	}

	if candidateCount > maxCandidates {
		panic(fmt.Sprintf("Maximum number of candidates is %v\n", maxCandidates))
	}

	for i := range candidateCount {
		candidates[i] = os.Args[1+i]
	}

	voterCount := lib50.Getint("Number of voters: ")

	// Query for votes
	for i := 0; i < voterCount; i++ {
		ranks := make([]int, candidateCount)

		// Query for each rank
		for j := range candidateCount {
			name := lib50.GetString("Rank %v: ", j+1)

			if !vote(j, name, ranks) {
				panic("Invalid vote.\n")
			}
		}

		recordPreferences(ranks)

		fmt.Println()
	}
	fmt.Println(tideman())
}

func tideman() string {
	addPairs()
	sortPairs()
	lockPairs()
	return printWinner()
}
