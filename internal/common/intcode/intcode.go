package intcode

import (
	"fmt"
	"math"
)

type IntCodeComputer struct {
	Program        []int
	Register       [3]int
	ProgramPointer int
	Output         []int
}

func NewIntCodeComputer(program []int) IntCodeComputer {
	return IntCodeComputer{program, [3]int{0, 0, 0}, 0, nil}
}

func NewIntCodeComputerWithRegister(program []int, register [3]int) IntCodeComputer {
	return IntCodeComputer{program, register, 0, nil}
}

func (ic *IntCodeComputer) Run() {
	for ic.ProgramPointer < len(ic.Program) {
		op := ic.Program[ic.ProgramPointer]
		switch op {
		case 0:
			ic.adv()
		case 1:
			ic.bxl()
		case 2:
			ic.bst()
		case 3:
			ic.jnz()
		case 4:
			ic.bxc()
		case 5:
			ic.out()
		case 6:
			ic.bdv()
		case 7:
			ic.cdv()
		default:
			panic(fmt.Sprint("op", op))
		}
	}
}

func (ic *IntCodeComputer) comboValue(combo int) int {
	if combo <= 3 {
		return combo
	} else if combo == 4 {
		return ic.Register[0]
	} else if combo == 5 {
		return ic.Register[1]
	} else if combo == 6 {
		return ic.Register[2]
	} else {
		panic(fmt.Sprint("comboValue", combo))
	}
}

func (ic *IntCodeComputer) adv() {
	denominator := math.Pow(2, float64(ic.comboValue(ic.Program[ic.ProgramPointer+1])))
	ic.Register[0] = ic.Register[0] / int(denominator)
	ic.ProgramPointer += 2
}

func (ic *IntCodeComputer) bxl() {
	ic.Register[1] = ic.Register[1] ^ ic.Program[ic.ProgramPointer+1]
	ic.ProgramPointer += 2
}

func (ic *IntCodeComputer) bst() {
	ic.Register[1] = ic.comboValue(ic.Program[ic.ProgramPointer+1]) % 8
	ic.ProgramPointer += 2
}

func (ic *IntCodeComputer) jnz() {
	if ic.Register[0] == 0 {
		ic.ProgramPointer += 2
	} else {
		ic.ProgramPointer = ic.Program[ic.ProgramPointer+1]
	}
}

func (ic *IntCodeComputer) bxc() {
	ic.Register[1] = ic.Register[1] ^ ic.Register[2]
	ic.ProgramPointer += 2
}

func (ic *IntCodeComputer) out() {
	ic.Output = append(ic.Output, ic.comboValue(ic.Program[ic.ProgramPointer+1])%8)
	ic.ProgramPointer += 2
}

func (ic *IntCodeComputer) bdv() {
	denominator := math.Pow(2, float64(ic.comboValue(ic.Program[ic.ProgramPointer+1])))
	ic.Register[1] = ic.Register[0] / int(denominator)
	ic.ProgramPointer += 2
}

func (ic *IntCodeComputer) cdv() {
	denominator := math.Pow(2, float64(ic.comboValue(ic.Program[ic.ProgramPointer+1])))
	ic.Register[2] = ic.Register[0] / int(denominator)
	ic.ProgramPointer += 2
}
