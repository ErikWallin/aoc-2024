package main

import (
	"slices"

	"github.com/ErikWallin/aoc-2024/internal/common"
	. "github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

type region struct {
	plant rune
	cs    []C
}

func (r *region) add(c C) {
	r.cs = append(r.cs, c)
}

func (r region) area() int {
	return len(r.cs)
}

func (r region) perimeter() int {
	p := 0
	for _, c := range r.cs {
		for _, n := range c.AllNeighbours(false) {
			if !slices.Contains(r.cs, n) {
				p++
			}
		}
	}
	return p
}

var rm map[C]*region
var rs []*region

func Run(input string) int {
	runes := common.ParseRuneListList(input)
	rm = map[C]*region{}
	for y, row := range runes {
		for x, p := range row {
			c := C{X: x, Y: y}
			lc := c.Minus(C{1, 0})
			uc := c.Minus(C{0, 1})
			if x > 0 && y > 0 && rm[lc].plant == p && rm[uc].plant == p && rm[lc] != rm[uc] {
				removedRegion := rm[uc]
				mergedRegion := rm[lc]
				mergedRegion.add(c)
				rm[c] = mergedRegion
				for _, c := range removedRegion.cs {
					mergedRegion.add(c)
				}
				toBeMoved := []C{}
				for cRemoved, r := range rm {
					if r == removedRegion {
						toBeMoved = append(toBeMoved, cRemoved)
					}
				}
				for _, cRemoved := range toBeMoved {
					rm[cRemoved] = mergedRegion
				}
				idx := slices.Index(rs, removedRegion)
				rs = slices.Delete(rs, idx, idx+1)
			} else if x > 0 && rm[lc].plant == p {
				rm[lc].add(c)
				rm[c] = rm[lc]
			} else if y > 0 && rm[uc].plant == p {
				rm[uc].add(c)
				rm[c] = rm[uc]
			} else {
				r := region{p, []C{c}}
				rs = append(rs, &r)
				rm[c] = &r
			}
		}
	}
	sum := 0
	for _, r := range rs {
		sum += r.area() * r.perimeter()
	}
	return sum
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 12, 1, submit, verbose)
}
