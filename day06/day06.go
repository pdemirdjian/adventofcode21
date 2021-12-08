package day06

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Sea struct {
	day  int
	fish [9]int
}

func Day06(inputFileName string) int {
	sea := processInputToSea(inputFileName)
	seaTotal := processSeaToFishByDays(80, sea)
	return seaTotal
}

func Day0602(inputFileName string) int {
	sea := processInputToSea(inputFileName)
	seaTotal := processSeaToFishByDays(256, sea)
	return seaTotal
}

func processSeaToFishByDays(days int, sea Sea) int {
	for day := 1; day <= days; day++ {
		for i, numFish := range sea.fish {
			if i == 0 {
				sea.fish[0] = sea.fish[0] - numFish
				sea.fish[6] = sea.fish[6] + numFish
				sea.fish[8] = sea.fish[8] + numFish
			} else {
				sea.fish[i-1] = sea.fish[i-1] + numFish
				sea.fish[i] = sea.fish[i] - numFish
			}
		}
	}
	total := 0
	for _, numFish := range sea.fish {
		total = total + numFish
	}
	return total
}

func processInputToSea(inputFileName string) Sea {
	byteContents, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	contents := string(byteContents)

	stringParts := strings.Split(contents, ",")

	fishes := [9]int{}

	for _, fishString := range stringParts {
		fishTimer, err := strconv.Atoi(fishString)
		if err != nil {
			log.Fatalln(err)
		}
		fishes[fishTimer]++
	}
	sea := Sea{
		day:  0,
		fish: fishes,
	}
	return sea
}
