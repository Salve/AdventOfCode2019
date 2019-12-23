package main

import (
	intcode "AdventOfCode2019/Intcode"
	"fmt"
	"log"
)

var mov = map[int]loc{
	1: loc{0, -1}, // North
	2: loc{0, 1},  // South
	3: loc{-1, 0}, // West
	4: loc{1, 0},  // East
}
var inverse = map[int]int{1: 2, 2: 1, 3: 4, 4: 3}

type loc struct {
	x, y int
}

func (l *loc) Add(l2 loc) loc {
	return loc{
		x: l.x + l2.x,
		y: l.y + l2.y,
	}
}

func printMaze(m map[loc]int, pos loc) {
	xMin, xMax, yMin, yMax := limits(m)
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			cl := loc{x, y}
			if cl == pos {
				fmt.Print("XXX")
				continue
			}
			if l, ok := m[cl]; ok {
				switch l {
				case 0:
					fmt.Print("███")
				case 1:
					fmt.Print("   ")
				case 2:
					fmt.Print("OOO")
				}
			} else {
				fmt.Print("???")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
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

func runDrone(p []int) (map[loc]int, loc) {
	in := make(chan int)
	out := make(chan int)
	halt := make(chan bool)
	reqin := make(chan bool, 1)

	maze := map[loc]int{loc{0, 0}: 1}
	var oxygen loc
	back := []int{}
	pos := loc{0, 0}

	go intcode.Run(p, in, out, halt, reqin)

Loop:
	for {
		//printMaze(maze, pos)
		for k, v := range mov {
			next := pos.Add(v)
			if _, exists := maze[next]; exists {
				continue
			}
			in <- k
			o := <-out
			maze[next] = o
			if o == 2 {
				oxygen = next
				fmt.Printf("Found oxygen at coords [%d, %d], %d moves from starting position.\n\n", oxygen.x, oxygen.y, len(back)+1)
			}
			if o > 0 { // did not hit a wall, moved
				back = append(back, inverse[k])
				pos = next
				continue Loop
			}
		}
		if len(back) == 0 {
			// Back at start - done exploring
			break Loop
		}
		// All moves exhausted, backtrack
		in <- back[len(back)-1]
		<-out // should never be a wall while backtracking
		pos = pos.Add(mov[back[len(back)-1]])
		back = back[:len(back)-1]
	}

	return maze, oxygen
}

func main() {
	inputprogram, err := intcode.FromFile("input")
	if err != nil {
		log.Fatalf("Failed to read intcode from file: %s\n", err)
	}

	fmt.Println("-- Part 1:")
	maze, oxygen := runDrone(intcode.Copy(inputprogram))

	fmt.Println("-- Part 2:")
	queue := []loc{oxygen}
	distance := map[loc]int{oxygen: 0}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		for _, v := range mov { // for each cardinal direction
			next := pos.Add(v)
			if _, exists := distance[next]; exists {
				continue
			}
			if maze[next] > 0 {
				distance[next] = distance[pos] + 1
				queue = append(queue, next)
			}
		}
	}

	max := 0
	for _, v := range distance {
		if v > max {
			max = v
		}
	}

	fmt.Printf("Oxygen restored in all locations after %d minutes.\n\n", max)
}
