package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type Material interface {
	Request(int, map[string]Material)
	Demand() int
}
type component struct {
	name   string
	amount int
}
type reaction struct {
	batchsize int
	input     []component
	demand    int
	runs      int
}
type minable struct {
	demand int
}

func (o *minable) Request(amount int, reactions map[string]Material) {
	o.demand += amount
}
func (o *minable) Demand() int {
	return o.demand
}

func (r *reaction) Request(amount int, reactions map[string]Material) {
	missing := amount - (r.runs*r.batchsize - r.demand)
	if missing > 0 {
		missingRuns := (missing + r.batchsize - 1) / r.batchsize // int division ceil
		for _, i := range r.input {
			reactions[i.name].Request(i.amount*missingRuns, reactions)
		}
		r.runs += missingRuns
	}
	r.demand += amount
}
func (r *reaction) Demand() int {
	return r.demand
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Could not open input file: %s\n", err)
	}

	fmt.Println("-- Part 1:")
	p1r := ReadReactions(bytes.NewReader(input))
	p1r["FUEL"].Request(1, p1r)
	ore1fuel := p1r["ORE"].Demand()
	fmt.Printf("Ore required for 1 unit of fuel: %d\n\n", ore1fuel)

	fmt.Println("-- Part 2:")
	lo := 1e12 / ore1fuel
	hi := int(1e12)
	for lo < hi {
		mid := (lo + hi + 1) / 2
		p2r := ReadReactions(bytes.NewReader(input))
		p2r["FUEL"].Request(mid, p2r)
		if p2r["ORE"].Demand() <= 1e12 {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	fmt.Printf("Amount of fuel that can be produced with a trillion ore: %d\n\n", lo)

}

func ReadReactions(i io.Reader) map[string]Material {
	reactions := make(map[string]Material)
	sc := bufio.NewScanner(i)
	re := regexp.MustCompile(`(\d+) ([A-Z]+)`)
	for sc.Scan() {
		r := re.FindAllStringSubmatch(sc.Text(), -1)
		productName := r[len(r)-1][2]
		productAmount, _ := strconv.Atoi(r[len(r)-1][1])
		re := reaction{
			batchsize: productAmount,
			input:     []component{},
			demand:    0,
			runs:      0,
		}
		for _, input := range r[:len(r)-1] {
			inputName := input[2]
			inputAmount, _ := strconv.Atoi(input[1])
			re.input = append(re.input, component{inputName, inputAmount})
		}
		reactions[productName] = &re
	}
	reactions["ORE"] = &minable{demand: 0}
	return reactions
}
