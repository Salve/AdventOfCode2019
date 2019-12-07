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
	part1program := make([]int, len(inputprogram))
	copy(part1program, inputprogram)
	for _, output := range runIntcode(part1program, 1) {
		fmt.Println(output)
	}

	fmt.Println("\n-- Part 2:")
	part2program := make([]int, len(inputprogram))
	copy(part2program, inputprogram)
	for _, output := range runIntcode(part2program, 5) {
		fmt.Println(output)
	}

}

func runIntcode(program []int, input int) (output []int) {
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
			program[params[1]] = input
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
