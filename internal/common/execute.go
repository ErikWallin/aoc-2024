package common

import (
	"fmt"
	"time"
)

func Run(run func(string) int, year int, day int, part int, submit bool, verbose bool) int {
	input := ReadFile(year, day, "")
	start := time.Now()
	answer := run(input)
	elapsed := time.Since(start)
	fmt.Printf("PART%v: %v\n", part, answer)
	fmt.Printf("Program took %s\n", elapsed)
	if submit && answer != 0 {
		Submit(year, day, part, answer)
	}
	return answer
}
