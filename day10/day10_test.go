package day10_test

import (
	"adventofcode21/day10"
	"strconv"
	"testing"
)

func TestDay10(t *testing.T) {
	result := day10.Day10("input_test.txt")
	t.Log("Day 09 Result is: " + strconv.Itoa(result))
	if result != 26397 {
		t.Error("Day 09 does not run correctly!")
	}
}
