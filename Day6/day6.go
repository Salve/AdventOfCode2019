package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	orbiting := readOrbits("input")

	fmt.Println("-- Part 1:")
	count := countOrbits(orbiting)
	fmt.Printf("Direct + indirect orbits: %d\n\n", count)

	fmt.Println("-- Part 2:")
	fmt.Printf("Shortest distance: %d", shortestDist("YOU", "SAN", orbiting))
}

func shortestDist(a, b string, orbiting map[string]string) int {
	apath := make(map[string]int)
	system := a
	dist := 0
	for { // traverse from a to root, store all distances
		if parent, exists := orbiting[system]; exists {
			dist++
			apath[parent] = dist
			system = parent
		} else {
			break
		}
	}

	dist = 0
	system = b
	for { // traverse from b towards root, until hitting path from a to root
		parent := orbiting[system]
		dist++
		if adist, exists := apath[parent]; exists {
			return dist + adist - 2 // don't count a and b
		}
		system = parent
	}
}

func countOrbits(orbiting map[string]string) int {
	count := 0
	for system := range orbiting {
		count += 1                 // count direct orbit
		system := orbiting[system] // check parent for indirect orbits
		for {
			if parent, exists := orbiting[system]; exists {
				count += 1 // count indirect orbit
				system = parent
			} else {
				break // reached top of chain
			}
		}
	}

	return count
}

func readOrbits(filename string) map[string]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	orbiting := make(map[string]string)
	for sc.Scan() {
		ln := strings.Split(sc.Text(), ")")
		orbiting[ln[1]] = ln[0] // key orbits value
	}

	return orbiting
}
