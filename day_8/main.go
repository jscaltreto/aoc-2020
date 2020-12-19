package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
)

const (
	NOP = "nop"
	ACC = "acc"
	JMP = "jmp"
)

type Instruction struct {
	instruction string
	value       int
	execCount   uint
}

type Instructions []*Instruction

func main() {
	instructions := Instructions{}
	loadInput(&instructions)
	b(instructions)
}

func loadInput(instructions *Instructions) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	instRe, _ := regexp.Compile(`^(\w{3}) ([+-]\d+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		match := instRe.FindAllStringSubmatch(line, -1)
		if match == nil {
			log.Fatalf("Invalid Instruction: %s", line)
		}
		inst := match[0][1]
		value, _ := strconv.Atoi(match[0][2])
		*instructions = append(*instructions, &Instruction{
			instruction: inst,
			value:       value,
			execCount:   0,
		})
	}
}

func b(instructions Instructions) {
	cursor := 0
	accum := 0

	var seen []int
	for cursor < len(instructions) {
		inst := instructions[cursor]
		if inst.execCount == 1 {
			break
		}
		switch inst.instruction {
		case NOP:
			seen = append(seen, cursor)
			cursor++
		case JMP:
			seen = append(seen, cursor)
			cursor += inst.value
		case ACC:
			cursor++
			accum += inst.value
		}
		inst.execCount++
	}
	fmt.Printf("Part A Accumulator: %d\n", accum)

	// We'll use goroutines to parallelize the search.
	// Once we find the correct instruction to flip we'll
	// cancel all of the others (which are in infinite loops!).
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	// We know that the instruction to be flipped MUST be one
	// that we already hit in Part A, so we'll only check those.
	for _, inst := range seen {
		wg.Add(1)
		go func(i int) {
			runWithFlip(ctx, instructions, i)
			cancel()
			wg.Done()
		}(inst)
	}
	wg.Wait()
}

func runWithFlip(ctx context.Context, instructions Instructions, flipInst int) {
	cursor := 0
	accum := 0
	for cursor < len(instructions) {
		select {
		case <-ctx.Done():
			return
		default:
			inst := instructions[cursor]
			instruction := inst.instruction
			if cursor == flipInst {
				if instruction == NOP {
					instruction = JMP
				} else {
					instruction = NOP
				}
			}
			switch instruction {
			case NOP:
				cursor++
			case JMP:
				cursor += inst.value
			case ACC:
				cursor++
				accum += inst.value
			}
		}
	}
	fmt.Printf("Part B Accumulator: %d (P.S. Instruction %d was flipped!)\n", accum, flipInst)
}
