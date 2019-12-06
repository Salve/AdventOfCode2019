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
	fmt.Printf("Closest intersect: %d", findClosestIntersect(layouts, startingLoc))
}

func findClosestIntersect(wirelayouts [][]wireSegment, startingLoc loc) (minDistance int) {
	// For each set of layouts/paths, generate a slice of locations visited
	// use a map of locations to indicate how many wire layouts have visited each location
	locVisits := make(map[loc]int)
	for _, layout := range wirelayouts {
		for loc := range locsForLayout(layout, startingLoc) {
			locVisits[loc]++
		}
	}

	// for each location where wire layouts intersect, calculate distance to starting point
	var intersects []int
	for loc, v := range locVisits {
		if v > 1 {
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

func manhattanDistance(locA loc, locB loc) int {
	return abs(locA.x-locB.x) + abs(locA.y-locB.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func locsForLayout(wirelayout []wireSegment, startingLoc loc) map[loc]struct{} {
	// Returns a set of locations visited by a given set of wire paths.
	// Using a map to avoid returning a location twice, as a wire's intersection with itself should not count

	locs := make(map[loc]struct{})
	currentLoc := startingLoc
	var movement loc
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
			currentLoc.x += movement.x
			currentLoc.y += movement.y
			locs[currentLoc] = struct{}{}
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
