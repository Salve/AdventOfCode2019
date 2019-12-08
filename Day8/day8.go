package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type image struct {
	layers [][][]byte
	width  int
	height int
}

func (img *image) fromBytes(b []byte) {
	pxPerLayer := img.width * img.height
	if len(b)%pxPerLayer != 0 {
		log.Fatalf("Invalid number of pixels for %dx%d layers: %d", img.width, img.height, len(b))
	}
	datalen := len(b)
	for i := 0; i < datalen/pxPerLayer; i++ {
		if img.layers == nil {
			img.layers = make([][][]byte, datalen/pxPerLayer)
		}
		for j := 0; j < img.height; j++ {
			if img.layers[i] == nil {
				img.layers[i] = make([][]byte, img.height)
			}
			img.layers[i][j] = b[0:img.width]
			b = b[img.width:]
		}
	}
}

func (img *image) countLayers(search byte) []int {
	count := make([]int, len(img.layers))
	for i, row := range img.layers {
		for _, column := range row {
			for _, digit := range column {
				if digit == search {
					count[i]++
				}
			}
		}
	}
	return count
}

func main() {
	fmt.Println("-- Part 1:")
	img := image{width: 25, height: 6}
	img.fromBytes(readBytes("input"))
	layer := minSlice(img.countLayers(0))
	fmt.Printf("Result: %d\n\n", img.countLayers(1)[layer]*img.countLayers(2)[layer])

}

func minSlice(s []int) (index int) {
	var min int
	for i, e := range s {
		if i == 0 || e < min {
			min = e
			index = i
		}
	}
	return
}

func readBytes(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	var result []byte
	for scanner.Scan() {
		b := scanner.Bytes()
		if b[0] < 48 || b[0] > 57 {
			continue
		}
		result = append(result, b[0]-48)
	}
	return result
}
