package day09

import (
	"adventofcode21/utility"
	"log"
	"strconv"
)

type Point struct {
	up     int
	down   int
	left   int
	right  int
	height int
	risk   int
	low    bool
}

func Day09(fileInputName string) int {
	fullRows := processToPointArray(utility.GetArrayOfInputFile(fileInputName))
	fullRows = processPoints(fullRows)
	sum := 0
	for _, rows := range fullRows {
		for _, point := range rows {
			if point.low {
				sum = sum + point.risk
			}
		}
	}
	return sum
}

func processPoints(points [][]Point) [][]Point {
	maxY := len(points)
	maxX := len(points[0])
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			point := points[y][x]
			if y != 0 {
				point.up = points[y-1][x].height
			}
			if y < maxY-1 {
				point.down = points[y+1][x].height
			}
			if x != 0 {
				point.left = points[y][x-1].height
			}
			if x < maxX-1 {
				point.right = points[y][x+1].height
			}
			point.risk = point.height + 1
			points[y][x] = point
		}
	}
	for y, pointsRows := range points {
		for x, point := range pointsRows {
			if point.height < point.left && point.height < point.up && point.height < point.right && point.height < point.down {
				point.low = true
				points[y][x] = point
			}
		}
	}
	return points
}

func processToPointArray(rowsStrings []string) [][]Point {
	full := [][]Point{}
	for _, rowString := range rowsStrings {
		row := []Point{}
		for _, colChar := range rowString {
			actualNum, err := strconv.Atoi(string(colChar))
			if err != nil {
				log.Fatalln(err)
			}
			point := Point{
				height: actualNum,
				up:     10,
				down:   10,
				left:   10,
				right:  10,
				risk:   actualNum + 1,
				low:    false,
			}
			row = append(row, point)
		}
		full = append(full, row)
	}
	return full
}
