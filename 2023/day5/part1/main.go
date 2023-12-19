package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Seed struct {
	Seed        int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

func main() {
	// filename := "../input"
	// if len(os.Args) > 2 {
	// 	filename += "_test"
	// }
	// filename += ".txt"
	//
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	raw_string := string(file)
	blocks := strings.Split(raw_string, "\n\n")
	runtime.GC()
	unparsed_seeds := strings.Split(blocks[0], ": ")[1]
	seeds := make([]int, 0)
	seeds = convertSliceOfStringsToInts(strings.Split(unparsed_seeds, " "))
	unparsed_seeds = ""
	runtime.GC()
	min := -1
	for i := range seeds {
		soil := findInTable(seeds[i], strings.Split(blocks[1], "\n")[1:])
		fertilizer := findInTable(soil, strings.Split(blocks[2], "\n")[1:])
		water := findInTable(fertilizer, strings.Split(blocks[3], "\n")[1:])
		light := findInTable(water, strings.Split(blocks[4], "\n")[1:])
		temperature := findInTable(light, strings.Split(blocks[5], "\n")[1:])
		humidity := findInTable(temperature, strings.Split(blocks[6], "\n")[1:])
		location := findInTable(humidity, strings.Split(blocks[7], "\n")[1:])
		if location < min || min == -1 {
			min = location
		}
	}
	runtime.GC()

	fmt.Println(min)

}

func findInTable(value int, lines []string) int {
	result := -1
	for _, line := range lines {
		if line == "" {
			continue
		}
		ns := strings.Split(line, " ")
		parsed_numbers := convertSliceOfStringsToInts(ns)
		if parsed_numbers[1]+parsed_numbers[2] < value {
			result = value
			continue
		}
		offset := value - parsed_numbers[1]

		if offset < 0 {
			result = value
			continue
		}
		result = parsed_numbers[0] + offset
		break

	}

	return result
}

func convertSliceOfStringsToInts(slice []string) []int {
	ints := make([]int, len(slice))
	for i, s := range slice {
		parsed, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] = parsed
	}
	return ints
}
