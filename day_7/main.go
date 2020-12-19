package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	color    string
	parents  map[*Bag]int
	children map[*Bag]int
}

func newBag(color string) *Bag {
	bag := new(Bag)
	bag.color = color
	bag.parents = make(map[*Bag]int)
	bag.children = make(map[*Bag]int)
	return bag
}

type Bags map[string]*Bag

func (b *Bags) getBag(color string) *Bag {
	if bag, ok := (*b)[color]; ok {
		return bag
	}
	bag := newBag(color)
	(*b)[color] = bag
	return bag
}

func main() {
	bags := make(Bags)
	loadInput(&bags)
	a(&bags)
	b(&bags)
}

func loadInput(bags *Bags) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	ruleRe, _ := regexp.Compile(`^(\w+ \w+) bags contain (.*)\.$`)
	containsRe, _ := regexp.Compile(`^(\d+) (\w+ \w+) bags?$`)
	for scanner.Scan() {
		line := scanner.Text()
		match := ruleRe.FindAllStringSubmatch(line, -1)
		if match == nil {
			log.Fatalf("Invalid Rule: %s", line)
		}
		parentColor := match[0][1]
		parent := bags.getBag(parentColor)
		for _, contents := range strings.Split(match[0][2], ", ") {
			childMatch := containsRe.FindAllStringSubmatch(contents, -1)
			if childMatch == nil {
				continue
			}
			numContains, _ := strconv.Atoi(childMatch[0][1])
			childBag := bags.getBag(childMatch[0][2])
			childBag.parents[parent] = numContains
			parent.children[childBag] = numContains
		}
	}
}

func a(bags *Bags) {
	myBag := bags.getBag("shiny gold")

	var checkBags []*Bag
	for b := range myBag.parents {
		checkBags = append(checkBags, b)
	}

	seenBags := make(map[*Bag]bool)
	for len(checkBags) > 0 {
		myBag, checkBags = checkBags[0], checkBags[1:]
		if _, ok := seenBags[myBag]; ok {
			continue
		}
		for b := range myBag.parents {
			checkBags = append(checkBags, b)
		}
		seenBags[myBag] = true
	}
	log.Printf("Outer Bags: %d", len(seenBags))
}

func b(bags *Bags) {
	checkBags := make([]*Bag, 1)
	checkBags[0] = bags.getBag("shiny gold")

	totalBags := 0
	for len(checkBags) > 0 {
		myBag := checkBags[0]
		checkBags = checkBags[1:]
		for b, num := range myBag.children {
			for n := 0; n < num; n++ {
				checkBags = append(checkBags, b)
				totalBags++
			}
		}
	}
	log.Printf("Total Bags: %d", totalBags)
}
