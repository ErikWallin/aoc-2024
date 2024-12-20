package main

import (
	"slices"

	"github.com/ErikWallin/aoc-2024/internal/common"
	. "github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

var dir []C = []C{Up, Right, Down, Left}

func Run100(input string) int {
	return Run(input, 100)
}

func Run(input string, save int) int {
	f := common.ParseRuneListList(input)

	var start C
	for y, row := range f {
		for x, r := range row {
			if r == 'S' {
				start = C{x, y}
			}
		}
	}

	startScore := score{start, C{0, 0}, 0}
	current := []score{startScore}
	trackWithoutCheat := []C{startScore.c}
	t := 0
	var finalScores []int
	var timeWithoutCheat int = -1
	for len(current) > 0 {
		var next []score
		t++
		for _, cs := range current {
			for _, d := range dir {
				np := cs.c.Plus(d)
				if cs.p == np || slices.Contains(trackWithoutCheat, np) {
					// Cannot go back on single track
					continue
				}
				if np.X >= 0 && np.Y >= 0 && np.X < len(f[0]) && np.Y < len(f) {
					if f[np.Y][np.X] == '#' {
						if cs.cheat < 1 {
							next = append(next, score{np, cs.c, cs.cheat+1})
						}
					} else if f[np.Y][np.X] == 'E' {
						if cs.cheat == 0 && timeWithoutCheat == -1 {
							timeWithoutCheat = t
						}
						finalScores = append(finalScores, t)
					} else {
						next = append(next, score{np, cs.c, cs.cheat})
					}
				}
			}
		}
		if timeWithoutCheat != -1 {
			count := 0
			for _, fs := range finalScores {
				if fs <= timeWithoutCheat - save {
					count++
				}
			}
			return count
		}
		current = next
	}
	return 0
}

type score struct {
	c C
	p C
	cheat int
}

func main() {
	submit := false
	verbose := true
	common.Run(Run100, 2024, 20, 1, submit, verbose)
}
