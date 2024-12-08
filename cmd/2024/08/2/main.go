package main

import (
	"github.com/ErikWallin/aoc-2024/internal/common"
)

type C struct {
	x, y int
}

func (c C) minus(other C) C {
	return C{c.x - other.x, c.y - other.y}
}

func Run(input string) int {
	runes := common.ParseRuneListList(input)
	lenY := len(runes)
	lenX := len(runes[0])
	antennas := map[rune][]C{}
	antinodes := map[C]bool{}
	for y, rs := range runes {
		for x, r := range rs {
			if r != '.' {
				antennas[r] = append(antennas[r], C{x, y})
			}
		}
	}
	for _, v := range antennas {
		for i, a1 := range v {
			for j, a2 := range v {
				if i == j {
					continue
				}

				start := a1
				diff := a2.minus(a1)
				for {
					antinode := start.minus(diff)
					if antinode.x >= 0 && antinode.x < lenX && antinode.y >= 0 && antinode.y < lenY {
						antinodes[antinode] = true
						start = antinode
					} else {
						break
					}
				}

				start = a2
				diff = a1.minus(a2)
				for {
					antinode := start.minus(diff)
					if antinode.x >= 0 && antinode.x < lenX && antinode.y >= 0 && antinode.y < lenY {
						antinodes[antinode] = true
						start = antinode
					} else {
						break
					}
				}
			}
		}
	}

	/*for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			if antinodes[C{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(string(runes[y][x]))
			}
		}
		fmt.Println()
	}*/

	antennasMoreThan1 := 0
	for _, as := range antennas {
		if len(as) != 1 {
			for _, a := range as {
				if _, ok := antinodes[a]; !ok {
					antennasMoreThan1++
				}
			}
		}
	}
	return len(antinodes) + antennasMoreThan1
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 8, 2, submit, verbose)
}
