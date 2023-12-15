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
	lines := strings.Split(raw_string, "\n")
	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		replaced_line :=
			strings.ReplaceAll(line, "  ", " ")
		broken_line := strings.Split(strings.Split(replaced_line, ": ")[1], " | ")
		winning_numbers := strings.Split(broken_line[0], " ")
		my_numbers := strings.Split(broken_line[1], " ")
		sum := 0
		for _, number := range my_numbers {
			parsed_my_number, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			for _, winning_number := range winning_numbers {
				parsed_winning_number, err := strconv.Atoi(winning_number)
				if err != nil {
					panic(err)
				}
				if parsed_my_number == parsed_winning_number {
					if sum == 0 {
						sum = 1
					} else {
						sum = sum * 2
					}
				}
			}
		}
		total += sum
	}
	fmt.Println(total)
}
