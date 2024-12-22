package main

import (
	"testing"
)

func TestSequenceKey(t *testing.T) {
	expected := 1226
	got := sequenceKey([]int{1, 2, 3, -4})
	if got != expected {
		t.Errorf("SequenceKey = %d; expected %d", got, expected)
	}
}

func TestSequenceKey2(t *testing.T) {
	expected := -1834
	got := sequenceKey([]int{-2, 2, -3, -4})
	if got != expected {
		t.Errorf("SequenceKey = %d; expected %d", got, expected)
	}
}

func TestSecret(t *testing.T) {
	expected := 15887950
	expectedDiff := -3
	got, diff := secret(123)
	if got != expected {
		t.Errorf("Secret = %d; expected %d", got, expected)
	}
	if diff != expectedDiff {
		t.Errorf("Secret diff = %d; expected %d", diff, expectedDiff)
	}
}

func TestRun(t *testing.T) {
	expected := 23
	got := Run(`1
2
3
2024`)
	if got != expected {
		t.Errorf("Run = %d; expected %d", got, expected)
	}
}