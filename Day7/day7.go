package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const ( // intcode commands
	ADD      int = 1
	MULT     int = 2
	INPUT    int = 3
	OUTPUT   int = 4
	JMPTRUE  int = 5
	JMPFALSE int = 6
	LESS     int = 7
	EQL      int = 8
	HALT     int = 99
)

const ( // parameter modes
	POS = 0
	IMM = 1
)

type state int

const (
	INIT state = iota
	RUNNING
	WAITOUT
	WAITIN
	PAUSED
	HALTED
)

type program struct {
	intcode []int
	loc     int
	state   state
	input   []int
	output  int
}

func (p *program) copy() *program {
	i2 := make([]int, len(p.intcode))
	copy(i2, p.intcode)
	p2 := program{
		intcode: i2,
		loc:     p.loc,
		state:   p.state,
		input:   p.input,
		output:  p.output,
	}
	return &p2
}

func main() {
	inputprogram := program{
		intcode: readIntcode("input"),
		loc:     0,
		state:   INIT,
	}

	fmt.Println("-- Part 1:")
	max := 0
	permute([]int{0, 1, 2, 3, 4}, func(p []int) {
		if out := tryPhaseSequence(p, &inputprogram); out > max {
			max = out
		}
	})
	fmt.Printf("Max thruster signal achieved: %d\n\n", max)

	fmt.Println("-- Part 2:")
	max = 0
	permute([]int{5, 6, 7, 8, 9}, func(p []int) {
		if out := tryFeedbackSequence(p, &inputprogram); out > max {
			max = out
		}
	})
	fmt.Printf("Max thruster signal achieved: %d\n\n", max)

}

func tryPhaseSequence(seq []int, prg *program) int {
	input := 0
	for i := range seq {
		p := prg.copy()
		p.input = []int{seq[i], input}
		for {
			if p.state == HALTED {
				break
			} else if p.state == WAITOUT {
				input = p.output
			}
			p.run()
		}
	}
	return input
}

func tryFeedbackSequence(seq []int, prg *program) int {
	amps := [5]*program{}
	for i := 0; i < 5; i++ {
		amps[i] = prg.copy()
		amps[i].input = []int{seq[i]} // set phase as first input
	}

	s := []int{0} // current signal value stack
	i := 0        // currently running amp
	for {
		if i == 4 && amps[i].state == HALTED {
			return s[0]
		}
		if amps[i].state == WAITIN {
			amps[i].input = append(amps[i].input, s[0])
			s = s[1:]
		}
		if amps[i].state == WAITOUT {
			s = append([]int{amps[i].output}, s...) // push output to s
			// run next amp
			amps[i].state = PAUSED
			if i == 4 {
				i = 0
			} else {
				i++
			}
			continue
		}
		amps[i].run()
		if amps[i].state == HALTED {
			if i == 4 {
				i = 0
			} else {
				i++
			}
			continue
		}
	}
}

func (p *program) run() {
	p.state = RUNNING
	for {
		instruction := decodeInstruction(p.intcode[p.loc])
		switch instruction[0] {
		case ADD:
			params := getParams(p.intcode, instruction, p.loc, 4)
			p.intcode[params[3]] = p.intcode[params[1]] + p.intcode[params[2]]
			p.loc += 4

		case MULT:
			params := getParams(p.intcode, instruction, p.loc, 4)
			p.intcode[params[3]] = p.intcode[params[1]] * p.intcode[params[2]]
			p.loc += 4

		case INPUT:
			params := getParams(p.intcode, instruction, p.loc, 2)
			if len(p.input) == 0 {
				p.state = WAITIN
				return
			}
			p.intcode[params[1]] = p.input[0]
			if len(p.input) > 0 {
				p.input = p.input[1:]
			}
			p.loc += 2

		case OUTPUT:
			params := getParams(p.intcode, instruction, p.loc, 2)
			p.output = p.intcode[params[1]]
			p.state = WAITOUT
			p.loc += 2
			return

		case JMPTRUE:
			params := getParams(p.intcode, instruction, p.loc, 3)
			if p.intcode[params[1]] != 0 {
				p.loc = p.intcode[params[2]]
			} else {
				p.loc += 3
			}

		case JMPFALSE:
			params := getParams(p.intcode, instruction, p.loc, 3)
			if p.intcode[params[1]] == 0 {
				p.loc = p.intcode[params[2]]
			} else {
				p.loc += 3
			}

		case LESS:
			params := getParams(p.intcode, instruction, p.loc, 4)
			if p.intcode[params[1]] < p.intcode[params[2]] {
				p.intcode[params[3]] = 1
			} else {
				p.intcode[params[3]] = 0
			}
			p.loc += 4

		case EQL:
			params := getParams(p.intcode, instruction, p.loc, 4)
			if p.intcode[params[1]] == p.intcode[params[2]] {
				p.intcode[params[3]] = 1
			} else {
				p.intcode[params[3]] = 0
			}
			p.loc += 4

		case HALT:
			p.state = HALTED
			return

		default:
			log.Fatalf("Unexpected opcode: %v", p.intcode[p.loc])
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
