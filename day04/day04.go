package day04

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	turn   int
	number int
}

type Entry struct {
	x      int
	y      int
	number int
	hit    bool
	turn   int
}

type BingoCard struct {
	entries [5][5]Entry
	bingo   bool
}

type Game struct {
	orders []Order
	cards  []BingoCard
}

func bingo(card BingoCard) bool {
	bingoColumn := checkColumns(card)
	bingoRow := checkRows(card)
	return bingoColumn || bingoRow
}

func checkRows(card BingoCard) bool {
	for y := 0; y < 5; y++ {
		if checkRow(card, y) {
			return true
		}
	}
	return false
}

func checkRow(card BingoCard, row int) bool {
	for x := 0; x < 5; x++ {
		if !card.entries[x][row].hit {
			return false
		}
	}
	return true
}

func checkColumns(card BingoCard) bool {
	for x := 0; x < 5; x++ {
		if checkColumn(card, x) {
			return true
		}
	}
	return false
}

func checkColumn(card BingoCard, column int) bool {
	for y := 0; y < 5; y++ {
		if !card.entries[column][y].hit {
			return false
		}
	}
	return true
}

func sumAllNotHit(card BingoCard) int {
	sumNotHit := 0
	for _, entryList := range card.entries {
		for _, entry := range entryList {
			if !entry.hit {
				sumNotHit = sumNotHit + entry.number
			}
		}
	}
	return sumNotHit
}

func Day04(inputFileName string) int {
	game := setupGame(inputFileName)
	winningNum := 0
	sumOfNotHit := 0
	for _, order := range game.orders {
		for i, card := range game.cards {
			for j, entryList := range card.entries {
				for k, entry := range entryList {
					if entry.number == order.number {
						entry.turn = order.turn
						entry.hit = true
						entryList[k] = entry
					}
				}
				card.entries[j] = entryList
			}
			game.cards[i] = card
		}
		for _, card := range game.cards {
			if bingo(card) {
				winningNum = order.number
				sumOfNotHit = sumAllNotHit(card)
				return winningNum * sumOfNotHit
			}
		}
	}
	return winningNum * sumOfNotHit
}

func Day0402(inputFileName string) int {
	game := setupGame(inputFileName)
	winningNum := 0
	sumOfNotHit := 0
	for _, order := range game.orders {
		for i, card := range game.cards {
			for j, entryList := range card.entries {
				for k, entry := range entryList {
					if entry.number == order.number {
						entry.turn = order.turn
						entry.hit = true
						entryList[k] = entry
					}
				}
				card.entries[j] = entryList
			}
			game.cards[i] = card
		}
		for i, card := range game.cards {
			if !card.bingo && bingo(card) {
				otherNotBingo := false
				for j, card2 := range game.cards {
					if i != j && !card2.bingo {
						otherNotBingo = true
					}
				}
				if !otherNotBingo {
					winningNum = order.number
					sumOfNotHit = sumAllNotHit(card)
					return winningNum * sumOfNotHit
				}
				card.bingo = true
				game.cards[i] = card
			}
		}
	}
	return winningNum * sumOfNotHit
}

func setupGame(inputFileName string) Game {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	game := Game{}

	scanner := bufio.NewScanner(file)
	ordersIn := false
	newCard := false
	cardRow := 0
	cards := []BingoCard{}
	card := BingoCard{}
	for scanner.Scan() {
		line := scanner.Text()
		if !ordersIn {
			game.orders = setupOrders(line)
			ordersIn = true
			newCard = true
		} else if newCard {
			if line != "" {
				cardRow = 0
				card = addRowToCard(BingoCard{}, cardRow, line)
				cardRow = cardRow + 1
				newCard = false
			}
		} else {
			if line == "" {
				newCard = true
			} else {
				card = addRowToCard(card, cardRow, line)
				cardRow = cardRow + 1
				if cardRow == 5 {
					cards = append(cards, card)
				}
			}
		}
	}
	game.cards = cards
	return game
}

func addRowToCard(card BingoCard, cardRow int, line string) BingoCard {
	row := strings.Fields(line)
	if len(row) != 5 {
		log.Fatalln("ISN'T 5 ROWS!")
	}
	for i, item := range row {
		num, err := strconv.Atoi(item)
		if err != nil {
			log.Fatalln(err)
		}
		entry := Entry{
			x:      i,
			y:      cardRow,
			number: num,
			hit:    false,
		}
		card.entries[i][cardRow] = entry
	}
	return card
}

func setupOrders(line string) []Order {
	orders := strings.Split(line, ",")
	returnOrders := []Order{}
	for i, order := range orders {
		intOrder, err := strconv.Atoi(order)
		if err != nil {
			log.Fatalln(err)
		}
		orderStruct := Order{
			turn:   i,
			number: intOrder,
		}
		returnOrders = append(returnOrders, orderStruct)
	}
	return returnOrders
}
