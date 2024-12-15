package main

import (
	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	ints := common.ParseDigitList(input)
	drive := []int{}
	isFile := true
	for i, n := range ints {
		for j := 0; j < n; j++ {
			if isFile {
				drive = append(drive, i/2)
			} else {
				drive = append(drive, -1)
			}
		}
		isFile = !isFile
	}
	lastPos := len(drive)
	for i, n := range drive {
		if i >= lastPos {
			break
		}
		if n == -1 {
			for {
				lastPos--
				if i >= lastPos {
					break
				}
				if drive[lastPos] != -1 {
					drive[i] = drive[lastPos]
					drive[lastPos] = -1
					break
				}
			}
		}
	}

	sum := 0
	for i, n := range drive {
		if n == -1 {
			break
		}
		sum += i * n
	}

	return sum
}

/*func print(ints []int) {
	for _, i := range ints {
		if i == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(i)
		}
	}
	fmt.Println()
}*/

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 9, 1, submit, verbose)
}
