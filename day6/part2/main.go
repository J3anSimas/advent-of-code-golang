package main

import (
	"fmt"
	"log"
	"os"
)

func Run(filename string) int {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file")
	}
	inputData := string(raw_data)
	raw_data = make([]byte, 0)
	var chunk string
	for i := 0; i < len(inputData); i++ {
		if len(chunk) == 14 {
			chunk = chunk[1:]
		}
		chunk += string(inputData[i])

		if len(chunk) == 14 {
			if !checkRepetitiveChars(chunk) {
				return i + 1
			}
		}
	}
	return 0
}
func checkRepetitiveChars(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}
	return false
}
func main() {
	// fmt.Println(checkRepetitiveChars("mjqj"))
	// fmt.Println(checkRepetitiveChars("jpqm"))
	fmt.Println(Run("../input.txt"))
}
