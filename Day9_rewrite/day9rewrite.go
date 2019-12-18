package main

import (
	intcode "AdventOfCode2019/Intcode"
	"fmt"
	"log"
)

func main() {
	inputprogram, err := intcode.FromFile("input")
	if err != nil {
		log.Fatalf("Failed to read intcode from file: %s\n", err)
	}

	fmt.Println("-- Part 1:")
	p1 := intcode.Copy(inputprogram)
	fmt.Printf("BOOST keycode: %d\n\n", runInOut(p1, 1))
	fmt.Println("-- Part 2:")
	p2 := intcode.Copy(inputprogram)
	fmt.Printf("BOOST keycode: %d\n\n", runInOut(p2, 2))

}

func runInOut(p []int, input int) (output []int) {
	in := make(chan int, 1)
	out := make(chan int)
	halt := make(chan bool)
	reqin := make(chan bool)

	go intcode.Run(p, in, out, halt, reqin)

	for {
		select {
		case <-halt:
			return output
		case <-reqin:
			in <- input
		case o := <-out:
			output = append(output, o)
		}
	}
}
