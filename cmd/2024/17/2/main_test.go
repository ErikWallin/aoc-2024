package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 117440
	got := Run(`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
