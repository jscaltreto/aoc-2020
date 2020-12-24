package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

const (
	INPUT = "input.txt"
)

type Coord struct {
	x int
	y int
}

func (c *Coord) getAllNeighbors() map[string]Coord {
	return map[string]Coord{
		"nw": Coord{c.x - 1, c.y - 1},
		"ne": Coord{c.x, c.y - 1},
		"sw": Coord{c.x, c.y + 1},
		"se": Coord{c.x + 1, c.y + 1},
		"w":  Coord{c.x - 1, c.y},
		"e":  Coord{c.x + 1, c.y},
	}
}

type Floor map[Coord]bool

var floor Floor

var instructions []string

func main() {
	floor = make(Floor)
	loadInput()
	a()
	b()
}

func loadInput() {
	data, err := ioutil.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	instructions = strings.Split(string(data), "\n")
}

func a() {
	start := Coord{0, 0}
	floor[start] = false
	re := regexp.MustCompile(`(se|sw|ne|nw|w|e)`)
	for _, line := range instructions {
		match := re.FindAllStringSubmatch(line, -1)
		if match == nil {
			log.Fatalf("Invalid instructions: %s", line)
		}
		cursor := start
		for _, m := range match {
			cursor = cursor.getAllNeighbors()[m[1]]
		}
		if _, ok := floor[cursor]; !ok {
			floor[cursor] = true
		} else {
			floor[cursor] = !floor[cursor]
		}
	}
	blackTiles := 0
	for _, v := range floor {
		if v {
			blackTiles++
		}
	}
	fmt.Println("Part A:", blackTiles)
}

func flipFloor() {
	newFloor := make(Floor)
	checked := make(Floor)
	tilesToCheck := []Coord{}
	for c, _ := range floor {
		tilesToCheck = append(tilesToCheck, c)
	}
	for len(tilesToCheck) > 0 {
		c := tilesToCheck[0]
		tilesToCheck = tilesToCheck[1:]
		if _, ok := checked[c]; ok {
			continue
		}
		checked[c] = true
		v := floor[c]
		neighbors := 0
		for _, n := range c.getAllNeighbors() {
			if b := floor[n]; b {
				neighbors++
			}
			if v {
				tilesToCheck = append(tilesToCheck, n)
			}
		}
		if v && (neighbors == 0 || neighbors > 2) {
			newFloor[c] = false
		} else if !v && neighbors == 2 {
			newFloor[c] = true
		} else {
			newFloor[c] = v
		}
	}
	floor = newFloor
}

func b() {
	for d := 1; d <= 100; d++ {
		flipFloor()
	}
	blackTiles := 0
	for _, v := range floor {
		if v {
			blackTiles++
		}
	}
	fmt.Println("Part B:", blackTiles)
}
