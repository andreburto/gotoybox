package main

import (
	"testing"
)

func TestCounter1(t *testing.T) {
	var expected []int = []int{0, 1, 2, 3, 4, 5}
	var actual []int = getRange(0, 5)

	if len(expected) != len(actual) {
		t.Errorf("Expected %d, but got %d", len(expected), len(actual))
	}
}

func TestCounter2(t *testing.T) {
	var first_test []string = []string{"0", "10"}
	var actual_start1, actual_end1 int = getStartEnd(first_test)
	
	if actual_start1 != 0 || actual_end1 != 10 {
		t.Errorf("Expected start 0 and end 10, but got start %d and end %d", actual_start1, actual_end1)
	}

	var second_test []string = []string{"5"}
	var actual_start2, actual_end2 int = getStartEnd(second_test)
	if actual_start2 != 0 || actual_end2 != 5 {
		t.Errorf("Expected start 0 and end 5, but got start %d and end %d", actual_start2, actual_end2)
	}
}