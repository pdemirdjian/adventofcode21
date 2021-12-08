package day04_test

import (
	"adventofcode21/day04"
	"strconv"
	"testing"
)

func TestDay04(t *testing.T) {
	result := day04.Day04("input_test.txt")
	t.Log("Day 04 Result is: " + strconv.Itoa(result))
	if result != 4512 {
		t.Error("Day 04 does not run correctly!")
	}
}

func TestDay0402(t *testing.T) {
	result := day04.Day0402("input_test.txt")
	t.Log("Day 04 Result is: " + strconv.Itoa(result))
	if result != 1924 {
		t.Error("Day 04 does not run correctly!")
	}
}
