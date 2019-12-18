package intcode

import (
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
	RBO      int = 9
	HALT     int = 99
)

const ( // parameter modes
	POS = 0
	IMM = 1
	REL = 2
)

func Run(p []int, in <-chan int, out chan<- int, halted chan<- bool, reqin chan<- bool) {
	loc := 0
	rbase := 0

	for {
		instruction := decodeInstruction(p[loc])
		switch instruction[0] {
		case ADD:
			params := getParams(&p, loc, rbase, instruction, 4)
			p[params[3]] = p[params[1]] + p[params[2]]
			loc += 4

		case MULT:
			params := getParams(&p, loc, rbase, instruction, 4)
			p[params[3]] = p[params[1]] * p[params[2]]
			loc += 4

		case INPUT:
			params := getParams(&p, loc, rbase, instruction, 2)
			if cap(reqin) == 0 {
				// If passed an unbuffered reqin channel, use it to sync with driver code and request input
				reqin <- true
			}
			p[params[1]] = <-in
			loc += 2

		case OUTPUT:
			params := getParams(&p, loc, rbase, instruction, 2)
			out <- p[params[1]]
			loc += 2

		case JMPTRUE:
			params := getParams(&p, loc, rbase, instruction, 3)
			if p[params[1]] != 0 {
				loc = p[params[2]]
			} else {
				loc += 3
			}

		case JMPFALSE:
			params := getParams(&p, loc, rbase, instruction, 3)
			if p[params[1]] == 0 {
				loc = p[params[2]]
			} else {
				loc += 3
			}

		case LESS:
			params := getParams(&p, loc, rbase, instruction, 4)
			if p[params[1]] < p[params[2]] {
				p[params[3]] = 1
			} else {
				p[params[3]] = 0
			}
			loc += 4

		case EQL:
			params := getParams(&p, loc, rbase, instruction, 4)
			if p[params[1]] == p[params[2]] {
				p[params[3]] = 1
			} else {
				p[params[3]] = 0
			}
			loc += 4

		case RBO:
			params := getParams(&p, loc, rbase, instruction, 2)
			rbase += p[params[1]]
			loc += 2

		case HALT:
			halted <- true
			return

		default:
			log.Fatalf("Unexpected opcode: %v", p[loc])
		}
	}
}

func getParams(p *[]int, loc int, rbase int, instruction [4]int, num int) (params [4]int) {
	checkLen(p, loc+num)
	for i := 1; i < num; i++ {
		switch instruction[i] {
		case IMM:
			params[i] = loc + i
		case POS:
			checkLen(p, (*p)[loc+i])
			params[i] = (*p)[loc+i]
		case REL:
			checkLen(p, (*p)[loc+i]+rbase)
			params[i] = (*p)[loc+i] + rbase
		default:
			log.Fatalf("Undefined mode: %d", instruction[i])
		}
	}
	return params
}

func checkLen(p *[]int, loc int) {
	if loc >= len(*p) {
		*p = append(*p, make([]int, loc-len(*p)+1)...)
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

func FromFile(filename string) ([]int, error) {
	var program []int

	f, err := os.Open(filename)
	if err != nil {
		return program, err
	}
	defer f.Close()

	fc, err := ioutil.ReadFile(filename)
	if err != nil {
		return program, err
	}

	for _, value := range strings.Split(strings.TrimSpace(string(fc)), ",") {
		intvalue, err := strconv.Atoi(value)
		if err != nil {
			return program, err
		}
		program = append(program, intvalue)
	}
	return program, err
}

func Copy(p []int) []int {
	t := append([]int{}, p...)
	return t
}
