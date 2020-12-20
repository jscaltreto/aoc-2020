package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	INPUT       = "input.txt"
	SEA_MONSTER = "                  # \n" +
		"#    ##    ##    ###\n" +
		" #  #  #  #  #  #   "
)

type TileRow []rune
type TileData []TileRow

type Coord struct {
	x, y int
}

type Tile struct {
	id    int
	data  TileData
	edges [4]*Tile
	x     int
	y     int
}

func (t *Tile) getEdges() []string {
	n := string(t.data[0][:])
	s := string(t.data[len(t.data)-1][:])
	var e, w string
	for _, r := range t.data {
		w += string(r[0])
		e += string(r[len(r)-1])
	}
	return []string{n, e, s, w}
}

func (t *Tile) rot() {
	newData := make(TileData, len(t.data))
	for rid, row := range t.data {
		newData[rid] = make(TileRow, len(row))
	}
	for rid, row := range t.data {
		for cid, char := range row {
			newData[cid][len(row)-1-rid] = char
		}
	}
	t.data = newData
}

func (t *Tile) flipx() {
	newData := make(TileData, len(t.data))
	for rid, row := range t.data {
		newData[rid] = make(TileRow, len(row))
		for cid, char := range row {
			newData[rid][len(row)-1-cid] = char
		}
	}
	t.data = newData
}

func (t *Tile) flipy() {
	newData := make(TileData, len(t.data))
	for rid, row := range t.data {
		newData[len(t.data)-1-rid] = row
	}
	t.data = newData
}

func (t *Tile) checkMatch(edge string, edgeIdx int) bool {
	if t.edges[edgeIdx] != nil {
		return false
	}
	for r := 0; r <= 3; r++ {
		checkEdge := t.getEdges()[edgeIdx]
		if checkEdge == edge {
			return true
		}
		if (edgeIdx % 2) == 0 {
			t.flipx()
		} else {
			t.flipy()
		}
		checkEdge = t.getEdges()[edgeIdx]
		if checkEdge == edge {
			return true
		}
		if (edgeIdx % 2) == 0 {
			t.flipx()
		} else {
			t.flipy()
		}
		t.rot()
	}
	return false
}

func (t *Tile) toString() string {
	rows := []string{}
	for _, row := range t.data {
		rows = append(rows, string(row[:]))
	}
	return strings.Join(rows, "\n")
}

func (t *Tile) countHashes() int {
	hashCount := 0
	for _, row := range t.data {
		for _, char := range row {
			if char == '#' {
				hashCount++
			}
		}
	}
	return hashCount
}
func (t *Tile) monsterCheck() int {
	monsterCoords := make(map[Coord]bool)
	for l, line := range t.data[monsterHeight-1:] {
		lid := l + monsterHeight - 1
	CHECKLINE:
		for cid := range line[:len(line)-monsterWidth] {
			checkCoords := []Coord{}
			for _, trans := range monsterMap {
				coord := Coord{cid + trans.x, lid + trans.y}
				if t.data[coord.y][coord.x] != '#' {
					continue CHECKLINE
				}
				checkCoords = append(checkCoords, coord)
			}
			for _, coord := range checkCoords {
				monsterCoords[coord] = true
			}
		}
	}
	return len(monsterCoords)
}

func newTile(data string) *Tile {
	tileData := strings.Split(data, "\n")
	t := new(Tile)
	id, _ := strconv.Atoi(tileData[0][5:9])
	t.id = id
	t.data = make(TileData, len(tileData)-1)
	for rid, row := range tileData[1:] {
		t.data[rid] = make(TileRow, len(row))
		for cid, char := range row {
			t.data[rid][cid] = char
		}
	}
	t.getEdges()
	return t
}

var tiles []*Tile
var min, max Coord
var tileMap map[Coord]*Tile
var monsterHeight, monsterWidth int
var monsterMap []Coord

func main() {
	min, max = Coord{0, 0}, Coord{0, 0}
	tileMap = make(map[Coord]*Tile)
	loadInput()
	a()
	b()
}

func loadInput() {
	data, err := ioutil.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	rawTiles := strings.Split(string(data), "\n\n")
	for _, t := range rawTiles {
		tiles = append(tiles, newTile(t))
	}
}

func a() {
	matched := []*Tile{tiles[0]}
	tileMap[min] = tiles[0]

	for len(tileMap) < len(tiles) {
	CHECK:
		for _, tile := range matched {
			for eid, edge := range tile.getEdges() {
				if tile.edges[eid] != nil {
					continue
				}
				checkEdge := (eid + 2) % 4
				for _, checkTile := range tiles {
					if checkTile.id == tile.id {
						continue
					}
					if checkTile.checkMatch(edge, checkEdge) {
						switch eid {
						case 0:
							checkTile.x = tile.x
							checkTile.y = tile.y - 1
						case 1:
							checkTile.x = tile.x + 1
							checkTile.y = tile.y
						case 2:
							checkTile.x = tile.x
							checkTile.y = tile.y + 1
						case 3:
							checkTile.x = tile.x - 1
							checkTile.y = tile.y
						}
						if checkTile.x < min.x {
							min.x = checkTile.x
						} else if checkTile.x > max.x {
							max.x = checkTile.x
						}
						if checkTile.y < min.y {
							min.y = checkTile.y
						} else if checkTile.y > max.y {
							max.y = checkTile.y
						}
						tileMap[Coord{checkTile.x, checkTile.y}] = checkTile
						tile.edges[eid] = checkTile
						checkTile.edges[checkEdge] = tile
						matched = append(matched, checkTile)
						break CHECK
					}
				}
			}
		}

	}
	product := tileMap[Coord{min.x, min.y}].id *
		tileMap[Coord{max.x, min.y}].id *
		tileMap[Coord{min.x, max.y}].id *
		tileMap[Coord{max.x, max.y}].id

	fmt.Println("Part A:", product)
}

func processSeaMonster() {
	seaMonster := strings.Split(SEA_MONSTER, "\n")
	monsterHeight = len(seaMonster)
	monsterMap = make([]Coord, 0)
	for lid, line := range seaMonster {
		monsterWidth = len(line)
		for cid, c := range line {
			if c == '#' {
				monsterMap = append(monsterMap, Coord{cid, 1 - monsterHeight + lid})
			}
		}
	}
}

func b() {
	imageData := []string{}
	for row := min.y; row <= max.y; row++ {
		lines := make([]string, 8)
		for col := min.x; col <= max.x; col++ {
			tile := tileMap[Coord{col, row}]
			for lid, line := range tile.data[1:9] {
				lines[lid] += string(line[1:9])
			}
		}
		for _, line := range lines {
			imageData = append(imageData, line)
		}
	}
	image := newTile("Imag 0000\n" + strings.Join(imageData, "\n"))

	processSeaMonster()

	monsterCoords := 0
	for r := 0; r <= 3; r++ {
		monsterCoords = image.monsterCheck()
		if monsterCoords > 0 {
			break
		}
		image.flipx()
		monsterCoords = image.monsterCheck()
		if monsterCoords > 0 {
			break
		}
		image.flipx()
		image.rot()
	}

	fmt.Println("Part B:", image.countHashes()-monsterCoords)
}
