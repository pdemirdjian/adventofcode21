package day05

import (
	"adventofcode21/utility"
	"log"
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x                int
	y                int
	linesOverlapping int
}

type Line struct {
	startingCoordinate Coordinate
	endingCoordinate   Coordinate
}

type Grid struct {
	coordinates [][]Coordinate
}

func Day05(inputFileName string) int {
	lines := turnInputIntoLines(inputFileName)
	largestX := largestX(lines)
	largestY := largestY(lines)
	gridCoordinates := [][]Coordinate{}
	for x := 0; x < largestX+1; x++ {
		var row []Coordinate = make([]Coordinate, largestY+1)
		for y := 0; y < largestY+1; y++ {
			coordinate := Coordinate{
				x: x,
				y: y,
			}
			row[y] = coordinate
		}
		gridCoordinates = append(gridCoordinates, row)
	}
	grid := Grid{
		coordinates: gridCoordinates,
	}
	for _, line := range lines {
		if line.startingCoordinate.x == line.endingCoordinate.x {
			x := line.startingCoordinate.x
			smallerY := 0
			biggerY := 0
			if line.startingCoordinate.y < line.endingCoordinate.y {
				smallerY = line.startingCoordinate.y
				biggerY = line.endingCoordinate.y
			} else {
				smallerY = line.endingCoordinate.y
				biggerY = line.startingCoordinate.y
			}
			for i := smallerY; i <= biggerY; i++ {
				currentCoordinate := grid.coordinates[x][i]
				currentCoordinate.x = x
				currentCoordinate.y = i
				currentCoordinate.linesOverlapping = currentCoordinate.linesOverlapping + 1
				grid.coordinates[x][i] = currentCoordinate
			}
		} else if line.startingCoordinate.y == line.endingCoordinate.y {
			y := line.startingCoordinate.y
			smallerX := 0
			biggerX := 0
			if line.startingCoordinate.x < line.endingCoordinate.x {
				smallerX = line.startingCoordinate.x
				biggerX = line.endingCoordinate.x
			} else {
				smallerX = line.endingCoordinate.x
				biggerX = line.startingCoordinate.x
			}
			for i := smallerX; i <= biggerX; i++ {
				currentCoordinate := grid.coordinates[i][y]
				currentCoordinate.x = i
				currentCoordinate.y = y
				currentCoordinate.linesOverlapping = currentCoordinate.linesOverlapping + 1
				grid.coordinates[i][y] = currentCoordinate
			}
		}
	}
	sumOverlappingMoreThan2 := 0
	for _, coordinateRows := range grid.coordinates {
		for _, coordinate := range coordinateRows {
			if coordinate.linesOverlapping >= 2 {
				sumOverlappingMoreThan2 = sumOverlappingMoreThan2 + 1
			}
		}
	}
	return sumOverlappingMoreThan2
}

