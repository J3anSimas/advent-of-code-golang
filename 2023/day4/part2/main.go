package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	winning_numbers []int
	my_numbers      []int
	count           int
}

func (c *Card) calculate() int {
	count := 0
	for _, my_number := range c.my_numbers {
		for _, winning_number := range c.winning_numbers {
			if my_number == winning_number {
				count = count + 1
			}
		}
	}
	return count
}

func main() {
	dt := time.Now()
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	raw_string := string(file)
	lines := strings.Split(raw_string, "\n")
	cards := make([]Card, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}
		replaced_line :=
			strings.ReplaceAll(line, "  ", " ")
		broken_line := strings.Split(strings.Split(replaced_line, ": ")[1], " | ")
		winning_numbers := strings.Split(broken_line[0], " ")
		my_numbers := strings.Split(broken_line[1], " ")
		cards[i] = Card{
			winning_numbers: convertSliceOfStringsToInts(winning_numbers),
			my_numbers:      convertSliceOfStringsToInts(my_numbers),
			count:           1,
		}
	}
	// sum := 0
	// for _, number := range my_numbers {
	// 	parsed_my_number, err := strconv.Atoi(number)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	for _, winning_number := range winning_numbers {
	// 		parsed_winning_number, err := strconv.Atoi(winning_number)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		if parsed_my_number == parsed_winning_number {
	// 			if sum == 0 {
	// 				sum = 1
	// 			} else {
	// 				sum = sum * 2
	// 			}
	// 		}
	// 	}
	// }

	for i := range cards {
		calculate(&cards, i)
	}
	sum := 0
	for _, card := range cards {
		sum += card.count
	}
	fmt.Println(sum)
	elapsed := time.Since(dt)
	fmt.Printf("Solution took %s", elapsed)
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

func calculate(cards *[]Card, idx int) {
	count := (*cards)[idx].calculate()
	for i := 1; i <= count; i++ {
		(*cards)[idx+i].count += 1
		calculate(cards, idx+i)
	}
}
