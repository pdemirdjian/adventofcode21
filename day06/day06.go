package day06

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Sea struct {
	day  int
	fish []Fish
}

type Fish struct {
	dayCreated int
	timer      int
	fishNumber int
}

func Day06(inputFileName string) int {
	sea := processInputToSea(inputFileName)
	sea = processSeaToFishByDays(80, sea)
	return len(sea.fish)
}

func Day0602(inputFileName string) int {
	sea := processInputToSea(inputFileName)
	sea = processSeaToFishByDays(256, sea)
	return len(sea.fish)
}

func processSeaToFishByDays(days int, sea Sea) Sea {
	for day := 1; day <= days; day++ {
		fishes := sea.fish
		totalFish := len(fishes)
		for fishNumber, fish := range fishes {
			if fish.timer == 0 {
				fish.timer = 6
				totalFish = totalFish + 1
				newFish := Fish{
					dayCreated: day,
					timer:      8,
					fishNumber: totalFish,
				}
				fishes = append(fishes, newFish)
			} else {
				fish.timer = fish.timer - 1
			}
			fishes[fishNumber] = fish
		}
		sea.fish = fishes
	}
	return sea
}

func processInputToSea(inputFileName string) Sea {
	sea := Sea{
		day: 0,
	}
	byteContents, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	contents := string(byteContents)

	stringParts := strings.Split(contents, ",")

	fishes := []Fish{}

	for i, fishString := range stringParts {
		fishTimer, err := strconv.Atoi(fishString)
		if err != nil {
			log.Fatalln(err)
		}
		fish := Fish{
			dayCreated: 0,
			timer:      fishTimer,
			fishNumber: i,
		}
		fishes = append(fishes, fish)
	}

	sea.fish = fishes

	return sea
}
