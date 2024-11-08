package gohangul

import "testing"

func TestJamo_Empty(t *testing.T) {
	input := Jamo(0)

	if got := input.Empty(); got != true {
		t.Errorf("Jamo.Empty() = %v, want %v", got, true)
	}
}

func TestJamo_Equals(t *testing.T) {
	input := Jamo('안')
	want := Jamo('안')

	if got := input.Equals(want); got != true {
		t.Errorf("Jamo.Equals() = %v, want %v", got, true)
	}
}

func TestJamo_String(t *testing.T) {
	input := Jamo('안')

	if got := input.String(); got != "안" {
		t.Errorf("Jamo.String() = %v, want %v", got, "안")
	}
}
