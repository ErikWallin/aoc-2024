package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := "4,6,3,5,6,3,5,2,1,0"
	got := RunString(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`)
	if got != expected {
		t.Errorf("Run = %s; expected %s", got, expected)
	}
}
