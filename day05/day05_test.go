package day05_test

import (
	"adventofcode21/day05"
	"strconv"
	"testing"
)

func TestDay05(t *testing.T) {
	result := day05.Day05("input_test.txt")
	t.Log("Day 05 Result is: " + strconv.Itoa(result))
	if result != 5 {
		t.Error("Day 05 does not run correctly!")
	}
}

func TestDay0502(t *testing.T) {
	result := day05.Day0502("input_test.txt")
	t.Log("Day 05 Part 2 Result is: " + strconv.Itoa(result))
	if result != 12 {
		t.Error("Day 05 Part 2 does not run correctly!")
	}
}
