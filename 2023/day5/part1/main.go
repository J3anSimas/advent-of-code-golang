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
	unparsed_seeds := strings.Split(blocks[0], ": ")
	seeds := make([]int, 0)
	seeds = append(seeds, convertSliceOfStringsToInts(strings.Split(unparsed_seeds[1], " "))...)
	unparsed_seeds = nil
	var m runtime.MemStats
	fmt.Println("Começando parseTable")
	soils := parseTable(strings.Split(blocks[1], "\n")[1:])
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		m.Alloc/1024/1024,
		m.TotalAlloc/1024/1024,
		m.Sys/1024/1024,
		m.NumGC)
	blocks[1] = ""
	fertilizers := parseTable(strings.Split(blocks[2], "\n")[1:])
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		m.Alloc/1024/1024,
		m.TotalAlloc/1024/1024,
		m.Sys/1024/1024,
		m.NumGC)
	blocks[2] = ""
	waters := parseTable(strings.Split(blocks[3], "\n")[1:])
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		m.Alloc/1024/1024,
		m.TotalAlloc/1024/1024,
		m.Sys/1024/1024,
		m.NumGC)
	blocks[3] = ""
	lights := parseTable(strings.Split(blocks[4], "\n")[1:])
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		m.Alloc/1024/1024,
		m.TotalAlloc/1024/1024,
		m.Sys/1024/1024,
		m.NumGC)
	blocks[4] = ""
	temperatures := parseTable(strings.Split(blocks[5], "\n")[1:])
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		m.Alloc/1024/1024,
		m.TotalAlloc/1024/1024,
		m.Sys/1024/1024,
		m.NumGC)
	blocks[5] = ""
	humidities := parseTable(strings.Split(blocks[6], "\n")[1:])
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		m.Alloc/1024/1024,
		m.TotalAlloc/1024/1024,
		m.Sys/1024/1024,
		m.NumGC)
	blocks[6] = ""
	locations := parseTable(strings.Split(blocks[7], "\n")[1:])
	runtime.GC()
	min := -1
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		m.Alloc/1024/1024,
		m.TotalAlloc/1024/1024,
		m.Sys/1024/1024,
		m.NumGC)
	// for i := range seeds {
	// 	soil := findInTable(seeds[i], soils)
	// 	fertilizer := findInTable(soil, fertilizers)
	// 	water := findInTable(fertilizer, waters)
	// 	light := findInTable(water, lights)
	// 	temperature := findInTable(light, temperatures)
	// 	humidity := findInTable(temperature, humidities)
	// 	location := findInTable(humidity, locations)
	// 	runtime.ReadMemStats(&m)
	// 	fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
	// 		m.Alloc/1024/1024,
	// 		m.TotalAlloc/1024/1024,
	// 		m.Sys/1024/1024,
	// 		m.NumGC)
	// 	if location < min || min == -1 {
	// 		min = location
	// 	}
	// }
	runtime.GC()
	// tables := make([][][]int, len(blocks[1:]))
	// for i, b := range blocks[1:] {
	// 	lines := strings.Split(b, "\n")
	// 	table := parseTable(lines[1:])
	// 	if len(table) == 0 {
	// 		continue
	// 	}
	// 	tables[i] = table
	// }
	// for _, table := range tables {
	// 	seeds = goToNextLevel(seeds, table)
	// }
	_ = soils
	_ = fertilizers
	_ = waters
	_ = lights
	_ = temperatures
	_ = humidities
	_ = locations

	fmt.Println(min)

}

func parseTable(lines []string) [][]int {
	result := make([][]int, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		ns := strings.Split(line, " ")
		parsed_numbers := convertSliceOfStringsToInts(ns)
		for i := 0; i < parsed_numbers[2]; i++ {
			result = append(result, []int{parsed_numbers[0] + i, parsed_numbers[1] + i})
		}
	}
	fmt.Println(len(result))
	return result

}
func findInTable(value int, table [][]int) int {
	result := -1
	for _, coordinate := range table {
		if value == coordinate[1] {
			result = coordinate[0]
		}
	}
	if result == -1 {
		result = value
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
