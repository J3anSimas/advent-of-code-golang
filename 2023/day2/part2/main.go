package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	raw_string := string(file)
	lines := strings.Split(strings.Replace(raw_string, "\r\n", "\n", -1), "\n")
	sum := 0
	for _, line := range lines {
		min_red := 0
		min_green := 0
		min_blue := 0
		s := strings.Split(line, ": ")
		subsets := strings.Split(s[1], "; ")
		for _, subset := range subsets {
			cubes := strings.Split(subset, ", ")
			for _, cube := range cubes {
				b := strings.Split(cube, " ")
				amount, err := strconv.Atoi(b[0])
				if err != nil {
					panic(err)
				}
				color := b[1]
				switch color {
				case "red":
					if amount > min_red {
						min_red = amount
					}
				case "green":
					if amount > min_green {
						min_green = amount
					}
				case "blue":
					if amount > min_blue {
						min_blue = amount
					}
				}

			}

		}
		sum += min_red * min_green * min_blue
	}

	fmt.Println(sum)
}
