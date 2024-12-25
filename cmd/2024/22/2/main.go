package main

import (
	"github.com/ErikWallin/aoc-2024/internal/common"
)

func Run(input string) int {
	secrets := common.ParseIntList(input, "\n")

	// map of sequence-key to map of secret-index to price
	var prices map[int]map[int]int = map[int]map[int]int{}

	var existingKeys map[int]bool = map[int]bool{}

	// map of secret-index to sequence
	var currentSequence map[int][]int = map[int][]int{}
	for i := range secrets {
		currentSequence[i] = []int{}
	}

	for range 2000 {
		for idx, currentSecret := range secrets {
			newSecret, diff := secret(currentSecret)
			secrets[idx] = newSecret
			currentSequence[idx] = append(currentSequence[idx], diff)
			if len(currentSequence[idx]) > 4 {
				currentSequence[idx] = currentSequence[idx][1:]
			}
			if len(currentSequence[idx]) == 4 {
				key := sequenceKey(currentSequence[idx])
				existingKeys[key] = true
				if _, ok := prices[key]; !ok {
					prices[key] = map[int]int{}
				}
				if _, ok := prices[key][idx]; !ok {
					prices[key][idx] = newSecret % 10
				}
			}
		}
	}

	max := 0
	for k := range existingKeys {
		sum := 0
		for _, p := range prices[k] {
			sum += p
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func secret(previous int) (int, int) {
	secret := previous
	secret = (secret * 64) ^ secret
	secret = secret % 16777216
	secret = (secret / 32) ^ secret
	secret = secret % 16777216
	secret = (secret * 2048) ^ secret
	secret = secret % 16777216
	return secret, secret%10 - previous%10
}

func sequenceKey(s []int) int {
	return s[0]*10000000 + s[1]*10000 + s[2]*100 + s[3]
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 22, 2, submit, verbose)
}
