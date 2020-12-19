package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	INPUT    = "input.txt"
	OCCUPIED = '#'
	EMPTY    = 'L'
	FLOOR    = '.'
	NEWLINE  = '\n'
)

type CountFunc func(string, int) int

var translations []int
var seats string

func main() {
	loadInput()
	countOccupied(countPartA, 4)
	countOccupied(countPartB, 5)
}

func loadInput() {
	data, err := ioutil.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	seats = string(data)
	rowWidth := len(strings.Split(seats, "\n")[0]) + 1
	translations = []int{
		-(rowWidth - 1),
		-rowWidth,
		-(rowWidth + 1),
		rowWidth - 1,
		rowWidth,
		rowWidth + 1,
		-1,
		1,
	}
}

func countPartA(state string, seatId int) int {
	numOccupied := 0
	for _, check := range translations {
		checkSeat := seatId + check
		if checkSeat < 0 || checkSeat >= len(state) || state[checkSeat] == '\n' {
			continue
		}
		if state[checkSeat] == OCCUPIED {
			numOccupied++
		}
	}
	return numOccupied
}

func countPartB(state string, seatId int) int {
	numOccupied := 0
	for _, check := range translations {
		checkSeat := seatId
		for {
			checkSeat += check
			if checkSeat < 0 || checkSeat >= len(state) || state[checkSeat] == '\n' {
				break
			}
			if state[checkSeat] == OCCUPIED {
				numOccupied++
				break
			} else if state[checkSeat] == EMPTY {
				break
			}
		}
	}
	return numOccupied
}

func countOccupied(countFunc CountFunc, checkOccupied int) {
	seatsOccupied := 0
	nextState := seats
	for {
		state := nextState
		for seatId, seat := range state {
			if seat == FLOOR || seat == NEWLINE {
				continue
			}
			numOccupied := countFunc(state, seatId)

			if numOccupied == 0 {
				seat = OCCUPIED
			} else if numOccupied >= checkOccupied {
				seat = EMPTY
			}
			if seat == OCCUPIED {
				seatsOccupied++
			}
			nextState = nextState[:seatId] + string(seat) + nextState[seatId+1:]
		}
		if state == nextState {
			break
		}
		seatsOccupied = 0
	}
	fmt.Println("Seats Occupied:", seatsOccupied)
}
