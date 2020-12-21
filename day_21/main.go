package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

const (
	INPUT = "input.txt"
)

type IngredientsList []string
type AllFoods []IngredientsList
type Allergens map[string]IngredientsList
type AllIngredients map[string]bool

var allergens Allergens
var allIngredients AllIngredients
var allFoods AllFoods

func main() {
	allergens = make(Allergens)
	allIngredients = make(AllIngredients)
	allFoods = make(AllFoods, 0)
	loadInput()
	a()
	b()
}

func loadInput() {
	data, err := ioutil.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	listRe := regexp.MustCompile(`^([\w ]+) \(contains ([\w, ]*)\)$`)
	for _, line := range strings.Split(string(data), "\n") {
		match := listRe.FindStringSubmatch(line)
		if match == nil {
			log.Fatalf("Invalid ingredients list: %s", line)
		}
		ingredients := strings.Split(match[1], " ")
		allFoods = append(allFoods, ingredients)
		for _, i := range ingredients {
			allIngredients[i] = true
		}
		for _, a := range strings.Split(match[2], ", ") {
			var ilist IngredientsList
			if cur, ok := allergens[a]; !ok {
				ilist = ingredients
			} else {
				ilist = make(IngredientsList, 0)
				for _, i := range ingredients {
					for _, ci := range cur {
						if i == ci {
							ilist = append(ilist, i)
						}
					}
				}
			}
			allergens[a] = ilist
		}
	}
}

func a() {
	noAllergenCount := 0
CHECKALL:
	for i, _ := range allIngredients {
		for _, allergens := range allergens {
			for _, ai := range allergens {
				if i == ai {
					continue CHECKALL
				}
			}
		}
		for _, food := range allFoods {
			for _, fi := range food {
				if i == fi {
					noAllergenCount++
				}
			}
		}
	}
	fmt.Println("Part A:", noAllergenCount)
}

func b() {
	knownAllergens := make(map[string]string)
	numAllergens := len(allergens)
	allergenNames := []string{}
	for a := range allergens {
		allergenNames = append(allergenNames, a)
	}
	sort.Strings(allergenNames)
	for len(knownAllergens) < numAllergens {
		for a, il := range allergens {
			newList := IngredientsList{}
		CHECKLIST:
			for _, i := range il {
				for _, ki := range knownAllergens {
					if ki == i {
						continue CHECKLIST
					}
				}
				newList = append(newList, i)
			}
			if len(newList) == 1 {
				knownAllergens[a] = newList[0]
				delete(allergens, a)
				break
			}
		}
	}
	badIngredients := []string{}
	for _, a := range allergenNames {
		badIngredients = append(badIngredients, knownAllergens[a])
	}
	fmt.Println("Part B:", strings.Join(badIngredients, ","))
}
