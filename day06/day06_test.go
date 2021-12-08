package day06_test

import (
	"adventofcode21/day06"
	"strconv"
	"testing"
)

func TestDay06(t *testing.T) {
	result := day06.Day06("input_test.txt")
	t.Log("Day 06 Result is: " + strconv.Itoa(result))
	if result != 5934 {
		t.Error("Day 06 does not run correctly!")
	}
}

func TestDay0602(t *testing.T) {
	result := day06.Day0602("input_test.txt")
	t.Log("Day 06 Part 2 Result is: " + strconv.Itoa(result))
	if result != 26984457539 {
		t.Error("Day 06 Part 2 does not run correctly!")
	}
}
