package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readDigits("input")

	fmt.Println("-- Part 1:")
	fmt.Printf("%v\n\n", fft(input, 100)[:8])

	fmt.Println("-- Part 2:")
	i2 := multSlice(input, 10000)
	offset := i2[:7]
	fmt.Printf("%d\n%d\n", len(i2), offset)
	//fmt.Printf("%v\n\n", fft(i2, 1))

}

func fft(in []int, phase int) []int {
	out := make([]int, len(in))
	pattern := []int{0, 1, 0, -1}
	for i, _ := range in {
		sum := 0
		for j, v := range in {
			sum += v * pattern[(j+1)%(4*(i+1))/(i+1)]
		}
		out[i] = Abs(sum % 10)
	}
	if phase == 1 {
		return out
	}
	return fft(out, phase-1)
}

func multSlice(s []int, x int) []int {
	out := []int{}
	for i := 0; i < x; i++ {
		out = append(out, s...)
	}
	return out
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func readDigits(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	var result []int
	for scanner.Scan() {
		b := scanner.Bytes()
		if b[0] < 48 || b[0] > 57 {
			continue
		}
		result = append(result, int(b[0]-48))
	}
	return result
}
