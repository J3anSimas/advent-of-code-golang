package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInputAndFindMostCalories(filename string) int {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file")
	}
	inputData := string(raw_data)
	raw_data = make([]byte, 0)
	lines := strings.Split(inputData, "\n")
	top1, top2, top3 := 0, 0, 0
	currentCalories := 0
	for _, line := range lines {
		if line == "" {
			if currentCalories > top1 {
				top3 = top2
				top2 = top1
				top1 = currentCalories
			} else if currentCalories > top2 {
				top3 = top2
				top2 = currentCalories
			} else if currentCalories > top3 {
				top3 = currentCalories
			}
			currentCalories = 0
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Error converting string to int")
		}
		currentCalories += calories
	}
	return top1 + top2 + top3
}
func main() {

	fmt.Println(readInputAndFindMostCalories("../input.txt"))
}
