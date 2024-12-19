package main

import (
	"strconv"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
	"github.com/ErikWallin/aoc-2024/internal/common/intcode"
)

func RunString(input string) string {
	sections := common.ParseStringList(input, "\n\n")
	r0, _ := strings.CutPrefix(strings.Split(sections[0], "\n")[0], "Register A: ")
	r1, _ := strings.CutPrefix(strings.Split(sections[0], "\n")[1], "Register B: ")
	r2, _ := strings.CutPrefix(strings.Split(sections[0], "\n")[2], "Register C: ")
	register := [3]int{common.MustAtoi(r0), common.MustAtoi(r1), common.MustAtoi(r2)}
	p, _ := strings.CutPrefix(sections[1], "Program: ")
	program := common.ParseIntList(p, ",")
	intCodeComputer := intcode.NewIntCodeComputerWithRegister(program, register)

	intCodeComputer.Run()

	stringOutput := make([]string, len(intCodeComputer.Output))
	for i, o := range intCodeComputer.Output {
		stringOutput[i] = strconv.Itoa(o)
	}

	return strings.Join(stringOutput, ",")
}

func Run(input string) int {
	RunString(input)
	return 0
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 17, 1, submit, verbose)
}
