package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type wireSegment struct {
	direction byte
	distance  int
}
type loc struct {
	x int
	y int
}

func main() {
	layouts := readWireLayouts("input")
	startingLoc := loc{0, 0}

	fmt.Println("-- Part 1:")
	fmt.Printf("Closest intersect: %d\n\n", closestIntersect(layouts, startingLoc))
	fmt.Println("-- Part 2:")
	fmt.Printf("Shortest intersect: %d steps\n\n", shortestIntersect(layouts, startingLoc))
}

func closestIntersect(wirelayouts [][]wireSegment, startingLoc loc) (minDistance int) {
	locVisits := visitedLocations(wirelayouts, startingLoc)

	// for each location where wire layouts intersect, calculate distance to starting point
	var intersects []int
	for loc, v := range locVisits {
		if len(v) > 1 {
			intersects = append(intersects, manhattanDistance(startingLoc, loc))
		}
	}

	// find smallest distance to arrive at closest intersect
	for i, distance := range intersects {
		if i == 0 || distance < minDistance {
			minDistance = distance
		}
	}

	return minDistance
}

func shortestIntersect(wirelayouts [][]wireSegment, startingLoc loc) (minSteps int) {
	locVisits := visitedLocations(wirelayouts, startingLoc)

	// for each location where wire layouts intersect, calculate sum of steps to get there
	var intersects []int
	for _, steps := range locVisits {
		if len(steps) > 1 {
			intersects = append(intersects, addArray(steps))
		}
	}

	// find smallest sum of steps to an intersect
	for i, sumsteps := range intersects {
		if i == 0 || sumsteps < minSteps {
			minSteps = sumsteps
		}
	}

	return minSteps
}

func visitedLocations(wirelayouts [][]wireSegment, startingLoc loc) map[loc][]int {
	// Given a set of wire paths/layouts, return a map of locations containing
	// a slice with one element per wire which visited the location,
	// with the value of the slice element being the number of steps that wire needed to get to the location

	locVisits := make(map[loc][]int)
	for _, layout := range wirelayouts {
		for loc, steps := range locsForLayout(layout, startingLoc) {
			locVisits[loc] = append(locVisits[loc], steps)
		}
	}

	return locVisits
}

func manhattanDistance(locA loc, locB loc) int {
	return abs(locA.x-locB.x) + abs(locA.y-locB.y)
}

func addArray(values []int) (sum int) {
	for _, v := range values {
		sum += v
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func locsForLayout(wirelayout []wireSegment, startingLoc loc) map[loc]int {
	// Returns a set of locations visited by a given set of wire paths, and the number of steps to get there

	locs := make(map[loc]int)
	currentLoc := startingLoc
	var movement loc
	var step int
	for _, segment := range wirelayout {
		switch segment.direction {
		case "U"[0]:
			movement = loc{0, 1}
		case "R"[0]:
			movement = loc{1, 0}
		case "D"[0]:
			movement = loc{0, -1}
		case "L"[0]:
			movement = loc{-1, 0}
		default:
			log.Fatalf("Undefined direction: %v", segment.direction)
		}

		for i := 0; i < segment.distance; i++ {
			step++
			currentLoc.x += movement.x
			currentLoc.y += movement.y
			if _, exists := locs[currentLoc]; !exists {
				locs[currentLoc] = step
			}
		}
	}

	return locs
}

func readWireLayouts(filename string) (wirelayouts [][]wireSegment) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var layout []wireSegment
		for _, segmentStr := range strings.Split(sc.Text(), ",") {
			dir := segmentStr[0]
			dist, _ := strconv.Atoi(segmentStr[1:])
			layout = append(layout, wireSegment{direction: dir, distance: dist})
		}
		wirelayouts = append(wirelayouts, layout)
	}

	return wirelayouts
}
