package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInputAndUseResult(filename string) int {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file")
	}
	inputData := string(raw_data)
	raw_data = make([]byte, 0)
	lines := strings.Split(inputData, "\n")
	var count int
	for _, line := range lines {
		pair := strings.Split(line, ",")
		first, err := convertStringSliceToIntSlice(strings.Split(pair[0], "-"))
		if err != nil {
			log.Fatal("Could not convert string slice to int slice", err)
		}
		second, err := convertStringSliceToIntSlice(strings.Split(pair[1], "-"))
		if err != nil {
			log.Fatal("Could not convert string slice to int slice", err)
		}
		if checkIfOverlaps(first, second) {
			count++
		}
	}
	return count
}

func checkIfOverlaps(first []int, second []int) bool {
	if first[0] <= second[0] && first[1] >= second[0] {
		return true
	}
	if second[0] <= first[0] && second[1] >= first[0] {
		return true
	}
	return false
}
func convertStringSliceToIntSlice(first []string) ([]int, error) {
	var result []int
	for _, item := range first {
		conv, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		result = append(result, conv)
	}
	return result, nil
}
func main() {
	fmt.Println(readInputAndUseResult("../input.txt"))
}
