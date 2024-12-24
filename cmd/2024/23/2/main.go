package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ErikWallin/aoc-2024/internal/common"
)

var connections map[string][]string = map[string][]string{}

func Run(input string) int {
	fmt.Println(RunPassword(input))
	return 0
}

func RunPassword(input string) string {
	rows := common.ParseStringList(input, "\n")
	for _, row := range rows {
		parts := strings.Split(row, "-")
		first := parts[0]
		second := parts[1]
		if _, found := connections[first]; !found {
			connections[first] = []string{}
		}
		if _, found := connections[second]; !found {
			connections[second] = []string{}
		}
		connections[first] = append(connections[first], second)
		connections[second] = append(connections[second], first)
	}
	passwords := []string{}

	for k, v := range connections {
		v = append(v, k)
		for l := len(v); len(v) > 3 && l > 3; l-- {
			for _, c := range combinations(v, l) {
				slices.Sort(c)
				passwords = append(passwords, strings.Join(c, ","))
			}
		}
	}

	//fmt.Println(passwords)
	occurances := map[string]int{}
	for _, p := range passwords {
		if _, ok := occurances[p]; !ok {
			occurances[p] = 0
		}
		occurances[p]++
	}

	for _, p := range passwords {
		length := (strings.LastIndex(p, ",") + 1) / 3 + 1
		if occurances[p] == length {
			fmt.Println(p)
			return p
		}
	}
	return "Not found"
}

func combinations(nodes []string, amount int) [][]string {
	if len(nodes) == amount {
		return [][]string{nodes}
	}
	if amount == 1 {
		res := [][]string{}
		for _, n := range nodes {
			res = append(res, []string{n})
		}
		return res
	}
	res := [][]string{}
	for i, n := range nodes {
		if i >= len(nodes)-1 {
			break
		}
		rest := nodes[i+1:]
		restCombinations := combinations(rest, amount - 1)
		for _, rc := range restCombinations {
			combination := []string{n}
			combination = append(combination, rc...)
			res = append(res, combination)
		}
	}
	return res
}

func main() {
	submit := false
	verbose := true
	common.Run(Run, 2024, 23, 1, submit, verbose)
}
