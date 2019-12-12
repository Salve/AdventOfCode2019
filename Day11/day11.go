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
	RBO      int = 9
	HALT     int = 99
)

const ( // parameter modes
	POS = 0
	IMM = 1
	REL = 2
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

const (
	LEFT  = 0
	RIGHT = 1
	BLACK = 0
	WHITE = 1
)

type program struct {
	intcode []int
	loc     int
	rbase   int
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

type location struct {
	x, y int
}

func (l *location) add(a location) location {
	return location{x: l.x + a.x, y: l.y + a.y}
}

type robot struct {
	location
	facing   location
	paintlog map[location]int
}

func (r *robot) paint(input int) {
	r.paintlog[r.location] = input
}

func (r *robot) turn(input int) {
	up := location{0, 1}
	right := location{1, 0}
	down := location{0, -1}
	left := location{-1, 0}
	switch r.facing {
	case up:
		switch input {
		case LEFT:
			r.facing = left
		case RIGHT:
			r.facing = right
		}
	case right:
		switch input {
		case LEFT:
			r.facing = up
		case RIGHT:
			r.facing = down
		}
	case down:
		switch input {
		case LEFT:
			r.facing = right
		case RIGHT:
			r.facing = left
		}
	case left:
		switch input {
		case LEFT:
			r.facing = down
		case RIGHT:
			r.facing = up
		}
	}
	r.location = r.location.add(r.facing)
}

func (r *robot) limits() (xMin, xMax, yMin, yMax int) {
	i := 0
	for loc, _ := range r.paintlog {
		if loc.x < xMin || i == 0 {
			xMin = loc.x
		}
		if loc.x > xMax || i == 0 {
			xMax = loc.x
		}
		if loc.y < yMin || i == 0 {
			yMin = loc.y
		}
		if loc.y > yMax || i == 0 {
			yMax = loc.y
		}
		i++
	}
	return
}

func (r *robot) visualPaintlog() {
	xMin, xMax, yMin, yMax := r.limits()
	for y := yMax; y >= yMin; y-- {
		for x := xMin; x <= xMax; x++ {
			if r.paintlog[location{x, y}] == WHITE {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	inputprogram := program{
		intcode: readIntcode("input"),
		loc:     0,
		state:   INIT,
	}

	fmt.Println("-- Part 1:")
	r := robot{
		location: location{0, 0},
		facing:   location{0, 1},
		paintlog: make(map[location]int),
	}
	inputprogram.copy().runPaintingRobot(&r)
	fmt.Printf("Unique painted locations: %d\n\n", len(r.paintlog))

	fmt.Println("-- Part 2:")
	r2 := robot{
		location: location{0, 0},
		facing:   location{0, 1},
		paintlog: make(map[location]int),
	}
	p2 := inputprogram.copy()
	p2.input = []int{1} // start on white panel
	p2.runPaintingRobot(&r2)
	r2.visualPaintlog()
}

func (p *program) runPaintingRobot(r *robot) {
	painted := false
	for {
		if p.state == HALTED {
			break
		} else if p.state == WAITOUT {
			if !painted {
				r.paint(p.output)
				painted = true
			} else {
				r.turn(p.output)
				painted = false
			}
		} else if p.state == WAITIN {
			p.input = []int{r.paintlog[r.location]}
		}
		p.run()
	}

}

func (p *program) run() {
	p.state = RUNNING
	for {
		instruction := decodeInstruction(p.intcode[p.loc])
		switch instruction[0] {
		case ADD:
			params := getParams(p, instruction, 4)
			p.intcode[params[3]] = p.intcode[params[1]] + p.intcode[params[2]]
			p.loc += 4

		case MULT:
			params := getParams(p, instruction, 4)
			p.intcode[params[3]] = p.intcode[params[1]] * p.intcode[params[2]]
			p.loc += 4

		case INPUT:
			params := getParams(p, instruction, 2)
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
			params := getParams(p, instruction, 2)
			p.output = p.intcode[params[1]]
			p.state = WAITOUT
			p.loc += 2
			return

		case JMPTRUE:
			params := getParams(p, instruction, 3)
			if p.intcode[params[1]] != 0 {
				p.loc = p.intcode[params[2]]
			} else {
				p.loc += 3
			}

		case JMPFALSE:
			params := getParams(p, instruction, 3)
			if p.intcode[params[1]] == 0 {
				p.loc = p.intcode[params[2]]
			} else {
				p.loc += 3
			}

		case LESS:
			params := getParams(p, instruction, 4)
			if p.intcode[params[1]] < p.intcode[params[2]] {
				p.intcode[params[3]] = 1
			} else {
				p.intcode[params[3]] = 0
			}
			p.loc += 4

		case EQL:
			params := getParams(p, instruction, 4)
			if p.intcode[params[1]] == p.intcode[params[2]] {
				p.intcode[params[3]] = 1
			} else {
				p.intcode[params[3]] = 0
			}
			p.loc += 4

		case RBO:
			params := getParams(p, instruction, 2)
			p.rbase += p.intcode[params[1]]
			p.loc += 2

		case HALT:
			p.state = HALTED
			return

		default:
			log.Fatalf("Unexpected opcode: %v", p.intcode[p.loc])
		}
	}
}

func getParams(p *program, instruction [4]int, num int) (params [4]int) {
	p.checkLen(p.loc + num)
	for i := 1; i < num; i++ {
		switch instruction[i] {
		case IMM:
			params[i] = p.loc + i
		case POS:
			p.checkLen(p.intcode[p.loc+i])
			params[i] = p.intcode[p.loc+i]
		case REL:
			p.checkLen(p.intcode[p.loc+i] + p.rbase)
			params[i] = p.intcode[p.loc+i] + p.rbase
		default:
			log.Fatalf("Undefined mode: %d", instruction[i])
		}
	}
	return params
}

func (p *program) checkLen(loc int) {
	if loc >= len(p.intcode) {
		p.intcode = append(p.intcode, make([]int, loc-len(p.intcode)+1)...)
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
	for _, value := range strings.Split(strings.TrimSpace(string(fc)), ",") {
		intvalue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Invalid integer in intcode: %v", err)
		}
		program = append(program, intvalue)
	}
	return program
}
