package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type x3d struct {
	x, y, z int
}
type moon struct {
	pos x3d
	vel x3d
}

func (m *moon) Velocity() {
	m.pos.x += m.vel.x
	m.pos.y += m.vel.y
	m.pos.z += m.vel.z
}

func Gravity(a *moon, b *moon) {
	switch {
	case a.pos.x > b.pos.x:
		a.vel.x--
		b.vel.x++
	case a.pos.x < b.pos.x:
		a.vel.x++
		b.vel.x--
	}
	switch {
	case a.pos.y > b.pos.y:
		a.vel.y--
		b.vel.y++
	case a.pos.y < b.pos.y:
		a.vel.y++
		b.vel.y--
	}
	switch {
	case a.pos.z > b.pos.z:
		a.vel.z--
		b.vel.z++
	case a.pos.z < b.pos.z:
		a.vel.z++
		b.vel.z--
	}
}

func (m *moon) Energy() int {
	potential := IntAbs(m.pos.x) + IntAbs(m.pos.y) + IntAbs(m.pos.z)
	kinetic := IntAbs(m.vel.x) + IntAbs(m.vel.y) + IntAbs(m.vel.z)
	return potential * kinetic
}

func TotalEnergy(moons []moon) (sum int) {
	for _, m := range moons {
		sum += m.Energy()
	}
	return sum
}

func Simulate(moons []moon) {
	for i := 0; i < len(moons); i++ {
		for j := i + 1; j < len(moons); j++ {
			Gravity(&moons[i], &moons[j])
		}
		moons[i].Velocity()
	}
}

func EqualX(m1 []moon, m2 []moon) bool {
	for i := range m1 {
		if m1[i].pos.x != m2[i].pos.x || m1[i].vel.x != m2[i].vel.x {
			return false
		}
	}
	return true
}
func EqualY(m1 []moon, m2 []moon) bool {
	for i := range m1 {
		if m1[i].pos.y != m2[i].pos.y || m1[i].vel.y != m2[i].vel.y {
			return false
		}
	}
	return true
}
func EqualZ(m1 []moon, m2 []moon) bool {
	for i := range m1 {
		if m1[i].pos.z != m2[i].pos.z || m1[i].vel.z != m2[i].vel.z {
			return false
		}
	}
	return true
}

func CopyMoons(m []moon) []moon {
	c := make([]moon, len(m))
	copy(c, m)
	return c
}

func main() {
	i, err := os.Open("input")
	if err != nil {
		log.Fatalf("Could not open input file: %s\n", err)
	}
	defer i.Close()
	inputmoons := MakeMoons(i)

	p1moons := CopyMoons(inputmoons)
	fmt.Println("-- Part 1:")
	for i := 0; i < 1000; i++ {
		Simulate(p1moons)
	}
	fmt.Printf("Total energy after 1000 steps: %d\n\n", TotalEnergy(p1moons))

	fmt.Println("-- Part 2:")
	p2moons := CopyMoons(inputmoons)
	step := 0
	xp, yp, zp := 0, 0, 0
	for xp == 0 || yp == 0 || zp == 0 {
		Simulate(p2moons)
		step++
		if xp == 0 && EqualX(inputmoons, p2moons) {
			xp = step
			fmt.Printf("X period found: %d\n", xp)
		}
		if yp == 0 && EqualY(inputmoons, p2moons) {
			yp = step
			fmt.Printf("Y period found: %d\n", yp)
		}
		if zp == 0 && EqualZ(inputmoons, p2moons) {
			zp = step
			fmt.Printf("Z period found: %d\n", zp)
		}
	}
	fmt.Printf("System repeating after: %d steps\n\n", LCM(xp, yp, zp))

}

func MakeMoons(i io.Reader) (moons []moon) {
	sc := bufio.NewScanner(i)
	re := regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)
	for sc.Scan() {
		r := re.FindStringSubmatch(sc.Text())
		x, _ := strconv.Atoi(r[1])
		y, _ := strconv.Atoi(r[2])
		z, _ := strconv.Atoi(r[3])
		moons = append(moons, moon{pos: x3d{x, y, z}, vel: x3d{0, 0, 0}})
	}

	return moons
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, ints ...int) int {
	res := a * b / GCD(a, b)
	for i := 0; i < len(ints); i++ {
		res = LCM(res, ints[i])
	}
	return res
}

func IntAbs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
