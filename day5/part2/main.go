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

func (s *Stack) Push(items []string) {
	s.items = append(s.items, items...)
}

func (s *Stack) Pop(quantity int) []string {
	if len(s.items) == 0 {
		return []string{}
	}
	items := s.items[len(s.items)-quantity:]
	s.items = s.items[:len(s.items)-quantity]
	return items
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
	result := ""
	for _, stack := range stacks {
		result += strings.Join(stack.Pop(1), "")
	}
	fmt.Print(result)
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
				stacks[i/4].Push([]string{string((line[i+1]))})
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
	items := stacks[source-1].Pop(amount)
	stacks[dest-1].Push(items)
	return stacks
}
func main() {
	Run("../input.txt")
}
