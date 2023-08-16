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
	MyRock        = "X"
	MyPaper       = "Y"
	MyScissors    = "Z"
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
	switch myPlay {
	case MyRock:
		score += 1
		switch enemyPlay {
		case RockEnemy:
			score += 3
		case PaperEnemy:
			score += 0
		case ScissorsEnemy:
			score += 6
		}
	case MyPaper:
		score += 2
		switch enemyPlay {
		case RockEnemy:
			score += 6
		case PaperEnemy:
			score += 3
		case ScissorsEnemy:
			score += 0
		}
	case MyScissors:
		score += 3
		switch enemyPlay {
		case RockEnemy:
			score += 0
		case PaperEnemy:
			score += 6
		case ScissorsEnemy:
			score += 3
		}
	}
	return score
}
func main() {
	fmt.Println(readInputAndCalculateScores("../input.txt"))
}
