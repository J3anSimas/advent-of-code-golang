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
	groups := strings.Split(inputData, "\n\n")
	mostCalories := 0
	for _, group := range groups {
		sum := 0
		lines := strings.Split(group, "\n")
		for _, line := range lines {
			cal, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("Error converting string to int")
			}
			sum += cal
		}
		if sum > mostCalories {
			mostCalories = sum
		}
	}

	return mostCalories
}
func main() {
	fmt.Println(readInputAndFindMostCalories("../input.txt"))
}
