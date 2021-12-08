package day01_test

import (
	"adventofcode21/day01"
	"strconv"
	"testing"
)

func TestDay01(t *testing.T) {
	result := day01.Day01("input_test.txt")
	t.Log("Day01 Result is: " + strconv.Itoa(result))
	if result != 7 {
		t.Error("Day01 does not test correctly!")
	}
}

func TestDay0101(t *testing.T) {
	result := day01.Day0102("input_test.txt")
	t.Log("Day01 Part 2 Result is: " + strconv.Itoa(result))
	if result != 5 {
		t.Error("Day01 Part 2 does not test correctly!")
	}

}
