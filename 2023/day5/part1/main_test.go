package main

import "testing"

func TestFindInTableWithNoValueFound(t *testing.T) {
	value := 14
	lines := []string{"50 98 2", "52 50 48"}
	want := 14
	got := findInTable(value, lines)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
func TestFindInTableWithValueFound(t *testing.T) {
	value := 81
	lines := []string{"0 15 37", "37 52 2", "39 0 15"}
	want := 81
	got := findInTable(value, lines)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
