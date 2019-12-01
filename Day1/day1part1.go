package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var inputfile = flag.String("input", "", "path to input file")

func main() {
	f, err := os.OpenFile(*inputfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		os.Exit(1)
	}

}
