package main

import "testing"

func TestOverlaps(t *testing.T) {
	want := false
	first := []int{2, 4}
	second := []int{6, 8}
	got := checkIfOverlaps(first, second)
	if want != got {
		t.Errorf("checkIfOverlaps([2,4], [6,8]) = %t, want = %t", got, want)
	}
	want = false
	first = []int{2, 3}
	second = []int{4, 5}
	got = checkIfOverlaps(first, second)
	if want != got {
		t.Errorf("checkIfOverlaps([2,3], [4,5]) = %t, want = %t", got, want)
	}
	want = true
	first = []int{5, 7}
	second = []int{7, 9}
	got = checkIfOverlaps(first, second)
	if want != got {
		t.Errorf("checkIfOverlaps([5,7], [7,9]) = %t, want = %t", got, want)
	}
	want = true
	first = []int{2, 8}
	second = []int{3, 7}
	got = checkIfOverlaps(first, second)
	if want != got {
		t.Errorf("checkIfOverlaps([2,8], [3,7]) = %t, want = %t", got, want)
	}
	want = true
	first = []int{6, 6}
	second = []int{4, 6}
	got = checkIfOverlaps(first, second)
	if want != got {
		t.Errorf("checkIfOverlaps([6,6], [4,6]) = %t, want = %t", got, want)
	}
	want = true
	first = []int{2, 6}
	second = []int{4, 8}
	got = checkIfOverlaps(first, second)
	if want != got {
		t.Errorf("checkIfOverlaps([2,6], [4,8]) = %t, want = %t", got, want)
	}

}
