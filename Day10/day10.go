package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type asteroid struct {
	x, y int
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	fmt.Println("-- Part 1:")
	a := readMap(f)
	var maxA asteroid
	var maxV int
	var blocked bool

	// Incredibly inefficient......
	for k, _ := range a { // potential observatory asteroid
		v := 0
		for k2, _ := range a { // asteroid we want to test visibility of
			if k2 == k {
				continue
			}
			blocked = false
			for k3, _ := range a { // asteroid that might be blocking line of sight
				if k3 == k2 || k3 == k {
					continue
				}
				if k3.between(k, k2) {
					blocked = true
					break
				}
			}
			if blocked == false {
				v++
			}
		}
		if v > maxV {
			maxV = v
			maxA = k
		}
	}

	fmt.Printf("Best location is (%d, %d), with a view of %d asteroids\n\n", maxA.x, maxA.y, maxV)

}

func (c *asteroid) between(a, b asteroid) bool {
	// True if c lies between a and b
	return collinear(a, b, *c) && within(a.x, c.x, b.x) && within(a.y, c.y, b.y)
}
func collinear(a, b, c asteroid) bool {
	// True if a, b and c are on same line
	return (b.x-a.x)*(c.y-a.y) == (c.x-a.x)*(b.y-a.y)
}
func within(p, q, r int) bool {
	// True if q between p and r
	return (p <= q && q <= r) || (r <= q && q <= p)
}

func readMap(f io.Reader) map[asteroid]struct{} {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	asteroids := map[asteroid]struct{}{}
	x, y := 0, 0
	for scanner.Scan() {
		for _, b := range scanner.Bytes() {
			if b == "#"[0] {
				asteroids[asteroid{x: x, y: y}] = struct{}{}
			}
			x++
		}
		x = 0
		y++
	}
	return asteroids
}