func Day0502(inputFileName string) int {
	lines := turnInputIntoLines(inputFileName)
	largestX := largestX(lines)
	largestY := largestY(lines)
	gridCoordinates := [][]Coordinate{}
	for x := 0; x < largestX+1; x++ {
		var row []Coordinate = make([]Coordinate, largestY+1)
		for y := 0; y < largestY+1; y++ {
			coordinate := Coordinate{
				x: x,
				y: y,
			}
			row[y] = coordinate
		}
		gridCoordinates = append(gridCoordinates, row)
	}
	grid := Grid{
		coordinates: gridCoordinates,
	}
	for _, line := range lines {
		if line.startingCoordinate.x == line.endingCoordinate.x {
			x := line.startingCoordinate.x
			smallerY := 0
			biggerY := 0
			if line.startingCoordinate.y < line.endingCoordinate.y {
				smallerY = line.startingCoordinate.y
				biggerY = line.endingCoordinate.y
			} else {
				smallerY = line.endingCoordinate.y
				biggerY = line.startingCoordinate.y
			}
			for i := smallerY; i <= biggerY; i++ {
				currentCoordinate := grid.coordinates[x][i]
				currentCoordinate.x = x
				currentCoordinate.y = i
				currentCoordinate.linesOverlapping = currentCoordinate.linesOverlapping + 1
				grid.coordinates[x][i] = currentCoordinate
			}
		} else if line.startingCoordinate.y == line.endingCoordinate.y {
			y := line.startingCoordinate.y
			smallerX := 0
			biggerX := 0
			if line.startingCoordinate.x < line.endingCoordinate.x {
				smallerX = line.startingCoordinate.x
				biggerX = line.endingCoordinate.x
			} else {
				smallerX = line.endingCoordinate.x
				biggerX = line.startingCoordinate.x
			}
			for i := smallerX; i <= biggerX; i++ {
				currentCoordinate := grid.coordinates[i][y]
				currentCoordinate.x = i
				currentCoordinate.y = y
				currentCoordinate.linesOverlapping = currentCoordinate.linesOverlapping + 1
				grid.coordinates[i][y] = currentCoordinate
			}
		} else {
			diffY := line.endingCoordinate.y - line.startingCoordinate.y
			diffX := line.endingCoordinate.x - line.startingCoordinate.x
			rateX := diffX / int(math.Abs(float64(diffX)))
			rateY := diffY / int(math.Abs(float64(diffY)))
			for i := 0; i <= int(math.Abs(float64(diffX))); i++ {
				currentCoordinate := gridCoordinates[(i*rateX)+line.startingCoordinate.x][(i*rateY)+line.startingCoordinate.y]
				currentCoordinate.linesOverlapping = currentCoordinate.linesOverlapping + 1
				grid.coordinates[(i*rateX)+line.startingCoordinate.x][(i*rateY)+line.startingCoordinate.y] = currentCoordinate
			}
		}
	}
	sumOverlappingMoreThan2 := 0
	for _, coordinateRows := range grid.coordinates {
		for _, coordinate := range coordinateRows {
			if coordinate.linesOverlapping >= 2 {
				sumOverlappingMoreThan2 = sumOverlappingMoreThan2 + 1
			}
		}
	}
	return sumOverlappingMoreThan2
}

func largestY(lines []Line) int {
	largest := 0
	for _, line := range lines {
		if line.startingCoordinate.y > largest && line.startingCoordinate.y > line.endingCoordinate.y {
			largest = line.startingCoordinate.y
		} else if line.endingCoordinate.y > largest {
			largest = line.endingCoordinate.y
		}
	}
	return largest
}

func largestX(lines []Line) int {
	largest := 0
	for _, line := range lines {
		if line.startingCoordinate.x > largest && line.startingCoordinate.x > line.endingCoordinate.x {
			largest = line.startingCoordinate.x
		} else if line.endingCoordinate.x > largest {
			largest = line.endingCoordinate.x
		}
	}
	return largest
}

func turnInputIntoLines(inputFileName string) []Line {
	inputs := utility.GetArrayOfInputFile(inputFileName)
	lines := []Line{}

	for _, input := range inputs {
		coordinates := strings.Split(input, " -> ")
		startingCoordinatesSlice := strings.Split(coordinates[0], ",")
		startingCoordinatesXStr := startingCoordinatesSlice[0]
		startingCoordinatesYStr := startingCoordinatesSlice[1]
		startingCoordinatesX, err := strconv.Atoi(startingCoordinatesXStr)
		if err != nil {
			log.Fatalln(err)
		}
		startingCoordinatesY, err := strconv.Atoi(startingCoordinatesYStr)
		if err != nil {
			log.Fatalln(err)
		}
		startingCoordinate := Coordinate{
			x: startingCoordinatesX,
			y: startingCoordinatesY,
		}
		endingCoordinatesSlice := strings.Split(coordinates[1], ",")
		endingCoordinatesXStr := endingCoordinatesSlice[0]
		endingCoordinatesYStr := endingCoordinatesSlice[1]
		endingCoordinatesX, err := strconv.Atoi(endingCoordinatesXStr)
		if err != nil {
			log.Fatalln(err)
		}
		endingCoordinatesY, err := strconv.Atoi(endingCoordinatesYStr)
		if err != nil {
			log.Fatalln(err)
		}
		endingCoordinate := Coordinate{
			x: endingCoordinatesX,
			y: endingCoordinatesY,
		}
		line := Line{
			startingCoordinate: startingCoordinate,
			endingCoordinate:   endingCoordinate,
		}
		lines = append(lines, line)
	}
	return lines
}
