package common

import (
	"testing"
)

func TestParseIntList(t *testing.T) {
	var input = `1
23
45`
	var ints = ParseIntList(input, "\n")
	if len(ints) != 3 {
		t.Errorf("ParseIntList expected len 3, got %d", len(ints))
	}
	if ints[1] != 23 {
		t.Errorf("ParseIntList expected item 23, got %d", ints[1])
	}
}

func TestParseDigitList(t *testing.T) {
	var input = `12345`
	var ints = ParseDigitList(input)
	if len(ints) != 5 {
		t.Errorf("ParseDigitList expected len 5, got %d", len(ints))
	}
	if ints[1] != 2 {
		t.Errorf("ParseDigitList expected item 23, got %d", ints[1])
	}
}

func TestParseIntCsvList(t *testing.T) {
	var input = `1,23,45`
	var ints = ParseIntList(input, ",")
	if len(ints) != 3 {
		t.Errorf("ParseIntList expected len 3, got %d", len(ints))
	}
	if ints[1] != 23 {
		t.Errorf("ParseIntList expected item 23, got %d", ints[1])
	}
}

func TestParseStringList(t *testing.T) {
	var input = `hej
på
dig`
	var words = ParseStringList(input, "\n")
	if len(words) != 3 {
		t.Errorf("ParseIntList expected len 3, got %d", len(words))
	}
	if words[1] != "på" {
		t.Errorf("ParseIntList expected item <på>, got %s", words[1])
	}
}

func TestParseStringList_sentense(t *testing.T) {
	var input = `hej på dig`
	var words = ParseStringList(input, " ")
	if len(words) != 3 {
		t.Errorf("ParseIntList expected len 3, got %d", len(words))
	}
	if words[1] != "på" {
		t.Errorf("ParseIntList expected item <på>, got %s", words[1])
	}
}

func TestParseRuneListList(t *testing.T) {
	var input = `abc
def
ghi`
	var runes = ParseRuneListList(input)
	if len(runes) != 3 {
		t.Errorf("ParseRuneListList expected len 3, got %d", len(runes))
	}
	if len(runes[0]) != 3 {
		t.Errorf("ParseRuneListList expected len 3, got %d", len(runes[0]))
	}
	if runes[2][1] != 'h' {
		t.Errorf("ParseRuneListList expected rune x, got %q", runes[2][1])
	}
	if runes[1][1] != 'e' {
		t.Errorf("ParseRuneListList expected rune e, got %q", runes[1][1])
	}
}

func TestParseDigitListList(t *testing.T) {
	var input = `123
345
567`
	var digits = ParseDigitListList(input)
	if len(digits) != 3 {
		t.Errorf("ParseDigitListList expected len 3, got %d", len(digits))
	}
	if len(digits[0]) != 3 {
		t.Errorf("ParseDigitListList expected len 3, got %d", len(digits[0]))
	}
	if digits[2][1] != 6 {
		t.Errorf("ParseDigitListList expected int 6, got %q", digits[2][1])
	}
	if digits[1][1] != 4 {
		t.Errorf("ParseDigitListList expected int 4, got %q", digits[1][1])
	}
}

func TestParseRuneListListList(t *testing.T) {
	var input = `abc
def
ghi

123
456
789`
	var runes = ParseRuneListListList(input)
	if len(runes) != 2 {
		t.Errorf("ParseRuneListListList expected len 2, got %d", len(runes))
	}
	if len(runes[0]) != 3 {
		t.Errorf("ParseRuneListListList expected len 3, got %d", len(runes[0]))
	}
	if len(runes[0][1]) != 3 {
		t.Errorf("ParseRuneListListList expected len 3, got %d", len(runes[0][1]))
	}
	if runes[0][2][1] != 'h' {
		t.Errorf("ParseRuneListListList expected rune h, got %q", runes[0][2][1])
	}
	if runes[1][1][0] != '4' {
		t.Errorf("ParseRuneListList expected rune 4, got %q", runes[1][1][0])
	}
}

func TestParseIntListList(t *testing.T) {
	var input = `2   5
19   21
123   456`
	var ints = ParseIntListList(input, "   ")
	if len(ints) != 3 {
		t.Errorf("ParseIntListList expected len 2, got %d", len(ints))
	}
	if len(ints[0]) != 2 {
		t.Errorf("ParseIntListList expected len 2, got %d", len(ints[0]))
	}
	if ints[0][0] != 2 {
		t.Errorf("ParseIntListList expected int 3, got %d", ints[0][0])
	}
	if ints[2][0] != 123 {
		t.Errorf("ParseIntListList expected int 123, got %q", ints[2][0])
	}
	if ints[1][1] != 21 {
		t.Errorf("ParseIntListList expected int 21, got %q", ints[1][1])
	}
}
