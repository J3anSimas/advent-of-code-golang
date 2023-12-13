package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	raw_string := string(file)
	lines := strings.Split(raw_string, "\n")
	_ = lines
	s := 0
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			char := line[j]
			parsed_char := string(char)
			if parsed_char == "*" {
				s += get_numbers(lines, i, j)
			}
		}
	}
	fmt.Println(s)
}

type Number struct {
	Number int
	start  int
	end    int
}

func get_numbers(lines []string, i, g int) int {
	numbers := []Number{}
	for j := -1; j < 2; j++ {
		if i+j < 0 {
			continue
		}
		if i+j >= len(lines) {
			continue
		}
		n := ""
		start := 0
		end := 0
		for k, char := range lines[i+j] {
			if unicode.IsDigit(char) {
				if n == "" {
					start = k
				}
				end = k
				n += string(char)
			} else {
				if n != "" {
					if isAdjacent(start, end, g) {
						nn, err := strconv.Atoi(n)
						if err != nil {
							panic(err)
						}
						numbers = append(numbers, Number{Number: nn, start: start, end: end})
					}
					n = ""
				}
			}
		}

	}
	if len(numbers) == 2 {
		return numbers[0].Number * numbers[1].Number
	}
	return 0
}

func isAdjacent(start, end, p int) bool {
	if end >= p-1 && end <= p+1 {
		return true
	}
	if start >= p-1 && start <= p+1 {
		return true
	}
	return false
}
