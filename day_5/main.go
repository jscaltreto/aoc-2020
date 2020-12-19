package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type seat struct {
	row int
	col int
}

func (s *seat) id() int {
	return 8*s.row + s.col
}

type seats []seat

func main() {
	var s seats
	loadInput(&s)
	a(s)
	b(s)
}

func loadInput(s *seats) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		i := 0
		for _, r := range line {
			i <<= 1
			if string(r) == "B" || string(r) == "R" {
				i |= 1
			}
		}
		*s = append(*s, seat{(i >> 3), (i & 0b111)})
	}
}

func a(s seats) {
	id := 0
	for _, i := range s {
		if i.id() > id {
			id = i.id()
		}
	}
	fmt.Printf("Max ID: %d\n", id)
}

func b(s seats) {
	id := 0
	var lastSeats uint = 0
CHECK:
	for id = 0; id < (127*8 + 7); id++ {
		lastSeats <<= 1
		for _, i := range s {
			if i.id() == id {
				lastSeats |= 1
				if (lastSeats & 0b111) == 0b101 {
					break CHECK
				}
				continue CHECK
			}
		}

	}
	fmt.Printf("Your Seat: %d\n", id-1)
}
