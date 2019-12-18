package main

import (
	intcode "AdventOfCode2019/Intcode"
	"fmt"
	"log"
)

// Tiles
const (
	EMPTY   = 0
	WALL    = 1
	BLOCK   = 2
	HPADDLE = 3
	BALL    = 4
)

// Joystick inputs
const (
	RIGHT = 1
	LEFT  = -1
	IDLE  = 0
)

type loc struct {
	x, y int
}

func limits(m map[loc]int) (xMin, xMax, yMin, yMax int) {
	i := 0
	for loc := range m {
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

func display(m map[loc]int, score int) {
	fmt.Printf("Score: %d\n", score)
	xMin, xMax, yMin, yMax := limits(m)
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			switch m[loc{x, y}] {
			case EMPTY:
				fmt.Print("  ")
			case WALL:
				fmt.Print("██")
			case BLOCK:
				fmt.Print("▒▒")
			case HPADDLE:
				fmt.Print("__")
			case BALL:
				fmt.Print("<>")
			}
		}
		fmt.Print("\n")
	}
}

func ballLoc(m map[loc]int) loc {
	for k, v := range m {
		if v == BALL {
			return k
		}
	}
	return loc{0, 0} // can't find the ball - not output yet?
}
func paddleLoc(m map[loc]int) loc {
	for k, v := range m {
		if v == HPADDLE {
			return k
		}
	}
	return loc{0, 0} // can't find the paddle - not output yet?
}
func move(m map[loc]int) int {
	ball := ballLoc(m)
	paddle := paddleLoc(m)
	if ball.x > paddle.x {
		return RIGHT
	}
	if ball.x < paddle.x {
		return LEFT
	}
	return IDLE
}

func game(p []int) (map[loc]int, int) {
	in := make(chan int)
	out := make(chan int)
	halt := make(chan bool)
	reqin := make(chan bool)

	score := 0
	grid := make(map[loc]int)
	b := []int{}

	go intcode.Run(p, in, out, halt, reqin)

	for {
		select {
		case <-halt:
			return grid, score
		case o := <-out:
			b = append(b, o)
			if len(b) == 3 {
				if b[0] == -1 && b[1] == 0 {
					score = b[2]
				} else {
					grid[loc{b[0], b[1]}] = b[2]
				}
				b = []int{}
			}
		case <-reqin:
			// display(grid, score)
			// time.Sleep(time.Millisecond * 50)
			in <- move(grid)
		}
	}
}

func main() {
	inputprogram, err := intcode.FromFile("input")
	if err != nil {
		log.Fatalf("Failed to read intcode from file: %s\n", err)
	}

	fmt.Println("-- Part 1:")
	p1 := intcode.Copy(inputprogram)
	grid, _ := game(p1)
	c := 0
	for _, v := range grid {
		if v == BLOCK {
			c++
		}
	}
	fmt.Printf("Blocks on starting screen: %d\n\n", c)

	fmt.Println("-- Part 2:")
	p2 := intcode.Copy(inputprogram)
	p2[0] = 2 // set freeplay
	_, finalScore := game(p2)
	fmt.Printf("Final score: %d\n", finalScore)

}
