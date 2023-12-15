package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var alpha = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func readInputAndUseResult(filename string) int {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file")
	}
	inputData := string(raw_data)
	raw_data = make([]byte, 0)
	lines := strings.Split(inputData, "\n")
	var score int
	for i := 0; i < len(lines); i += 3 {
		score += getBadgeScore(lines[i : i+3])
	}
	return score
}

func getBadgeScore(group []string) int {
	var badArrangedItems string
	for _, item := range group[0] {
		if strings.Contains(group[1], string(item)) {
			if strings.Contains(group[2], string(item)) {
				badArrangedItems = string(item)
				break
			}
		}
	}

	return strings.Index(strings.Join(alpha, ""), badArrangedItems) + 1
}
func main() {
	fmt.Println(readInputAndUseResult("../input.txt"))
  num := 1
}
