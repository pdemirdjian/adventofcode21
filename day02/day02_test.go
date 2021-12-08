package day02_test

import (
	"adventofcode21/day02"
	"strconv"
	"testing"
)

func TestDay02(t *testing.T) {
	result := day02.Day02("input_test.txt")
	t.Log("Day 02 Result is: " + strconv.Itoa(result))
	if result != 150 {
		t.Error("Day 02 does not run correctly!")
	}
}

func TestDay0202(t *testing.T) {
	result := day02.Day0202("input_test.txt")
	t.Log("Day 02 Part 2 Result is: " + strconv.Itoa(result))
	if result != 900 {
		t.Error("Day 02 Part 2 does not run correctly!")
	}
}
