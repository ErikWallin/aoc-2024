package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

type gate struct {
	input1    string
	input2    string
	operation string
	output    string
}

func Run(input string) int {
	fmt.Println(RunGates(input))
	return 0
}

func RunGates(input string) string {

	sections := common.ParseStringList(input, "\n\n")

	var gates []gate
	for _, r := range common.ParseStringList(sections[1], "\n") {
		ps := strings.Split(r, " -> ")
		output := ps[1]
		inputs := strings.Split(ps[0], " ")
		gates = append(gates, gate{inputs[0], inputs[2], inputs[1], output})
	}

	/*
			vdn OR qtn -> z45
			y44 AND x44 -> vdn
			nnt AND jnj -> qtn
			x44 XOR y44 -> nnt
			hsd OR nsf -> jnj

			cat cmd/2024/24/puzzle.txt | grep -z " -> z" | grep " XOR"
		    --> kth, gsd, tbt
			cat cmd/2024/24/puzzle.txt | grep " -> z" | grep -v " XOR" | grep -v "z45"
		    --> z32, z26, z12
			Errors: z32, z26, z12, kth, gsd, tbt
			x00 AND y00 -> bdj (rest from 00)
			y01 XOR x01 -> twd
			y01 AND x01 -> gwd
			twd AND bdj -> cbq
			twd XOR bdj -> z01
			cbq OR gwd -> rhr

			rsk XOR rhr -> z02

			rest-in = bdj
			z01 = twd XOR bdj = y01 XOR x01 XOR bdj
			rest-out = rhr = cbq OR gwd = (twd AND bdj) OR (y01 AND x01) = ((y01 XOR x01) AND x00 AND y00) OR (y01 AND x01)
	*/

	zOutputsNotFromXorErrors := []string{}
	for _, g := range gates {
		if _, isZ := strings.CutPrefix(g.output, "z"); isZ {
			if g.operation != "XOR" && g.output != "z45" {
				zOutputsNotFromXorErrors = append(zOutputsNotFromXorErrors, g.output)
			}
		}
	}

	xorOutputsNotZErrors := []string{}
	for _, g := range gates {
		if g.operation == "XOR" {
			_, isZOutput := strings.CutPrefix(g.output, "z")
			_, isXInput1 := strings.CutPrefix(g.input1, "x")
			_, isYInput1 := strings.CutPrefix(g.input1, "y")
			_, isXInput2 := strings.CutPrefix(g.input2, "x")
			_, isYInput2 := strings.CutPrefix(g.input2, "y")
			isXInput := isXInput1 || isXInput2
			isYInput := isYInput1 || isYInput2
			if !isZOutput && !(isXInput && isYInput) {
				xorOutputsNotZErrors = append(xorOutputsNotZErrors, g.output)
			}
		}
	}

	andOutputs := []string{}
	for _, g := range gates {
		if g.operation == "AND" {
			if !slices.Contains(zOutputsNotFromXorErrors, g.output) {
				andOutputs = append(andOutputs, g.output)
			}
		}
	}
	andOutputNotInOrInputErrors := []string{}
	for _, o := range andOutputs {
		inputInOr := false
		for _, g := range gates {
			if g.operation == "OR" && (g.input1 == o || g.input2 == o) {
				inputInOr = true
			}
		}
		if !inputInOr && o != "bdj" {
			andOutputNotInOrInputErrors = append(andOutputNotInOrInputErrors, o)
		}
	}

	orOutputs := []string{}
	for _, g := range gates {
		if g.operation == "OR" {
			if !slices.Contains(zOutputsNotFromXorErrors, g.output) && g.output != "z45" {
				orOutputs = append(orOutputs, g.output)
			}
		}
	}
	fmt.Println("orOutputs", len(orOutputs), orOutputs)
	orOutpurNotInXorInputWithZOutputErrors := []string{}
	for _, o := range orOutputs {
		inputInXorWithZOut := false
		inputInAnd := false
		for _, g := range gates {
			if g.operation == "XOR" && (g.input1 == o || g.input2 == o) {
				if _, isZ := strings.CutPrefix(g.output, "z"); isZ || slices.Contains(xorOutputsNotZErrors, g.output) {
					inputInXorWithZOut = true
				}
			} else if g.operation == "AND" && (g.input1 == o || g.input2 == o) {
				inputInAnd = true
			}
		}
		fmt.Println(inputInAnd, inputInXorWithZOut)
		if !(inputInXorWithZOut && inputInAnd) {
			orOutpurNotInXorInputWithZOutputErrors = append(orOutpurNotInXorInputWithZOutputErrors, o)
		}
	}
	fmt.Println("orOutpurNotInXorInputWithZOutputErrors", orOutpurNotInXorInputWithZOutputErrors)

	/*
		I could only find 7 erros with this method, but I have at least the on part
		of the last pair since I have 7.

		Manual matching of pairs

		Starting rest
		x00 AND y00 -> bdj (rest from 00)

		T0 calculate the next bit, this is what happens, which is locical in addition
		y01 XOR x01 -> twd
		y01 AND x01 -> gwd
		twd AND bdj -> cbq
		twd XOR bdj -> z01
		cbq OR gwd -> rhr

		Check them on by on by building up the sequence and see what is wrong

		nhb AND cdq -> nng
		x12 AND y12 -> psw
		psw OR nng -> z12 (should be rest to 13)

		x13 XOR y13 -> mtp
		y13 AND x13 -> qrs
		mtp AND kth -> cns
		mtp XOR kth -> z13
		cns OR qrs -> rpt

		This gives kth <-> z12

		y26 XOR x26 -> dfp
		x26 AND y26 -> z26
		mbg AND dfp -> kbg
		dfp XOR mbg -> gsd
		gsd OR kbg -> cmf
		vvv
		y26 XOR x26 -> dfp
		x26 AND y26 -> gsd
		mbg AND dfp -> kbg
		dfp XOR mbg -> gsd
		gsd OR kbg -> cmf

		This gives z26 <-> gsd

		x32 XOR y32 -> bkh
		y32 AND x32 -> skt
		vtg AND bkh -> z32
		vtg XOR bkh -> tbt
		tbt OR skt -> nwm

		This gives tbt <-> z32

		Left is qnf and the missing bit
		===
		y36 XOR x36 -> vpm
		y36 AND x36 -> qnf
		htb AND qnf -> thg
		htb XOR qnf -> z36
		thg OR vpm -> wkk

		This gives qnf <-> vpm

		vpm is the missing output
	*/

	orOutpurNotInXorInputWithZOutputErrors = append(orOutpurNotInXorInputWithZOutputErrors, "vpm")

	errors := []string{}
	errors = append(errors, zOutputsNotFromXorErrors...)
	errors = append(errors, xorOutputsNotZErrors...)
	errors = append(errors, andOutputNotInOrInputErrors...)
	errors = append(errors, orOutpurNotInXorInputWithZOutputErrors...)
	slices.Sort(errors)
	return strings.Join(errors, ",")
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 24, 2, submit, verbose)
}
