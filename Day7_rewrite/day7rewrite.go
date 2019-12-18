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
	max := 0
	permute([]int{0, 1, 2, 3, 4}, func(p []int) {
		if out := tryPhaseSequence(p, inputprogram); out > max {
			max = out
		}
	})
	fmt.Printf("Max thruster signal achieved: %d\n\n", max)

	fmt.Println("-- Part 2:")
	max = 0
	permute([]int{5, 6, 7, 8, 9}, func(p []int) {
		if out := tryFeedbackSequence(p, inputprogram); out > max {
			max = out
		}
	})
	fmt.Printf("Max thruster signal achieved: %d\n\n", max)

}

func tryPhaseSequence(seq []int, prg []int) int {
	signal := 0
	for i := range seq {
		p := intcode.Copy(prg)
		in := make(chan int, 2)
		out := make(chan int)
		halt := make(chan bool)
		reqin := make(chan bool, 1) // buffered, won't be used
		halted := false

		go intcode.Run(p, in, out, halt, reqin)

		in <- seq[i]
		in <- signal
		for {
			select {
			case <-halt:
				halted = true
			case o := <-out:
				signal = o
			}
			if halted {
				break
			}
		}
	}
	return signal
}

func tryFeedbackSequence(seq []int, prg []int) (signal int) {
	ampA := intcode.Copy(prg)
	ampB := intcode.Copy(prg)
	ampC := intcode.Copy(prg)
	ampD := intcode.Copy(prg)
	ampE := intcode.Copy(prg)

	chAB := make(chan int)
	chBC := make(chan int)
	chCD := make(chan int)
	chDE := make(chan int)
	chEA := make(chan int)

	halt := make(chan bool)
	reqin := make(chan bool, 1) // buffered, won't be used

	go intcode.Run(ampA, chEA, chAB, halt, reqin)
	go intcode.Run(ampB, chAB, chBC, halt, reqin)
	go intcode.Run(ampC, chBC, chCD, halt, reqin)
	go intcode.Run(ampD, chCD, chDE, halt, reqin)
	go intcode.Run(ampE, chDE, chEA, halt, reqin)

	// set phase settings
	chEA <- seq[0]
	chAB <- seq[1]
	chBC <- seq[2]
	chCD <- seq[3]
	chDE <- seq[4]

	// initiate feedback sequence
	chEA <- 0

	halted := 0
	for {
		select {
		case <-halt:
			halted++
		default:
			if halted == 4 {
				signal = <-chEA // read final signal value
				halted++        // avoid deadlock, wait for final amp to halt
			}
			if halted == 6 {
				return signal
			}
		}
	}

}

// https://rosettacode.org/wiki/Permutations#Go
func permute(s []int, emit func([]int)) {
	// permute function.  takes a set to permute and a function
	// to call for each generated permutation.

	if len(s) == 0 {
		emit(s)
		return
	}
	// Steinhaus, implemented with a recursive closure.
	// arg is number of positions left to permute.
	// pass in len(s) to start generation.
	// on each call, weave element at pp through the elements 0..np-2,
	// then restore array to the way it was.
	var rc func(int)
	rc = func(np int) {
		if np == 1 {
			emit(s)
			return
		}
		np1 := np - 1
		pp := len(s) - np1
		// weave
		rc(np1)
		for i := pp; i > 0; i-- {
			s[i], s[i-1] = s[i-1], s[i]
			rc(np1)
		}
		// restore
		w := s[0]
		copy(s, s[1:pp+1])
		s[pp] = w
	}
	rc(len(s))
}
