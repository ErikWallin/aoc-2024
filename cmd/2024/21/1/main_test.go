package main

import (
	"testing"
)

func TestNumericPath(t *testing.T) {
	expected := ">^^^"
	got := getSequence(numericPositions, '0', '9')[0]
	if got != expected {
		t.Errorf("Run = %s; expected %s", got, expected)
	}
}

func TestNumericPaths(t *testing.T) {
	expected := "<A^A>^^AvvvA"
	got := getSequenceOnNumeric("029A")[0]
	if got != expected {
		t.Errorf("Run = %s; expected %s", got, expected)
	}
}

func TestRun(t *testing.T) {
	expected := 126384
	got := Run(`029A
980A
179A
456A
379A`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
