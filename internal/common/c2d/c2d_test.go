package c2d

import (
	"testing"
)

func TestRotateRight(t *testing.T) {
	actual := C{4, 2}.RotateRight()
	expected := C{2, -4}
	if actual != expected {
		t.Errorf("RotateRight expected %d, got %d", expected, actual)
	}
	actual = expected.RotateRight().RotateRight().RotateRight().RotateRight()
	if actual != expected {
		t.Errorf("RotateRight expected %d, got %d", expected, actual)
	}
}

func TestRotateLeft(t *testing.T) {
	actual := C{2, -4}.RotateLeft()
	expected := C{4, 2}
	if actual != expected {
		t.Errorf("RotateLeft expected %d, got %d", expected, actual)
	}
	actual = expected.RotateLeft().RotateLeft().RotateLeft().RotateLeft()
	if actual != expected {
		t.Errorf("RotateLeft expected %d, got %d", expected, actual)
	}
}

func TestRotate45Right(t *testing.T) {
	actual := C{-2, 2}.Rotate45Right()
	expected := C{0, 2}
	if actual != expected {
		t.Errorf("Rotate45Right expected %d, got %d", expected, actual)
	}
	actual = expected.Rotate45Right().Rotate45Right().Rotate45Right().Rotate45Right().Rotate45Right().Rotate45Right().Rotate45Right().Rotate45Right()
	if actual != expected {
		t.Errorf("Rotate45Right expected %d, got %d", expected, actual)
	}
}

func TestRotate45Left(t *testing.T) {
	actual := C{-2, 2}.Rotate45Left()
	expected := C{-2, 0}
	if actual != expected {
		t.Errorf("Rotate45Left expected %d, got %d", expected, actual)
	}
	actual = expected.Rotate45Left().Rotate45Left().Rotate45Left().Rotate45Left().Rotate45Left().Rotate45Left().Rotate45Left().Rotate45Left()
	if actual != expected {
		t.Errorf("Rotate45Left expected %d, got %d", expected, actual)
	}
}
