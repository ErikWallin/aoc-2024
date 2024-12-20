package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

var register [3]int
var program []int
var p int
var output []int

func comboValue(combo int) int {
	if combo <= 3 {
		return combo
	} else if combo == 4 {
		return register[0]
	} else if combo == 5 {
		return register[1]
	} else if combo == 6 {
		return register[2]
	} else {
		panic(fmt.Sprint("comboValue", combo))
	}
}

func adv() {
	denominator := math.Pow(2, float64(comboValue(program[p+1])))
	register[0] = register[0] / int(denominator)
	p += 2
}

func bxl() {
	register[1] = register[1] ^ program[p+1]
	p += 2
}

func bst() {
	register[1] = comboValue(program[p+1]) % 8
	p += 2
}

func jnz() {
	if register[0] == 0 {
		p += 2
	} else {
		p = program[p+1]
	}
}

func bxc() {
	register[1] = register[1] ^ register[2]
	p += 2
}

func out() {
	output = append(output, comboValue(program[p+1])%8)
	p += 2
}

func bdv() {
	denominator := math.Pow(2, float64(comboValue(program[p+1])))
	register[1] = register[0] / int(denominator)
	p += 2
}

func cdv() {
	denominator := math.Pow(2, float64(comboValue(program[p+1])))
	register[2] = register[0] / int(denominator)
	p += 2
}

func Run(input string) int {
	sections := common.ParseStringList(input, "\n\n")
	r, _ := strings.CutPrefix(strings.Split(sections[0], "\n")[0], "Register A: ")
	register[0] = common.MustAtoi(r)
	r, _ = strings.CutPrefix(strings.Split(sections[0], "\n")[1], "Register B: ")
	register[1] = common.MustAtoi(r)
	register1Origin := register[1]
	r, _ = strings.CutPrefix(strings.Split(sections[0], "\n")[2], "Register C: ")
	register[2] = common.MustAtoi(r)
	register2Origin := register[2]
	r, _ = strings.CutPrefix(sections[1], "Program: ")
	program = common.ParseIntList(r, ",")

	//a := 5
	//a := 61
	//a := 3773
	a := 15560381

	for {
		p = 0
		output = []int{}
		register[0] = a
		register[1] = register1Origin
		register[2] = register2Origin

		for p < len(program) && len(output) <= len(program) {
			if program[p] == 0 {
				adv()
			} else if program[p] == 1 {
				bxl()
			} else if program[p] == 2 {
				bst()
			} else if program[p] == 3 {
				jnz()
			} else if program[p] == 4 {
				bxc()
			} else if program[p] == 5 {
				out()
				i := len(output) - 1
				if output[i] != program[i] {
					/*if len(output) > 12 {
						fmt.Println(a)
					}*/
					break
				}
			} else if program[p] == 6 {
				bdv()
			} else if program[p] == 7 {
				cdv()
			} else {
				panic(fmt.Sprint("op", program[p]))
			}
		}
		if len(output) == len(program) {
			success := true
			for i, o := range output {
				if o != program[i] {
					success = false
					break
				}
			}
			if success {
				return a
			}
		}
		//a += 8
		//a += 64
		//a += 4096
		a += 16777216
	}
}

func main() {
	// 164279033360061 That's not the right answer; your answer is too high.
	// 164279024971453
	submit := false
	verbose := true
	common.Run(Run, 2024, 17, 2, submit, verbose)
}
