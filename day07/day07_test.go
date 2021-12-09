package day07_test

import (
	"adventofcode21/day07"
	"strconv"
	"testing"
)

func TestDay07(t *testing.T) {
	result := day07.Day07("input_test.txt")
	t.Log("Day 07 Result is: " + strconv.Itoa(result))
	if result != 37 {
		t.Error("Day 07 does not run correctly!")
	}
}

func TestDay0702(t *testing.T) {
	result := day07.Day0702("input_test.txt")
	t.Log("Day 07 Part 2 Result is: " + strconv.Itoa(result))
	if result != 168 {
		t.Error("Day 07 Part 2 does not run correctly!")
	}
}
