package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const ADD int = 1
const MULT int = 2
const INPUT int = 3
const OUTPUT int = 4
const JMPTRUE int = 5
const JMPFALSE int = 6
const LESS int = 7
const EQL int = 8
const HALT int = 99

const POS = 0 // parameter mode position
const IMM = 1 // parameter mode immediate

func main() {
	inputprogram := readIntcode("input")

	fmt.Println("-- Part 1:")
	max := 0
	permute([]int{0, 1, 2, 3, 4}, func(p []int) {
		if out := tryPhaseSequence(p, inputprogram); out > max {
			max = out
		}
	})
	fmt.Printf("Max thruster signal achieved: %d\n\n", max)

}

func tryPhaseSequence(seq []int, program []int) int {
	input := 0
	for i := range seq {
		p := make([]int, len(program))
		copy(p, program)

		input = runIntcode(p, []int{seq[i], input})[0]
	}
	return input
}

func runIntcode(program []int, input []int) (output []int) {
	loc := 0
	for {
		instruction := decodeInstruction(program[loc])
		switch instruction[0] {
		case ADD:
			params := getParams(program, instruction, loc, 4)
			program[params[3]] = program[params[1]] + program[params[2]]
			loc += 4

		case MULT:
			params := getParams(program, instruction, loc, 4)
			program[params[3]] = program[params[1]] * program[params[2]]
			loc += 4

		case INPUT:
			params := getParams(program, instruction, loc, 2)
			program[params[1]] = input[0]
			if len(input) > 1 {
				input = input[1:]
			}
			loc += 2

		case OUTPUT:
			params := getParams(program, instruction, loc, 2)
			output = append(output, program[params[1]])
			loc += 2

		case JMPTRUE:
			params := getParams(program, instruction, loc, 3)
			if program[params[1]] != 0 {
				loc = program[params[2]]
			} else {
				loc += 3
			}

		case JMPFALSE:
			params := getParams(program, instruction, loc, 3)
			if program[params[1]] == 0 {
				loc = program[params[2]]
			} else {
				loc += 3
			}

		case LESS:
			params := getParams(program, instruction, loc, 4)
			if program[params[1]] < program[params[2]] {
				program[params[3]] = 1
			} else {
				program[params[3]] = 0
			}
			loc += 4

		case EQL:
			params := getParams(program, instruction, loc, 4)
			if program[params[1]] == program[params[2]] {
				program[params[3]] = 1
			} else {
				program[params[3]] = 0
			}
			loc += 4

		case HALT:
			loc += 1
			return output

		default:
			log.Printf("Unexpected opcode: %v", program[loc])
			return []int{-1}
		}
	}
}

func getParams(program []int, instruction [4]int, loc int, num int) (params [4]int) {
	for i := 1; i < num; i++ {
		switch instruction[i] {
		case IMM:
			params[i] = loc + i
		case POS:
			params[i] = program[loc+i]
		default:
			log.Fatalf("Undefined mode: %d", instruction[i])
		}
	}
	return params
}

func decodeInstruction(v int) (instruction [4]int) {
	instruction[0] = v % 100 // opcode
	v = v / 100
	for i := 1; i < 4; i++ {
		instruction[i] = v % 10 // mode
		v = v / 10              // advance
	}
	return instruction
}

func readIntcode(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	fc, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading program from file: %v", err)
	}

	var program []int
	for _, value := range strings.Split(strings.TrimSpace(string(fc)), ",") {
		intvalue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Invalid integer in intcode: %v", err)
		}
		program = append(program, intvalue)
	}
	return program
}

// https://rosettacode.org/wiki/Permutations#Go
func permute(s []int, emit func([]int)) {
	// permute function.  takes a set to permute and a function
	// to call for each generated permutation.

	if len(s) == 0 {
		emit(s)
		return
	}
	// Steinhaus, implemented with a recursive closure.
	// arg is number of positions left to permute.
	// pass in len(s) to start generation.
	// on each call, weave element at pp through the elements 0..np-2,
	// then restore array to the way it was.
	var rc func(int)
	rc = func(np int) {
		if np == 1 {
			emit(s)
			return
		}
		np1 := np - 1
		pp := len(s) - np1
		// weave
		rc(np1)
		for i := pp; i > 0; i-- {
			s[i], s[i-1] = s[i-1], s[i]
			rc(np1)
		}
		// restore
		w := s[0]
		copy(s, s[1:pp+1])
		s[pp] = w
	}
	rc(len(s))
}
