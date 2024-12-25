package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

var start map[string]bool = map[string]bool{}
var xs []string
var ys []string
var zs []string

type gate struct {
	input1    string
	input2    string
	operation string
	output    string
}

func (g gate) process() bool {
	if _, ok := start[g.input1]; !ok {
		return false
	}
	if _, ok := start[g.input2]; !ok {
		return false
	}
	if g.operation == "AND" {
		start[g.output] = start[g.input1] && start[g.input2]
	} else if g.operation == "OR" {
		start[g.output] = start[g.input1] || start[g.input2]
	} else if g.operation == "XOR" {
		start[g.output] = start[g.input1] != start[g.input2]
	} else {
		panic(g)
	}
	return true
}

var gates []gate

func Run(input string) int {

	sections := common.ParseStringList(input, "\n\n")

	for _, r := range common.ParseStringList(sections[0], "\n") {
		ps := strings.Split(r, ": ")
		if x, ok := strings.CutPrefix(ps[0], "x"); ok {
			xs = append(xs, x)
		}
		if y, ok := strings.CutPrefix(ps[0], "y"); ok {
			ys = append(ys, y)
		}
		if ps[1] == "1" {
			start[ps[0]] = true
		} else {
			start[ps[0]] = false
		}
	}
	xValue := 0
	for i, x := range xs {
		if start[x] {
			xValue += int(math.Pow(2, float64(i)))
		}
	}
	yValue := 0
	for i, x := range ys {
		if start[x] {
			yValue += int(math.Pow(2, float64(i)))
		}
	}
	fmt.Println(xValue, yValue)
	for _, r := range common.ParseStringList(sections[1], "\n") {
		ps := strings.Split(r, " -> ")
		output := ps[1]
		inputs := strings.Split(ps[0], " ")
		gates = append(gates, gate{inputs[0], inputs[2], inputs[1], output})
		if _, ok := strings.CutPrefix(output, "z"); ok {
			zs = append(zs, output)
		}
	}
	slices.Sort(zs)

	zsDone := []string{}
	gatesLeft := gates
	for len(zsDone) < len(zs) {
		gatesNotProcessed := []gate{}
		for _, gate := range gatesLeft {
			r := gate.process()
			if r {
				if slices.Contains(zs, gate.output) {
					zsDone = append(zsDone, gate.output)
				}
			} else {
				gatesNotProcessed = append(gatesNotProcessed, gate)
			}
		}
		gatesLeft = gatesNotProcessed
	}

	result := float64(0)
	for i, z := range zs {
		if start[z] {
			result += math.Pow(2, float64(i))
		}
	}
	return int(result)
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 24, 1, submit, verbose)
}
