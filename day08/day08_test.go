package day08_test

import (
	"adventofcode21/day08"
	"strconv"
	"testing"
)

func TestDay08(t *testing.T) {
	result := day08.Day08("input_test.txt")
	t.Log("Day 08 Result is: " + strconv.Itoa(result))
	if result != 26 {
		t.Error("Day 08 does not run correctly!")
	}
}
