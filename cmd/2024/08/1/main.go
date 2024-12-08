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
				d1 := a2.minus(a1)
				antinode := a1.minus(d1)
				if antinode.x >= 0 && antinode.x < lenX && antinode.y >= 0 && antinode.y < lenY {
					antinodes[antinode] = true
				}
				d2 := a1.minus(a2)
				antinode = a2.minus(d2)
				if antinode.x >= 0 && antinode.x < lenX && antinode.y >= 0 && antinode.y < lenY {
					antinodes[antinode] = true
				}
			}
		}
	}
	return len(antinodes)
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 8, 1, submit, verbose)
}
