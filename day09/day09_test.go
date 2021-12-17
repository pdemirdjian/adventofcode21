package day09_test

import (
	"adventofcode21/day09"
	"strconv"
	"testing"
)

func TestDay09(t *testing.T) {
	result := day09.Day09("input_test.txt")
	t.Log("Day 09 Result is: " + strconv.Itoa(result))
	if result != 15 {
		t.Error("Day 09 does not run correctly!")
	}
}
