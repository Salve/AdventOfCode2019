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
	p1 := patterns(len(input))
	fmt.Printf("%v\n\n", fft(input, 100, p1))

	fmt.Println("-- Part 2:")
	i2 := multSlice(input, 10000)
	p2 := patterns(len(i2))
	fmt.Printf("%v\n\n", fft(i2, 100, p2))
}

func patterns(len int) [][]int {
	pattern := []int{0, 1, 0, -1}
	lines := make([][]int, len)
	for i := 1; i <= len; i++ {
		line := make([]int, len+i+1)
		for j := 0; j*i <= len+1; j++ {
			for k := 0; k < i; k++ {
				line[j*i+k] = pattern[j%4]
			}
		}
		lines[i-1] = line[1:]
	}
	return lines
}

func fft(in []int, phase int, p [][]int) []int {
	out := make([]int, len(in))
	for i, _ := range in {
		sum := 0
		for j, v := range in {
			sum += v * p[i][j]
		}
		out[i] = Abs(sum % 10)
	}
	if phase == 1 {
		return out
	}
	return fft(out, phase-1, p)
}

func multSlice(s []int, x int) []int {
	out := make([]int, len(s)*x)
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
