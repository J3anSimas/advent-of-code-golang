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
	mostCalories := 0
	currentCalories := 0
	for _, line := range lines {
		if line == "" {
			if currentCalories > mostCalories {
				mostCalories = currentCalories
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
	return mostCalories
}
func main() {
	fmt.Println(readInputAndFindMostCalories("../input.txt"))
}
