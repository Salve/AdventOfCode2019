package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputfile = flag.String("input", "", "path to input file")

func main() {
	flag.Parse()
	f, err := os.OpenFile(*inputfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	fuelSum := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		value, err := strconv.Atoi(sc.Text())
		if err == nil {
			fuelSum += fuelRequired(value)
		}
	}
	fmt.Printf("Total fuel required: %v", fuelSum)

}

func fuelRequired(mass int) int {
	return (mass / 3) - 2
}
