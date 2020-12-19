package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	INPUT  = "input.txt"
	ACTIVE = '#'
)

type Coord [4]int

type Grid map[Coord]bool
type NeighborMap map[Coord]int

var state Grid = make(Grid)

func main() {
	loadInput()
	a()
	b()
}

func loadInput() {
	data, err := ioutil.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	for x, row := range strings.Split(string(data), "\n") {
		for y, cube := range row {
			state[Coord{x, y, 0, 0}] = cube == ACTIVE
		}
	}
}

func updateNeighbors(neighborMap *NeighborMap, coord Coord, status bool) {
	if _, ok := (*neighborMap)[coord]; ok {
		if status {
			(*neighborMap)[coord]++
		}
	} else {
		if status {
			(*neighborMap)[coord] = 1
		} else {
			(*neighborMap)[coord] = 0
		}
	}

}

func getNextStateA(curState Grid) Grid {
	neighborMap := make(NeighborMap)
	for k, status := range curState {
		for x := k[0] - 1; x <= k[0]+1; x++ {
			for y := k[1] - 1; y <= k[1]+1; y++ {
				for z := k[2] - 1; z <= k[2]+1; z++ {
					coord := Coord{x, y, z, 0}
					if coord == k {
						continue
					}
					updateNeighbors(&neighborMap, coord, status)
				}
			}
		}
	}
	return updateState(neighborMap, curState)
}

func updateState(neighborMap NeighborMap, curState Grid) Grid {
	nextState := make(Grid)
	for k, countActive := range neighborMap {
		nextState[k] = curState[k]
		if curState[k] {
			if countActive != 2 && countActive != 3 {
				nextState[k] = false
			}
		} else {
			if countActive == 3 {
				nextState[k] = true
			}
		}
	}
	return nextState
}

func a() {
	curState := make(Grid)
	for c, status := range state {
		curState[c] = status
	}
	for i := 0; i < 6; i++ {
		curState = getNextStateA(curState)
	}
	active := 0
	for _, status := range curState {
		if status {
			active++

		}
	}
	fmt.Println("Part A:", active)
}

func getNextStateB(curState Grid) Grid {
	neighborMap := make(NeighborMap)
	for k, status := range curState {
		for x := k[0] - 1; x <= k[0]+1; x++ {
			for y := k[1] - 1; y <= k[1]+1; y++ {
				for z := k[2] - 1; z <= k[2]+1; z++ {
					for w := k[3] - 1; w <= k[3]+1; w++ {
						coord := Coord{x, y, z, w}
						if coord == k {
							continue
						}
						updateNeighbors(&neighborMap, coord, status)
					}
				}
			}

		}
	}
	return updateState(neighborMap, curState)
}

func b() {
	curState := make(Grid)
	for c, status := range state {
		curState[c] = status
	}
	for i := 0; i < 6; i++ {
		curState = getNextStateB(curState)
	}
	active := 0
	for _, status := range curState {
		if status {
			active++

		}
	}
	fmt.Println("Part B:", active)
}
