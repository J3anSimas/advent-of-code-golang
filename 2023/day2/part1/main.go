package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube = string

const (
	Blue  Cube = "blue"
	Red   Cube = "red"
	Green Cube = "green"
)

type InfoCube struct {
	Amount int
	Color  Cube
}

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	raw_string := string(file)
	lines := strings.Split(strings.Replace(raw_string, "\r\n", "\n", -1), "\n")
	sum := 0
	for _, line := range lines {
		couldBe := true
		s := strings.Split(line, ": ")
		gameId, err := strconv.Atoi(strings.Split(s[0], " ")[1])
		if err != nil {
			panic(err)
		}
		subsets := strings.Split(s[1], "; ")
		for _, subset := range subsets {
			red := 0
			green := 0
			blue := 0
			blocks := strings.Split(subset, ", ")
			for _, block := range blocks {
				b := strings.Split(block, " ")
				amount, err := strconv.Atoi(b[0])
				if err != nil {
					panic(err)
				}
				color := b[1]
				switch color {
				case "red":
					red += amount
				case "green":
					green += amount
				case "blue":
					blue += amount
				}
			}
			if red > 12 || green > 13 || blue > 14 {
				couldBe = false
			}
		}
		if couldBe {
			sum += gameId
		}

	}

	fmt.Println(sum)
}
