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
	program := readIntcode("input")
	program[1] = 12
	program[2] = 2
	fmt.Printf("Program utput: %v", runIntcode(program)[0])
}

func runIntcode(program []int) []int {
	loc := 0
	for {
		switch program[loc] {
		case ADD:
			program[program[loc+3]] = program[program[loc+1]] + program[program[loc+2]]
		case MULT:
			program[program[loc+3]] = program[program[loc+1]] * program[program[loc+2]]
		case HALT:
			return program
		default:
			log.Fatalf("Unexpected opcode: %v", program[loc])
		}

		loc += 4
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
