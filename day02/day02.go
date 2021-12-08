package day02

import (
	"adventofcode21/utility"
	"log"
	"strconv"
	"strings"
)

type Locaion struct {
	horizonal int
	depth     int
	aim       int
}

type Input struct {
	direction string
	length    int
}

func Day02(fileInputPath string) int {
	inputs := getInput(fileInputPath)

	location := Locaion{
		horizonal: 0,
		depth:     0,
	}

	for _, input := range inputs {
		if input.direction == "forward" {
			location.horizonal = location.horizonal + input.length
		} else if input.direction == "down" {
			location.depth = location.depth + input.length
		} else {
			location.depth = location.depth - input.length
		}
	}

	return location.depth * location.horizonal
}

func Day0202(fileInputPath string) int {
	inputs := getInput(fileInputPath)

	location := Locaion{
		horizonal: 0,
		depth:     0,
		aim:       0,
	}

	for _, input := range inputs {
		if input.direction == "forward" {
			location.horizonal = location.horizonal + input.length
			location.depth = location.depth + (location.aim * input.length)
		} else if input.direction == "down" {
			location.aim = location.aim + input.length
		} else {
			location.aim = location.aim - input.length
		}
	}

	return location.depth * location.horizonal
}

func getInput(fileInputPath string) []Input {
	lines := utility.GetArrayOfInputFile(fileInputPath)

	inputs := []Input{}

	for _, line := range lines {
		splits := strings.SplitN(line, " ", 2)
		length, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatalln(err)
		}
		lineInput := Input{
			direction: splits[0],
			length:    length,
		}
		inputs = append(inputs, lineInput)
	}

	return inputs
}
