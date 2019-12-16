package main

import (
	intcode "AdventOfCode2019/intcode"
	"fmt"
	"log"
)

const (
	LEFT  = 0
	RIGHT = 1
	BLACK = 0
	WHITE = 1
)

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
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	inputprogram, err := intcode.FromFile("input")
	if err != nil {
		log.Fatalf("Failed to read intcode from file: %s\n", err)
	}

	fmt.Println("-- Part 1:")
	r := robot{
		location: location{0, 0},
		facing:   location{0, 1},
		paintlog: make(map[location]int),
	}
	p1 := intcode.Copy(inputprogram)
	r.run(p1)

	fmt.Printf("Unique painted locations: %d\n\n", len(r.paintlog))

	fmt.Println("-- Part 2:")
	r2 := robot{
		location: location{0, 0},
		facing:   location{0, 1},
		paintlog: map[location]int{location{0, 0}: 1},
	}
	p2 := intcode.Copy(inputprogram)
	r2.run(p2)
	r2.visualPaintlog()
}

func (r *robot) run(p []int) {
	in := make(chan int, 1)
	out := make(chan int, 2)
	halt := make(chan bool)

	go intcode.Run(p, in, out, halt)

	for {
		select {
		case <-halt:
			close(in)
			return
		default:
			in <- r.paintlog[r.location]
			r.paint(<-out)
			r.turn(<-out)
		}
	}

}
