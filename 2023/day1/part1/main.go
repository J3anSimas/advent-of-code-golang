package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		res, err := getDigitsPart2(line)
		if err != nil {
			panic(err)
		}
		sum += res
	}
	fmt.Println(sum)

}
func getDigitsPart1(input string) (int, error) {
	var firstDigit string
	var lastDigit string
	firstDigitFount := false
	for _, char := range input {
		if char >= '0' && char <= '9' {
			if !firstDigitFount {
				firstDigit = string(char)
				firstDigitFount = true
			}
			lastDigit = string(char)
		}
	}
	digit, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		return 0, err
	}
	return digit, nil
}
func getDigitsPart2(input string) (int, error) {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	type info struct {
		is_char bool
		value   string
	}
	indexes := make([]info, len(input))

	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			indexes[i] = info{is_char: true, value: string(input[i])}
		}
	}
	for num, n := range numbers {
		strNum := strconv.Itoa(num + 1)
		for _, i := range findOcurrency(input, n) {
			indexes[i] = info{is_char: true, value: strNum}
		}

	}
	var firstDigit string
	var lastDigit string
	firstDigitFount := false
	for _, char := range indexes {
		if char.is_char == true {
			if !firstDigitFount {
				firstDigit = string(char.value)
				firstDigitFount = true
			}
			lastDigit = string(char.value)
		}
	}
	digit, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		return 0, err
	}
	return digit, nil

}
func findOcurrency(input string, search string) []int {
	indexes := []int{}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(search); j++ {
			if i+j >= len(input) {
				break
			}
			if input[i+j] != search[j] {
				break
			}
			if j == len(search)-1 {
				indexes = append(indexes, i)
			}
		}
	}
	return indexes
}
