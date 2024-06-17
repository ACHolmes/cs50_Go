package main

import (
	"fmt"
	"log"
	"testing"
)

type TidemanElection interface {
	simulateElection()
}

type election struct {
	name           string
	candidateCount int
	candidates     []string
	voterCount     int
	votes          [][]string
	expectedWinner string
}

func simulateVoter(votes []string, ranks []int) {
	if len(votes) != candidateCount {
		panic("simulateVoter testing error. CandidateCount != length of votes provided")
	}
	for j, name := range votes {
		if !vote(j, name, ranks) {
			panic("Invalid vote.\n")
		}
	}
	recordPreferences(ranks)
}

func resetGlobals() {
	pairs = pairs[:0]
	for i := range maxCandidates {
		for j := range maxCandidates {
			edges[i][j] = false
		}
	}
	clear(preferences)
}

func (election election) simulateElection() string {
	resetGlobals()
	candidateCount = election.candidateCount
	if candidateCount < 1 {
		panic("simulation error: Need >=1 candidates")
	}

	if candidateCount > maxCandidates {
		panic(fmt.Sprintf("Maximum number of candidates is %v\n", maxCandidates))
	}

	for i := range candidateCount {
		candidates[i] = election.candidates[i]
	}

	for i := 0; i < election.voterCount; i++ {
		ranks := make([]int, candidateCount)

		simulateVoter(election.votes[i], ranks)
	}
	return tideman()
}

func TestTideman(t *testing.T) {

	testElectionOne := election{
		name:           "Basic One Voter Election",
		candidateCount: 3,
		candidates:     []string{"Alice", "Bob", "Charlie"},
		voterCount:     1,
		votes: [][]string{
			{"Alice", "Bob", "Charlie"},
		},
		expectedWinner: "Alice",
	}

	testElectionTwo := election{
		name:           "Provided test case from website",
		candidateCount: 3,
		candidates:     []string{"Alice", "Bob", "Charlie"},
		voterCount:     9,
		votes: [][]string{
			{"Alice", "Bob", "Charlie"},
			{"Alice", "Bob", "Charlie"},
			{"Alice", "Bob", "Charlie"},
			{"Bob", "Charlie", "Alice"},
			{"Bob", "Charlie", "Alice"},
			{"Charlie", "Alice", "Bob"},
			{"Charlie", "Alice", "Bob"},
			{"Charlie", "Alice", "Bob"},
			{"Charlie", "Alice", "Bob"},
		},
		expectedWinner: "Charlie",
	}

	elections := []election{testElectionOne, testElectionTwo}

	for _, e := range elections {
		t.Run(e.name, func(t *testing.T) {
			winner := e.simulateElection()
			if winner == e.expectedWinner {
				fmt.Println(e.name, "passed")
			} else {
				log.Fatalf("Failed test (%v), winner: %v, expectedWinner: %v", e.name, winner, e.expectedWinner)
			}
		})

	}
}
