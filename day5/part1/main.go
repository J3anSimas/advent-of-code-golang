package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func Run(filename string) int {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file")
	}
	inputData := string(raw_data)
	raw_data = make([]byte, 0)
	raw_stacks_string, num_stacks := SeparateLastLine(strings.Split(inputData, "\n\n")[0])
	// fmt.Println(num_of_stacks)
	raw_instructions := strings.Split(inputData, "\n\n")[1]
	// raw_stacks := strings.Split(raw_stacks_string, " ")
	//fmt.Println(len(raw_stacks))
	stacks := ParseStacks(raw_stacks_string, num_stacks)
	for _, instruction := range strings.Split(raw_instructions, "\n") {
		stacks = ParseInstruction(instruction, stacks)
	}
	fmt.Println()
	for _, stack := range stacks {
		fmt.Print(stack.Pop())
	}
	return len(raw_instructions)
}
func SeparateLastLine(input string) (string, int) {
	last_line := strings.TrimSpace(input[strings.LastIndex(input, "\n")+1:])
	num, err := strconv.Atoi(last_line[strings.LastIndex(last_line, " ")+1:])
	if err != nil {
		log.Fatal("Error converting string to int")
	}
	return input[:strings.LastIndex(input, "\n")], num
}
func ParseStacks(input string, num_of_stacks int) []Stack {
	stacks := make([]Stack, num_of_stacks)
	lines := strings.Split(input, "\n")
	slices.Reverse(lines)
	for _, line := range lines {
		for i := 0; i < len(line); i += 4 {
			if line[i] == '[' {
				stacks[i/4].Push(string(line[i+1]))
			}
		}
	}
	return stacks
}
func ParseInstruction(input string, stacks []Stack) []Stack {
	tokens := strings.Split(input, " ")
	amount, _ := strconv.Atoi(tokens[1])
	source, _ := strconv.Atoi(tokens[3])
	dest, _ := strconv.Atoi(tokens[5])
	for i := 0; i < amount; i++ {
		item := stacks[source-1].Pop()
		stacks[dest-1].Push(item)
	}
	return stacks
}
func main() {
	Run("../input.txt")
}
