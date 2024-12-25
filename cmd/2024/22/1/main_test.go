package main

import (
	"testing"
)

func TestXor(t *testing.T) {
	expected := 37
	got := 42 ^ 15
	if got != expected {
		t.Errorf("Xor = %d; expected %d", got, expected)
	}
}

func TestSecret(t *testing.T) {
	expected := 15887950
	got := secret(123)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}

func TestRun(t *testing.T) {
	expected := 37327623
	got := Run(`1
10
100
2024`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}
