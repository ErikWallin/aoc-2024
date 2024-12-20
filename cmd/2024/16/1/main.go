package main

import (
	"fmt"

	"github.com/ErikWallin/aoc-2024/internal/common"
	. "github.com/ErikWallin/aoc-2024/internal/common/c2d"
)

type Score struct {
	value     int
	position  C
	direction C
}

func (s Score) key() string {
	return key(s.position, s.direction)
}

func key(position, direction C) string {
	return fmt.Sprintf("%d,%d+%d,%d", position.X, position.Y, direction.X, direction.Y)
}

func Run(input string) int {
	runes := common.ParseRuneListList(input)
	var start, end C
	for y, row := range runes {
		for x, r := range row {
			if r == 'S' {
				start = C{x, y}
			} else if r == 'E' {
				end = C{x, y}
			}
		}
	}
	records := map[string]Score{}
	startScore := Score{0, start, Right}
	currentRound := []Score{startScore}
	records[startScore.key()] = startScore
	for len(currentRound) > 0 {
		var nextRound []Score
		for _, current := range currentRound {
			nextPosition := current.position.Plus(current.direction)
			if runes[nextPosition.Y][nextPosition.X] != '#' {
				next := Score{current.value + 1, nextPosition, current.direction}
				previousScore, ok := records[next.key()]
				if !ok || next.value < previousScore.value {
					records[next.key()] = next
					nextRound = append(nextRound, next)
				}
			}
			if current.direction == Up || current.direction == Down {
				next := Score{current.value + 1000, current.position, Left}
				previousScore, ok := records[next.key()]
				if !ok || next.value < previousScore.value {
					records[next.key()] = next
					nextRound = append(nextRound, next)
				}
				next = Score{current.value + 1000, current.position, Right}
				previousScore, ok = records[next.key()]
				if !ok || next.value < previousScore.value {
					records[next.key()] = next
					nextRound = append(nextRound, next)
				}
			}
			if current.direction == Left || current.direction == Right {
				next := Score{current.value + 1000, current.position, Up}
				previousScore, ok := records[next.key()]
				if !ok || next.value < previousScore.value {
					records[next.key()] = next
					nextRound = append(nextRound, next)
				}
				next = Score{current.value + 1000, current.position, Down}
				previousScore, ok = records[next.key()]
				if !ok || next.value < previousScore.value {
					records[next.key()] = next
					nextRound = append(nextRound, next)
				}
			}
		}
		currentRound = nextRound
	}
	possibleEndKeys := []string{key(end, Up), key(end, Down), key(end, Right), key(end, Left)}
	minScore := -1
	for _, k := range possibleEndKeys {
		if val, ok := records[k]; ok {
			if minScore == -1 || val.value < minScore {
				minScore = val.value
			}
		}
	}
	return minScore
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 16, 1, submit, verbose)
}
