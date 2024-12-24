package main

import (
	"fmt"
	"testing"
)

func TestCombinations(t *testing.T) {
	expected := [][]string{{"b", "c"}, {"a", "c"}, {"a", "b"}}
	got := combinations([]string{"a", "b", "c"}, 2)
	if len(got) != len(expected) {
		t.Errorf("RunCombination lenth = %d; expected %d", len(got), len(expected))
	}
	fmt.Println(got)
	if got[0][0] != expected[0][0] {
		t.Errorf("RunCombination[0][0] = %s; expected %s", got[0][0], expected[0][0])
	}
}

func TestCombinations2(t *testing.T) {
	expected := [][]string{{"b", "c"}, {"a", "c"}, {"a", "b"}, {"a", "d"}, {"b", "d"}, {"c", "d"}}
	got := combinations([]string{"a", "b", "c", "d"}, 2)
	if len(got) != len(expected) {
		t.Errorf("RunCombination lenth = %d; expected %d", len(got), len(expected))
	}
	fmt.Println(got)
}

func TestCombinations3(t *testing.T) {
	expected := [][]string{{"a", "b", "c", "d"}}
	got := combinations([]string{"a", "b", "c", "d"}, 4)
	if len(got) != len(expected) {
		t.Errorf("RunCombination lenth = %d; expected %d", len(got), len(expected))
	}
	fmt.Println(got)
}

func TestCombinations4(t *testing.T) {
	expected := [][]string{{"a"}, {"b"}, {"c"}, {"d"}}
	got := combinations([]string{"a", "b", "c", "d"}, 1)
	if len(got) != len(expected) {
		t.Errorf("RunCombination lenth = %d; expected %d", len(got), len(expected))
	}
	fmt.Println(got)
}

func TestCombinations5(t *testing.T) {
	expectedLen := 20
	got := combinations([]string{"a", "b", "c", "d", "e", "f"}, 3)
	if len(got) != expectedLen {
		t.Errorf("RunCombination lenth = %d; expectedLen %d", len(got), expectedLen)
	}
	fmt.Println(got)
}

func TestRunPassword(t *testing.T) {
	expected := "co,de,ka,ta"
	got := RunPassword(`kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`)
	if got != expected {
		t.Errorf("RunPassword = %s; expected %s", got, expected)
	}
}
