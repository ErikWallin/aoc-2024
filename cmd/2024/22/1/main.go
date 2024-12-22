package main

import (
	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	secrets := common.ParseIntList(input, "\n")
	for range 2000 {
		for i, s := range secrets {
			secrets[i] = secret(s)
		}
	}
	sum := 0
	for _, s := range secrets {
		sum += s
	}
	return sum
}

func secret(previous int) int {
	secret := previous
	secret = (secret * 64) ^ secret
	secret = secret % 16777216
	secret = (secret / 32) ^ secret
	secret = secret % 16777216
	secret = (secret * 2048) ^ secret
	secret = secret % 16777216
	return secret
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 22, 1, submit, verbose)
}
