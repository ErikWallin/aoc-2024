package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 4
	got := Run(`x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -> z00
x01 XOR y01 -> z01
x02 OR y02 -> z02`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}