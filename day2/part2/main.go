package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	RockEnemy     = "A"
	PaperEnemy    = "B"
	ScissorsEnemy = "C"
	Lose          = "X"
	Draw          = "Y"
	Win           = "Z"
)

func readInputAndCalculateScores(filename string) int {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file")
	}
	inputData := string(raw_data)
	raw_data = make([]byte, 0)
	lines := strings.Split(inputData, "\n")
	var score int
	for _, line := range lines {
		plays := strings.Split(line, " ")
		score += getPlayScore(plays[0], plays[1])
	}
	return score
}

func getPlayScore(enemyPlay, myPlay string) int {
	var score int
	switch enemyPlay {
	case RockEnemy:
		switch myPlay {
		case Lose:
			score += 3
		case Draw:
			score += 4
		case Win:
			score += 8
		}
	case PaperEnemy:
		switch myPlay {
		case Lose:
			score += 1
		case Draw:
			score += 5
		case Win:
			score += 9
		}
	case ScissorsEnemy:
		switch myPlay {
		case Lose:
			score += 2
		case Draw:
			score += 6
		case Win:
			score += 7
		}
	}
	return score
}
func main() {
	fmt.Println(readInputAndCalculateScores("../input.txt"))
}
