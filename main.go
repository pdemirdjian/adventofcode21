package main

import (
	"adventofcode21/day01"
	"adventofcode21/day02"
	"adventofcode21/day03"
	"adventofcode21/day04"
	"adventofcode21/day05"
	"adventofcode21/day06"
	"adventofcode21/day07"
	"adventofcode21/day08"
	"adventofcode21/day09"
	"adventofcode21/day10"
	"flag"
	"log"
	"strconv"
)

func main() {

	dayPtr := flag.Int("day", 0, "Which Days Code To Run")

	flag.Parse()

	switch *dayPtr {
	case 0:
		log.Println("Pete's Advent of Code for 2021!")
		log.Println("Running Day01 Code.")
		log.Println("Day01's result is: " + strconv.Itoa(day01.Day01("./day01/input.txt")))
		log.Println("Day01 Code ran fine!")
		log.Println("Running Day01 Part 2 Code.")
		log.Println("Day01 Part 2's result is: " + strconv.Itoa(day01.Day0102("./day01/input.txt")))
		log.Println("Day01 Part 2 Code ran fine!")
		log.Println("Running Day02 Code.")
		log.Println("Day02's result is: " + strconv.Itoa(day02.Day02("./day02/input.txt")))
		log.Println("Day02 Code ran fine!")
		log.Println("Running Day02 Part 2 Code.")
		log.Println("Day02 Part 2's result is: " + strconv.Itoa(day02.Day0202("./day02/input.txt")))
		log.Println("Day02 Part 2 Code ran fine!")
		log.Println("Running Day03 Code.")
		log.Println("Day03's result is: " + strconv.Itoa(day03.Day03("./day03/input.txt")))
		log.Println("Day03 Code ran fine!")
		log.Println("Running Day03 Part 2 Code.")
		log.Println("Day03 Part 2's result is: " + strconv.Itoa(day03.Day0302("./day03/input.txt")))
		log.Println("Day03 Part 2 Code ran fine!")
		log.Println("Running Day04 Code.")
		log.Println("Day04's result is: " + strconv.Itoa(day04.Day04("./day04/input.txt")))
		log.Println("Day04 Code ran fine!")
		log.Println("Running Day04 Part 2 Code.")
		log.Println("Day04 Part 2's result is: " + strconv.Itoa(day04.Day0402("./day04/input.txt")))
		log.Println("Day04 Part 2 Code ran fine!")
		log.Println("Running Day05 Code.")
		log.Println("Day05's result is: " + strconv.Itoa(day05.Day05("./day05/input.txt")))
		log.Println("Day05 Code ran fine!")
		log.Println("Running Day05 Part 2 Code.")
		log.Println("Day05 Part 2's result is: " + strconv.Itoa(day05.Day0502("./day05/input.txt")))
		log.Println("Day05 Part 2 Code ran fine!")
		log.Println("Running Day06 Code.")
		log.Println("Day06's result is: " + strconv.Itoa(day06.Day06("./day06/input.txt")))
		log.Println("Day06 Code ran fine!")
		log.Println("Running Day06 Part 2 Code.")
		log.Println("Day06 Part 2's result is: " + strconv.Itoa(day06.Day0602("./day06/input.txt")))
		log.Println("Day06 Part 2 Code ran fine!")
		log.Println("Running Day07 Code.")
		log.Println("Day07's result is: " + strconv.Itoa(day07.Day07("./day07/input.txt")))
		log.Println("Day07 Code ran fine!")
		log.Println("Running Day07 Part 2 Code.")
		log.Println("Day07 Part 2's result is: " + strconv.Itoa(day07.Day0702("./day07/input.txt")))
		log.Println("Day07 Part 2 Code ran fine!")
		log.Println("Running Day08 Code.")
		log.Println("Day08's result is: " + strconv.Itoa(day08.Day08("./day08/input.txt")))
		log.Println("Day08 Code ran fine!")
		log.Println("Running Day09 Code.")
		log.Println("Day09's result is: " + strconv.Itoa(day09.Day09("./day09/input.txt")))
		log.Println("Day09 Code ran fine!")
		log.Println("Running Day10 Code.")
		log.Println("Day10's result is: " + strconv.Itoa(day10.Day10("./day10/input.txt")))
		log.Println("Day10 Code ran fine!")
	case 1:
		log.Println("Running Day01 Code.")
		log.Println("Day01's result is: " + strconv.Itoa(day01.Day01("./day01/input.txt")))
		log.Println("Day01 Code ran fine!")
		log.Println("Running Day01 Part 2 Code.")
		log.Println("Day01 Part 2's result is: " + strconv.Itoa(day01.Day0102("./day01/input.txt")))
		log.Println("Day01 Part 2 Code ran fine!")
	case 2:
		log.Println("Running Day02 Code.")
		log.Println("Day02's result is: " + strconv.Itoa(day02.Day02("./day02/input.txt")))
		log.Println("Day02 Code ran fine!")
		log.Println("Running Day02 Part 2 Code.")
		log.Println("Day02 Part 2's result is: " + strconv.Itoa(day02.Day0202("./day02/input.txt")))
		log.Println("Day02 Part 2 Code ran fine!")
	case 3:
		log.Println("Running Day03 Code.")
		log.Println("Day03's result is: " + strconv.Itoa(day03.Day03("./day03/input.txt")))
		log.Println("Day03 Code ran fine!")
		log.Println("Running Day03 Part 2 Code.")
		log.Println("Day03 Part 2's result is: " + strconv.Itoa(day03.Day0302("./day03/input.txt")))
		log.Println("Day03 Part 2 Code ran fine!")
	case 4:
		log.Println("Running Day04 Code.")
		log.Println("Day04's result is: " + strconv.Itoa(day04.Day04("./day04/input.txt")))
		log.Println("Day04 Code ran fine!")
		log.Println("Running Day04 Part 2 Code.")
		log.Println("Day04 Part 2's result is: " + strconv.Itoa(day04.Day0402("./day04/input.txt")))
		log.Println("Day04 Part 2 Code ran fine!")
	case 5:
		log.Println("Running Day05 Code.")
		log.Println("Day05's result is: " + strconv.Itoa(day05.Day05("./day05/input.txt")))
		log.Println("Day05 Code ran fine!")
		log.Println("Running Day05 Part 2 Code.")
		log.Println("Day05 Part 2's result is: " + strconv.Itoa(day05.Day0502("./day05/input.txt")))
		log.Println("Day05 Part 2 Code ran fine!")
	case 6:
		log.Println("Running Day06 Code.")
		log.Println("Day06's result is: " + strconv.Itoa(day06.Day06("./day06/input.txt")))
		log.Println("Day06 Code ran fine!")
		log.Println("Running Day06 Part 2 Code.")
		log.Println("Day06 Part 2's result is: " + strconv.Itoa(day06.Day0602("./day06/input.txt")))
		log.Println("Day06 Part 2 Code ran fine!")
	case 7:
		log.Println("Running Day07 Code.")
		log.Println("Day07's result is: " + strconv.Itoa(day07.Day07("./day07/input.txt")))
		log.Println("Day07 Code ran fine!")
		log.Println("Running Day07 Part 2 Code.")
		log.Println("Day07 Part 2's result is: " + strconv.Itoa(day07.Day0702("./day07/input.txt")))
		log.Println("Day07 Part 2 Code ran fine!")
	case 8:
		log.Println("Running Day08 Code.")
		log.Println("Day08's result is: " + strconv.Itoa(day08.Day08("./day08/input.txt")))
		log.Println("Day08 Code ran fine!")
	case 9:
		log.Println("Running Day09 Code.")
		log.Println("Day09's result is: " + strconv.Itoa(day09.Day09("./day09/input.txt")))
		log.Println("Day09 Code ran fine!")
	case 10:
		log.Println("Running Day10 Code.")
		log.Println("Day10's result is: " + strconv.Itoa(day10.Day10("./day10/input.txt")))
		log.Println("Day10 Code ran fine!")
	}

}
