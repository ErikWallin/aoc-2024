package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 55312
	got := Run(`125 17`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}

func TestBlinkNTimes1(t *testing.T) {
	expected := []int{253000}
	got := BlinkNTimes(125, 1)
	if len(got) != len(expected) {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
	for i := range len(got) {
		if got[i] != expected[i] {
			t.Errorf("Run = %d; expected %d", got[i], expected[i])
		}
	}
}

func TestBlinkNTimes2(t *testing.T) {
	expected := []int{253, 0}
	got := BlinkNTimes(253000, 1)
	if len(got) != len(expected) {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
	for i := range len(got) {
		if got[i] != expected[i] {
			t.Errorf("Run = %d; expected %d", got[i], expected[i])
		}
	}
}
