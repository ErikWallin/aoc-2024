package main

import (
	"slices"
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

func Run(input string) int {
	words := common.ParseStringList(input, "\n")

	sum := 0
	for _, word := range words {
		np, _ := strings.CutSuffix(word, string('A'))
		numPart := common.MustAtoi(np)
		minLen := 10000000
		seq1 := getSequenceOnNumeric(word)
		for _, s1 := range seq1 {
			seq2 := getSequenceOnDirectional(s1)
			for _, s2 := range seq2 {
				seq3 := getSequenceOnDirectional(s2)
				for _, s3 := range seq3 {
					if len(s3) < minLen {
						minLen = len(s3)
					}
				}
			}
		}
		sum += numPart * minLen
	}
	return sum
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

	// Traverse and skip hope paths
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
	common.Run(Run, 2024, 21, 1, submit, verbose)
}
