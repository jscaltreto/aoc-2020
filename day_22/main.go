package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	INPUT = "input.txt"
)

var decks [][]int

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
	for _, player := range strings.Split(string(data), "\n\n") {
		deck := []int{}
		for _, line := range strings.Split(player, "\n")[1:] {
			card, _ := strconv.Atoi(line)
			deck = append(deck, card)
		}
		decks = append(decks, deck)
	}
}

func cloneDecks(srcDecks [][]int) [][]int {
	gameDecks := make([][]int, 2)
	for player, deck := range srcDecks {
		newDeck := make([]int, len(deck))
		copy(newDeck, deck)
		gameDecks[player] = newDeck
	}
	return gameDecks
}

func a() {
	winner := 1
	gameDecks := cloneDecks(decks)
GAME:
	for {
		playerCards := make([]int, 2)
		for player, deck := range gameDecks {
			playerCards[player] = deck[0]
			gameDecks[player] = deck[1:]
		}
		if playerCards[0] > playerCards[1] {
			gameDecks[0] = append(gameDecks[0], playerCards...)
		} else {
			gameDecks[1] = append(gameDecks[1], playerCards[1], playerCards[0])
		}
		for player, deck := range gameDecks {
			if len(deck) == 0 {
				winner ^= player
				break GAME
			}
		}
	}
	score := 0
	for id, card := range gameDecks[winner] {
		score += card * (len(gameDecks[winner]) - id)
	}
	fmt.Println("Part A:", score)
}

func recursiveCombat(srcDecks [][]int) (int, [][]int) {
	memory := make([]map[string]bool, 2)
	memory[0], memory[1] = make(map[string]bool), make(map[string]bool)
	gameDecks := cloneDecks(srcDecks)
	for {
		playerCards := make([]int, 2)
		for player, deck := range gameDecks {
			deck_order := fmt.Sprint(deck)
			if _, ok := memory[player][deck_order]; ok {
				return 0, gameDecks
			}
			memory[player][deck_order] = true
			playerCards[player] = deck[0]
			gameDecks[player] = deck[1:]
		}
		if len(gameDecks[0]) >= playerCards[0] &&
			len(gameDecks[1]) >= playerCards[1] {
			subDecks := [][]int{
				gameDecks[0][:playerCards[0]],
				gameDecks[1][:playerCards[1]],
			}
			subwinner, _ := recursiveCombat(subDecks)
			gameDecks[subwinner] = append(gameDecks[subwinner], playerCards[subwinner], playerCards[subwinner^1])
		} else {
			for player, deck := range gameDecks {
				if len(deck) < playerCards[player] {
					break
				}
			}
			if playerCards[0] > playerCards[1] {
				gameDecks[0] = append(gameDecks[0], playerCards...)
			} else {
				gameDecks[1] = append(gameDecks[1], playerCards[1], playerCards[0])
			}
		}
		for player, deck := range gameDecks {
			if len(deck) == 0 {
				return player^1, gameDecks
			}
		}
	}
}

func b() {
	winner, gameDecks := recursiveCombat(decks)
	score := 0
	for id, card := range gameDecks[winner] {
		score += card * (len(gameDecks[winner]) - id)
	}
	fmt.Println("Part B:", score)
}
