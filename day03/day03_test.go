package day03_test

import (
	"adventofcode21/day03"
	"strconv"
	"testing"
)

func TestDay03(t *testing.T) {
	result := day03.Day03("input_test.txt")
	t.Log("Day 03 Result is: " + strconv.Itoa(result))
	if result != 198 {
		t.Error("Day 03 does not run correctly!")
	}
}

func TestDay0302(t *testing.T) {
	result := day03.Day0302("input_test.txt")
	t.Log("Day 03 Result is: " + strconv.Itoa(result))
	if result != 230 {
		t.Error("Day 03 does not run correctly!")
	}
}
