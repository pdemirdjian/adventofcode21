package day01

import (
	"adventofcode21/utility"
	"log"
	"strconv"
)

func Day01(inputFileName string) int {
	items := utility.GetArrayOfInputFile(inputFileName)
	nums := []int{}

	for _, current := range items {
		currentInt, err := strconv.Atoi(current)
		if err != nil {
			log.Fatalln(err)
		}
		nums = append(nums, currentInt)
	}

	larger := calculateLargerThanPrevious(nums)

	return larger

}

func Day0102(inputFileName string) int {
	items := utility.GetArrayOfInputFile(inputFileName)

	first := -1
	second := -1
	sums := []int{}

	for _, current := range items {
		third, err := strconv.Atoi(current)
		if err != nil {
			log.Fatalln(err)
		}
		if first != -1 && second != -1 {
			totalOfWindow := first + second + third
			sums = append(sums, totalOfWindow)
		}
		first = second
		second = third
	}
	larger := calculateLargerThanPrevious(sums)
	return larger
}

func calculateLargerThanPrevious(list []int) int {
	larger := 0
	previous := -1

	for _, current := range list {
		if current > previous && previous != -1 {
			larger = larger + 1
		}
		previous = current
	}

	return larger
}
