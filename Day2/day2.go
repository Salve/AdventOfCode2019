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
const HALT int = 99

func main() {
	inputprogram := readIntcode("input")

	fmt.Println("-- Part 1:")
	part1program := make([]int, len(inputprogram))
	copy(part1program, inputprogram)
	part1program[1] = 12
	part1program[2] = 2
	fmt.Printf("Program output: %v\n\n", runIntcode(part1program)[0])

	fmt.Println("-- Part 2:")
	var verb, noun int
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			part2program := make([]int, len(inputprogram))
			copy(part2program, inputprogram)
			part2program[1] = noun
			part2program[2] = verb
			result := runIntcode(part2program)[0]
			if result == 19690720 {
				fmt.Printf("Answer: %v", 100*noun+verb)
			}
		}
	}
}

func runIntcode(program []int) []int {
	loc := 0
	for {
		switch program[loc] {
		case ADD:
			program[program[loc+3]] = program[program[loc+1]] + program[program[loc+2]]
			loc += 4
		case MULT:
			program[program[loc+3]] = program[program[loc+1]] * program[program[loc+2]]
			loc += 4
		case HALT:
			loc += 1
			return program
		default:
			log.Fatalf("Unexpected opcode: %v", program[loc])
		}
	}
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
