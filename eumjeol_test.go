package gohangul

import "testing"

func TestEumjeol_Empty(t *testing.T) {
	input := Eumjeol{}

	if got := input.Empty(); got != true {
		t.Errorf("Eumjeol.Empty() = %v, want %v", got, true)
	}
}

func TestEumjeol_Equals(t *testing.T) {
	input := Eumjeol{Choseong: Jamo('ㅇ'), Jungseong: Jamo('ㅏ'), Jongseong: Jamo('ㄴ')}

	if got := input.Equals(Eumjeol{Choseong: Jamo('ㅇ'), Jungseong: Jamo('ㅏ'), Jongseong: Jamo('ㄴ')}); got != true {
		t.Errorf("Eumjeol.Equals() = %v, want %v", got, true)
	}
}

func TestEumjeol_String(t *testing.T) {
	input := []Eumjeol{
		{Choseong: Jamo('ㅇ'), Jungseong: Jamo('ㅏ'), Jongseong: Jamo('ㄴ')},
		{Choseong: Jamo('ㅇ'), Jungseong: Jamo('ㅏ'), Jongseong: Jamo('ㄲ').toChoseong()},
		{},
	}
	want := []string{"안", "앆", ""}

	for i, e := range input {
		if got := e.String(); got != want[i] {
			t.Errorf("Eumjeol.String() = %v, want %v", got, want[i])
		}
	}
}
