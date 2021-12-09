package day08_test

import (
	"adventofcode21/day08"
	"strconv"
	"testing"
)

func TestDay08(t *testing.T) {
	result := day08.Day08("input_test.txt")
	t.Log("Day 08 Result is: " + strconv.Itoa(result))
	if result != -1 {
		t.Error("Day 08 does not run correctly!")
	}
}

func TestDay0802(t *testing.T) {
	result := day08.Day0802("input_test.txt")
	t.Log("Day 08 Part 2 Result is: " + strconv.Itoa(result))
	if result != -1 {
		t.Error("Day 08 Part 2 does not run correctly!")
	}
}
