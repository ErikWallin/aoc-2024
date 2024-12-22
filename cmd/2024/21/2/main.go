package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
	. "github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

var dir map[C]rune = map[C]rune{
	Up:    '^',
	Down:  'v',
	Left:  '<',
	Right: '>',
}

var numeric [4][3]rune = [4][3]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'-', '0', 'A'},
}

var numericPositions map[rune]C = func() map[rune]C {
	np := map[rune]C{}
	for y, row := range numeric {
		for x, r := range row {
			np[r] = C{x, y}
		}
	}
	return np
}()

var directional [2][3]rune = [2][3]rune{
	{'-', '^', 'A'},
	{'<', 'v', '>'},
}

var directionalPositions map[rune]C = func() map[rune]C {
	np := map[rune]C{}
	for y, row := range directional {
		for x, r := range row {
			np[r] = C{x, y}
		}
	}
	return np
}()

func Run25(input string) int {
	return Run(input, 25)
}

func Run(input string, robots int) int {
	words := common.ParseStringList(input, "\n")
	sum := 0
	for _, word := range words {
		np, _ := strings.CutSuffix(word, string('A'))
		numPart := common.MustAtoi(np)
		seqs := getSequenceOnNumeric(word)
		var values []int
		for _, seq := range seqs {
			values = append(values, minLength(seq, robots))
		}
		min := slices.Min(values)
		sum += numPart * min
	}
	return sum
}

var minLengthCache map[string]int = map[string]int{}

func minLength(sequence string, depth int) int {
	key := sequence + strconv.Itoa(depth)
	if val, ok := minLengthCache[key]; ok {
		return val
	}
	var res int
	if depth == 0 {
		res = len(sequence)
	} else {
		res = 0
		parts := strings.Split(sequence, "A")
		for i, part := range parts {
			if i == len(parts)-1 {
				continue
			}
			values := []int{}
			newSequences := getSequenceOnDirectional(part + "A")
			for _, s := range newSequences {
				values = append(values, minLength(s, depth-1))
			}
			res += slices.Min(values)
		}
	}
	minLengthCache[key] = res
	return res
}

func getSequenceOnNumeric(numericSequence string) []string {
	sequences := []string{}
	for i, r := range numericSequence {
		if i == 0 {
			for _, s := range getSequence(numericPositions, 'A', r) {
				sequences = append(sequences, s+"A")
			}
		} else {
			newRes := []string{}
			for _, s := range getSequence(numericPositions, rune(numericSequence[i-1]), r) {
				for _, r := range sequences {
					newRes = append(newRes, r+s+"A")
				}
			}
			sequences = newRes
		}
	}
	return sequences
}

func getSequenceOnDirectional(directionalSequence string) []string {
	res := []string{}
	for i, r := range directionalSequence {
		if i == 0 {
			for _, s := range getSequence(directionalPositions, 'A', r) {
				res = append(res, s+"A")
			}
		} else {
			newRes := []string{}
			for _, s := range getSequence(directionalPositions, rune(directionalSequence[i-1]), r) {
				for _, r := range res {
					newRes = append(newRes, r+s+"A")
				}
			}
			res = newRes
		}
	}
	return res
}

var sequenceCache map[string][]string = map[string][]string{}

func getSequence(positions map[rune]C, start, end rune) []string {
	key := string(start) + string(end)
	if val, ok := sequenceCache[key]; ok {
		return val
	}
	startC := positions[start]
	endC := positions[end]
	hole := positions['-']
	diff := endC.Minus(startC)
	var dirs []C
	if diff.X < 0 {
		for range -diff.X {
			dirs = append(dirs, Left)
		}
	} else if diff.X > 0 {
		for range diff.X {
			dirs = append(dirs, Right)
		}
	}
	if diff.Y < 0 {
		for range -diff.Y {
			dirs = append(dirs, Up)
		}
	} else if diff.Y > 0 {
		for range diff.Y {
			dirs = append(dirs, Down)
		}
	}
	var res []string

	// Traverse and skip hole paths
	current := startC
	success := true
	for _, d := range dirs {
		current.Add(d)
		if current == hole {
			success = false
			break
		}
	}
	if success {
		r := ""
		for _, d := range dirs {
			r += string(dir[d])
		}
		res = append(res, r)
	}
	// Add the other path (up/down first)
	if startC.X != endC.X && startC.Y != endC.Y {
		slices.Reverse(dirs)
		current = startC
		success = true
		for _, d := range dirs {
			current.Add(d)
			if current == hole {
				success = false
				break
			}
		}
		if success {
			r := ""
			for _, d := range dirs {
				r += string(dir[d])
			}
			res = append(res, r)
		}
	}
	sequenceCache[key] = res
	return res
}

func main() {
	submit := false
	verbose := true
	common.Run(Run25, 2024, 21, 2, submit, verbose)
}
