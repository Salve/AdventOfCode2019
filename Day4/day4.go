package main

import (
	"fmt"
	"strconv"
)

const inputStart int = 245182
const inputEnd int = 790572

func main() {
	var numValid int
	var numValid2 int
	for i := inputStart; i <= inputEnd; i++ {
		if isValid(i) {
			numValid++
		}
		if isValid2(i) {
			numValid2++
		}
	}
	fmt.Println("-- Part 1:")
	fmt.Printf("Valid inputs: %d\n\n", numValid)
	fmt.Println("-- Part 2:")
	fmt.Printf("Valid inputs: %d\n\n", numValid2)
}

func isValid(v int) bool {
	doubleSeen := false
	var lastDigit int

	for _, e := range strconv.Itoa(v) {
		digit, _ := strconv.Atoi(string(e))
		if digit < lastDigit {
			return false
		}
		if digit == lastDigit {
			doubleSeen = true
		}
		lastDigit = digit
	}

	return doubleSeen
}

func isValid2(v int) bool {
	sequenceLen := 1
	doubleConfirmed := false
	var lastDigit int

	for _, e := range strconv.Itoa(v) {
		digit, _ := strconv.Atoi(string(e))
		if digit < lastDigit {
			return false
		}
		if sequenceLen == 2 && digit != lastDigit {
			doubleConfirmed = true
		}
		if digit == lastDigit {
			sequenceLen++
		} else {
			sequenceLen = 1
		}
		lastDigit = digit
	}
	return doubleConfirmed || sequenceLen == 2
}
