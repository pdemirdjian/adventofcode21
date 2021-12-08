package day07

import (
	"adventofcode21/utility"
	"math"

	"github.com/montanaflynn/stats"
)

func Day07(inputFileName string) int {
	crabs := utility.GetFloat64ArrFromStringArr(utility.GetCommaSeperatedStringsOneLine(inputFileName))
	median, _ := stats.Median(crabs)
	median, _ = stats.Round(median, 0)
	fuel := 0
	for _, crab := range crabs {
		fuel = fuel + int(math.Abs(crab-median))
	}
	return int(fuel)
}

func Day0702(inputFileName string) int {
	crabs := utility.GetFloat64ArrFromStringArr(utility.GetCommaSeperatedStringsOneLine(inputFileName))
	mean, _ := stats.Mean(crabs)
	meanFl := math.Floor(mean)
	meanCi := math.Ceil(mean)
	min, _ := stats.Min([]float64{calcFuelForExponentialExpensive(crabs, meanFl), calcFuelForExponentialExpensive(crabs, meanCi)})
	return int(min)
}

func calcFuelForExponentialExpensive(crabs []float64, mean float64) float64 {
	fuel := 0
	for _, crab := range crabs {
		diff := int(math.Abs(crab - mean))
		for i := 1; i <= diff; i++ {
			fuel = fuel + i
		}
	}
	return float64(fuel)
}
