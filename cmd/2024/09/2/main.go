package main

import (
	"fmt"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	ints := common.ParseDigitList(input)
	drive := []int{}
	isFile := true
	for i, n := range ints {
		for j := 0; j < n; j++ {
			if isFile {
				drive = append(drive, i / 2)
			} else {
				drive = append(drive, -1)
			}
		}
		isFile = !isFile
	}
	empties := []int{}
	for i, n := range drive {
		if n == -1 {
			empties = append(empties, i)
		}
	}
	lastPos := len(drive) - 1
	for {
		if lastPos <= 0 || drive[lastPos] == 0 {
			break
		}
		if drive[lastPos] == -1 {
			lastPos--
			continue
		}
		firstPos := lastPos
		for {
			if drive[firstPos-1] == drive[lastPos] {
				firstPos--
				continue
			}
			break
		}
		lenP := lastPos - firstPos + 1
		firstEmpty := -1
		firstEmptyI := -1
		for ei, i := range empties {
			if i + lenP > firstPos {
				break
			}
			succeded := true
			for j := range lenP {
				if i != -1 && ei + j < len(empties) && empties[ei + j] == empties[ei] + j {
					continue
				}
				succeded = false
				break
			}
			if succeded {
				firstEmpty = i
				firstEmptyI = ei
				break
			}
		}

		if firstEmpty != -1 && firstEmptyI < firstPos {
			for i := range lenP {
				drive[firstEmpty + i] = drive[firstPos + i]
				drive[firstPos + i] = -1
				empties[firstEmptyI + i] = -1
			}
			lastPos = firstPos - 1
		} else {
			lastPos = firstPos - 1
		}
	}

	sum := 0
	for i, n := range drive {
		if n == -1 {
			continue
		}
		sum += i * n
	}
	return sum
}

func print(ints []int) {
	for _, i := range ints {
		if i == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(i)
		}
	}
	fmt.Println()
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 9, 2, submit, verbose)
}
