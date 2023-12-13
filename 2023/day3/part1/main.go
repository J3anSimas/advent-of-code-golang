package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Point struct {
	x int
	y int
}
type Number struct {
	value      int
	startIndex Point
	endIndex   Point
}

func (n Number) String() string {
	return fmt.Sprintf("%d (%d,%d) (%d,%d)", n.value, n.startIndex.x, n.startIndex.y, n.endIndex.x, n.endIndex.y)
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	raw_string := string(file)
	lines := strings.Split(raw_string, "\n")

	numbers := []Number{}
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			char := line[j]
			if unicode.IsDigit(rune(char)) {
				startIndex := Point{x: j, y: i}
				endIndex := Point{x: j, y: i}
				count := 0
				n := ""
				for unicode.IsDigit(rune(line[j+count])) {
					n += string(line[j+count])
					count++
					endIndex.x++
				}
				j += count
				num, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				number := Number{value: num, startIndex: startIndex, endIndex: endIndex}
				numbers = append(numbers, number)
			}
		}
	}
	valid_numbers := []Number{}
	sum := 0
	for _, number := range numbers {
	out:
		for i := number.startIndex.y - 1; i <= number.endIndex.y+1; i++ {
			if i < 0 || i >= len(lines) {
				continue
			}
			for j := number.startIndex.x - 1; j <= number.endIndex.x; j++ {
				if j < 0 || j >= len(lines[i]) {
					continue
				}
				char := string(lines[i][j])
				if char != "." && !unicode.IsDigit(rune(lines[i][j])) {
					valid_numbers = append(valid_numbers, number)
					sum += number.value
					break out
				}
			}
		}
	}
	fmt.Println(sum)
	str := ""
	for _, number := range valid_numbers {
		for i := number.startIndex.y - 1; i <= number.endIndex.y+1; i++ {
			if i < 0 || i >= len(lines) {
				continue
			}
			for j := number.startIndex.x - 1; j <= number.endIndex.x; j++ {
				if j < 0 || j >= len(lines[i]) {
					continue
				}
				char := string(lines[i][j])
				str += fmt.Sprintf("%s", char)
			}
			str += "\n"
		}
		str += "\n===========================\n"
	}

	os.WriteFile("output.txt", []byte(str), 0644)
}
