package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	moduleMasses := readNumbers("input")
	fuelSum, recursiveSum := 0, 0

	for _, mass := range moduleMasses {
		fuelSum += fuelRequired(mass, false)
		recursiveSum += fuelRequired(mass, true)
	}

	fmt.Println("-- Part 1:")
	fmt.Printf("Total fuel required: %v\n\n", fuelSum)
	fmt.Println("-- Part 2:")
	fmt.Printf("Total fuel required: %v\n\n", recursiveSum)

}

func readNumbers(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var numbers []int

	for sc.Scan() {
		value, err := strconv.Atoi(sc.Text())
		if err != nil {
			log.Fatalf("error reading numbers from file: %v", err)
		}
		numbers = append(numbers, value)
	}

	return numbers
}

func fuelRequired(mass int, recursive bool) (fuel int) {
	fuel = (mass / 3) - 2
	if !recursive {
		return
	}
	if fuel <= 0 {
		return 0
	} else {
		return fuel + fuelRequired(fuel, true)
	}
}
