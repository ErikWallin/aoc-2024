package common

import (
	"strconv"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}


func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func ReverseSlice[T any](in []T) []T {
	out := make([]T, len(in))
	copy(out, in)
	for i := len(out)/2 - 1; i >= 0; i-- {
		opp := len(out) - 1 - i
		out[i], out[opp] = out[opp], out[i]
	}
	return out
}

func ReverseString(s string) string {
    n := len(s)
    runes := make([]rune, n)
    for _, rune := range s {
        n--
        runes[n] = rune
    }
    return string(runes[n:])
}
