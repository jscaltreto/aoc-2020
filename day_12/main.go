package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	INPUT   = "input.txt"
	NORTH   = 'N'
	SOUTH   = 'S'
	EAST    = 'E'
	WEST    = 'W'
	LEFT    = 'L'
	RIGHT   = 'R'
	FORWARD = 'F'
)

var DIRECTIONS []byte = []byte{'N', 'E', 'S', 'W'}

var instructions string

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
	instructions = string(data)
}

func a() {
	var ship Ship = Ship{0, 0, 1}
	for _, inst := range strings.Split(instructions, "\n") {
		cmd := inst[0]
		val, _ := strconv.Atoi(inst[1:])
		ship.move(cmd, val)
	}
	fmt.Println("Part A:", math.Abs(float64(ship.x))+math.Abs(float64(ship.y)))

}

func b() {
	var ship Ship = Ship{0, 0, 1}
	var waypoint Waypoint = Waypoint{10, 1}
	for _, inst := range strings.Split(instructions, "\n") {
		cmd := inst[0]
		dist, _ := strconv.Atoi(inst[1:])
		switch cmd {
		case RIGHT:
			waypoint.turnRight(dist)
		case LEFT:
			waypoint.turnLeft(dist)
		case NORTH:
			waypoint.y += dist
		case SOUTH:
			waypoint.y -= dist
		case EAST:
			waypoint.x += dist
		case WEST:
			waypoint.x -= dist
		case FORWARD:
			ship.x += waypoint.x * dist
			ship.y += waypoint.y * dist
		}
	}
	fmt.Println("Part B:", math.Abs(float64(ship.x))+math.Abs(float64(ship.y)))

}
