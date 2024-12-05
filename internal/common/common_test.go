package common

import (
	"testing"
)

func TestMustAtoi(t *testing.T) {
	i := MustAtoi("47")
	if i != 47 {
		t.Errorf("MustAtoi expected 47, got %d", i)
	}
}

func TestReverseString(t *testing.T) {
	s := ReverseString("odoM")
	if s != "Modo" {
		t.Errorf("ReverseString expected Modo, got %s", s)
	}
}