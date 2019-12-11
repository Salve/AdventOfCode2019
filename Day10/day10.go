package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
)

type asteroid struct {
	x, y  int
	angle float64
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	fmt.Println("-- Part 1:")
	roids := readMap(f)
	var v, maxV int // visible asteroids
	var maxA asteroid

	// This is very inefficient, but works...
	for _, roid := range roids { // potential observatory asteroid
		v = len(roid.visibleAsteroids(roids))
		if v > maxV {
			maxV = v
			maxA = roid
		}
	}

	fmt.Printf("Best location is (%d, %d), with a view of %d asteroids\n\n", maxA.x, maxA.y, maxV)

	fmt.Println("-- Part 2:")
	visible := maxA.visibleAsteroids(roids)

	var nq2 []asteroid // not 2. quadrant
	var q2 []asteroid  // 2. quadrant

	for _, roid := range visible {
		ang := roid.angleFrom(maxA)
		if ang > math.Pi/2 {
			// angle is in 2. quadrant and has to be sorted separately
			q2 = append(q2, asteroid{roid.x, roid.y, ang})
		} else {
			nq2 = append(nq2, asteroid{roid.x, roid.y, ang})
		}
	}

	// Sort quadrants 1, 3 and 4 descending (up going clockwise to left = 1/2pi -> 0 -> -1/2pi -> -pi)
	sort.Slice(nq2, func(i, j int) bool { return nq2[i].angle > nq2[j].angle })

	// Sort quadrant 2 descending (left going clockwise to up = pi -> 1/2pi)
	sort.Slice(q2, func(i, j int) bool { return q2[i].angle > q2[j].angle })

	sAsteroids := append(nq2, q2...)
	fmt.Printf("200th asteroid to be exploded is (%d, %d), answer: %d",
		sAsteroids[199].x, sAsteroids[199].y, sAsteroids[199].x*100+sAsteroids[199].y)

}

func (a *asteroid) visibleAsteroids(roids []asteroid) (vroids []asteroid) {
	for _, roid := range roids { // asteroid we want to test visibility of
		if roid.visibleFrom(*a, roids) {
			vroids = append(vroids, roid)
		}
	}
	return
}

func (a *asteroid) visibleFrom(b asteroid, roids []asteroid) bool {
	if a.x == b.x && a.y == b.y {
		return false // asteroid is not visible from itself
	}
	for _, roid := range roids {
		if (roid.x == a.x && roid.y == a.y) || (roid.x == b.x && roid.y == b.y) {
			continue // visiblity between a and b can't be blocked by themselves
		}
		if roid.between(b, *a) {
			return false
		}
	}
	return true
}

func (a *asteroid) angleFrom(b asteroid) float64 {
	relative := a.minus(b)
	// negative y as our coordinates are positive below the x axis
	return math.Atan2(float64(-relative.y), float64(relative.x))
}

func (a *asteroid) minus(b asteroid) asteroid {
	return asteroid{
		x: a.x - b.x,
		y: a.y - b.y,
	}
}

func (a *asteroid) between(b, c asteroid) bool {
	// True if a lies between b and c
	return collinear(b, c, *a) && within(b.x, a.x, c.x) && within(b.y, a.y, c.y)
}
func collinear(a, b, c asteroid) bool {
	// True if a, b and c are on same line
	return (b.x-a.x)*(c.y-a.y) == (c.x-a.x)*(b.y-a.y)
}
func within(p, q, r int) bool {
	// True if q between p and r
	return (p <= q && q <= r) || (r <= q && q <= p)
}

func readMap(f io.Reader) []asteroid {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var asteroids []asteroid
	x, y := 0, 0
	for scanner.Scan() {
		for _, b := range scanner.Bytes() {
			if b == "#"[0] {
				asteroids = append(asteroids, asteroid{x: x, y: y})
			}
			x++
		}
		x = 0
		y++
	}
	return asteroids
}
