package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 0
	got := Run(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}

func TestRunBst(t *testing.T) {
	expected := 1
	Run(`Register A: 0
Register B: 0
Register C: 9

Program: 2,6`)
	if register[1] != expected {
		t.Errorf("Run = %d; expected %d", register[1], expected)
	}
}
