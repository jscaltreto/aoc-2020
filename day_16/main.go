package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	INPUT = "input.txt"
)

type Field struct {
	name string
	mina int
	maxa int
	minb int
	maxb int
}

func (f *Field) isValid(n int) bool {
	return (n >= f.mina && n <= f.maxa) || (n >= f.minb && n <= f.maxb)
}

var fields []Field

type Ticket []int

var myTicket Ticket
var nearbyTickets []Ticket

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
	sections := strings.Split(string(data), "\n\n")

	fieldRe, _ := regexp.Compile(`^([\w ]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	for _, f := range strings.Split(sections[0], "\n") {
		match := fieldRe.FindAllStringSubmatch(f, -1)
		if match == nil {
			log.Fatalf("Invalid Field: %s", f)
		}
		minA, _ := strconv.Atoi(match[0][2])
		maxA, _ := strconv.Atoi(match[0][3])
		minB, _ := strconv.Atoi(match[0][4])
		maxB, _ := strconv.Atoi(match[0][5])
		fields = append(fields, Field{match[0][1], minA, maxA, minB, maxB})
	}

	for _, n := range strings.Split(strings.Split(sections[1], "\n")[1], ",") {
		num, _ := strconv.Atoi(n)
		myTicket = append(myTicket, num)
	}

	for _, t := range strings.Split(sections[2], "\n")[1:] {
		var ticket Ticket
		for _, n := range strings.Split(t, ",") {
			num, _ := strconv.Atoi(n)
			ticket = append(ticket, num)
		}
		nearbyTickets = append(nearbyTickets, ticket)
	}
}

func validateFields() (int, []Ticket) {
	var validTickets []Ticket
	errorRate := 0
	for _, t := range nearbyTickets {
		valid := true
	NUM:
		for _, n := range t {
			for _, f := range fields {
				if f.isValid(n) {
					continue NUM
				}
			}
			valid = false
			errorRate += n
		}
		if valid {
			validTickets = append(validTickets, t)
		}
	}
	return errorRate, validTickets
}

func a() {
	errorRate, _ := validateFields()
	fmt.Println("Part A:", errorRate)
}

func b() {
	_, validTickets := validateFields()
	fieldMap := map[int]Field{}
	fieldsAvail := map[Field]bool{}
	for _, field := range fields {
		fieldsAvail[field] = true
	}
	for len(fieldsAvail) > 0 {
		for fieldId := range fields {
			if _, ok := fieldMap[fieldId]; ok {
				continue
			}
			possibleFields := map[Field]bool{}
			for _, t := range validTickets {
				for f := range fieldsAvail {
					if !f.isValid(t[fieldId]) {
						possibleFields[f] = false
					}
				}
			}
			if len(possibleFields) == len(fieldsAvail)-1 {
				for f := range fieldsAvail {
					if _, ok := possibleFields[f]; !ok {

						fieldMap[fieldId] = f
						delete(fieldsAvail, f)
						break
					}
				}

			}
		}
	}
	product := 1
	for fieldId, field := range fieldMap {
		if len(field.name) >= 9 && field.name[:9] == "departure" {
			product *= myTicket[fieldId]
		}
	}
	fmt.Println("Part B:", product)
}
