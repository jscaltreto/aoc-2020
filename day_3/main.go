package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type treemap [][]bool

func main() {
	var tmap treemap
	loadInput(&tmap)
	a(tmap)
	b(tmap)
}

func loadInput(tmap *treemap) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var row []bool
		line := scanner.Text()
		if line == "" {
			continue
		}
		for _, c := range line {
			if string(c) == "#" {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		*tmap = append(*tmap, row)
	}
}

func check_slope(tmap treemap, slope_y int, slope_x int) int {
	var x, y, trees int = 0, 0, 0

	for true {
		x = (x + slope_x) % len(tmap[y])
		y += slope_y
		if y >= len(tmap) {
			break
		}
		if tmap[y][x] {
			trees += 1
		}
	}

	return trees
}

func a(tmap treemap) {
	trees := check_slope(tmap, 1, 3)

	fmt.Printf("Trees Encountered: %d\n", trees)
}

func b(tmap treemap) {
	var slopes = [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	var product = 1

	for _, slope := range slopes {
		product *= check_slope(tmap, slope[1], slope[0])
	}
	fmt.Printf("Product: %d\n", product)
}
