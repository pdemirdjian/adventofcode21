package day08

import (
	"adventofcode21/utility"
	"strings"
)

type Line struct {
	numSegments int
	normLetters []string
	normString  string
	value       int
}

type Display struct {
	pattern    []string
	output     []string
	a          string
	b          string
	c          string
	d          string
	e          string
	f          string
	g          string
	newOutput  []string
	newPattern []string
	total      int
}

func Day08(fileInputName string) int {
	lineArray := setupLineArray()
	displays := setupInput(fileInputName)
	totalSelected := 0
	for _, display := range displays {
		for _, output := range display.output {
			length := len(output)
			if length == lineArray[1].numSegments || length == lineArray[4].numSegments || length == lineArray[7].numSegments || length == lineArray[8].numSegments {
				totalSelected++
			}
		}
	}
	return totalSelected
}

func setupInput(fileInputName string) []Display {
	displays := []Display{}
	for _, line := range utility.GetArrayOfInputFile(fileInputName) {
		stringSplits := strings.Split(line, "|")
		pattern := stringSplits[0]
		output := stringSplits[1]
		patternSplits := strings.Fields(pattern)
		outputSplits := strings.Fields(output)
		display := Display{
			pattern: patternSplits,
			output:  outputSplits,
		}
		displays = append(displays, display)
	}
	return displays
}

func setupLineArray() [10]Line {
	lines := [10]Line{
		{
			numSegments: 6,
			normLetters: []string{"a", "b", "c", "e", "f", "g"},
			normString:  "abcefg",
			value:       0,
		},
		{
			numSegments: 2,
			normLetters: []string{"c", "f"},
			normString:  "cf",
			value:       1,
		},
		{
			numSegments: 5,
			normLetters: []string{"a", "c", "d", "e", "g"},
			normString:  "acdeg",
			value:       2,
		},
		{
			numSegments: 5,
			normLetters: []string{"a", "c", "d", "e", "g"},
			normString:  "acdeg",
			value:       3,
		},
		{
			numSegments: 4,
			normLetters: []string{"b", "c", "d", "f"},
			normString:  "bcdf",
			value:       4,
		},
		{
			numSegments: 5,
			normLetters: []string{"a", "b", "d", "f", "g"},
			normString:  "abdfg",
			value:       5,
		},
		{
			numSegments: 6,
			normLetters: []string{"a", "b", "d", "e", "f", "g"},
			normString:  "abdefg",
			value:       6,
		},
		{
			numSegments: 3,
			normLetters: []string{"a", "c", "f"},
			normString:  "acf",
			value:       7,
		},
		{
			numSegments: 7,
			normLetters: []string{"a", "b", "c", "d", "e", "f", "g"},
			normString:  "abcdefg",
			value:       8,
		},
		{
			numSegments: 6,
			normLetters: []string{"a", "b", "c", "d", "f", "g"},
			normString:  "abcdfg",
			value:       9,
		},
	}
	return lines
}
