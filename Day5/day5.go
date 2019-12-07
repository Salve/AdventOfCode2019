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

}

func runIntcode(program []int, input int) (output []int) {
	loc := 0
	for {
		instruction := decodeInstruction(program[loc])
		var params [4]int
		switch instruction[0] {
		case ADD:
			for i := 1; i < 4; i++ {
				switch instruction[i] {
				case IMM:
					params[i] = loc + i
				case POS:
					params[i] = program[loc+i]
				default:
					log.Fatalf("Undefined mode: %d", instruction[i])
				}
			}
			program[params[3]] = program[params[1]] + program[params[2]]
			loc += 4

		case MULT:
			for i := 1; i < 4; i++ {
				switch instruction[i] {
				case IMM:
					params[i] = loc + i
				case POS:
					params[i] = program[loc+i]
				default:
					log.Fatalf("Undefined mode: %d", instruction[i])
				}
			}
			program[params[3]] = program[params[1]] * program[params[2]]
			loc += 4

		case INPUT:
			switch instruction[1] {
			case IMM:
				log.Fatalf("Invalid mode IMMEDIATE for INPUT operation at location: %d", loc)
			case POS:
				program[program[loc+1]] = input
			default:
				log.Fatalf("Undefined mode: %d", instruction[1])
			}
			loc += 2

		case OUTPUT:
			switch instruction[1] {
			case IMM:
				output = append(output, program[loc+1])
			case POS:
				output = append(output, program[program[loc+1]])
			default:
				log.Fatalf("Undefined mode: %d", instruction[1])
			}
			loc += 2

		case HALT:
			loc += 1
			return

		default:
			log.Fatalf("Unexpected opcode: %v", program[loc])

		}
	}
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
	for _, value := range strings.Split(string(fc), ",") {
		intvalue, _ := strconv.Atoi(value)
		program = append(program, intvalue)
	}
	return program
}
