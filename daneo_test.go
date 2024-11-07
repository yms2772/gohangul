package gohangul

import "testing"

func TestDaneo_Assemble(t *testing.T) {
	input := Disassemble("안녕하세요")
	want := "안녕하세요"

	if got := input.Assemble(); got != want {
		t.Errorf("Daneo.Assemble() = %v, want %v", got, want)
	}
}

func TestDaneo_At(t *testing.T) {
	input := Disassemble("안녕")
	want := []Eumjeol{
		{Choseong: Jamo('ㅇ'), Jungseong: Jamo('ㅏ'), Jongseong: Jamo('ㄴ')},
		{Choseong: Jamo('ㄴ'), Jungseong: Jamo('ㅕ'), Jongseong: Jamo('ㅇ')},
	}

	for i, w := range want {
		if got := input.At(i); !got.Equals(w) {
			t.Errorf("Daneo.At() = %v, want %v", got, w)
		}
	}

	if got := input.At(3); !got.Choseong.Empty() || !got.Jungseong.Empty() || !got.Jongseong.Empty() {
		t.Errorf("Daneo.At() = %v, want empty", got)
	}
}

func TestDaneo_Each(t *testing.T) {
	input := Disassemble("안녕")
	want := []Eumjeol{
		{Choseong: Jamo('ㅇ'), Jungseong: Jamo('ㅏ'), Jongseong: Jamo('ㄴ')},
		{Choseong: Jamo('ㄴ'), Jungseong: Jamo('ㅕ'), Jongseong: Jamo('ㅇ')},
	}

	input.Each(func(i int, eumjeol Eumjeol) {
		if !eumjeol.Equals(want[i]) {
			t.Errorf("Daneo.Each() = %v, want %v", eumjeol, want[i])
		}
	})
}

func TestDaneo_Equals(t *testing.T) {
	input := []Daneo{
		Disassemble("안녕"),
		Disassemble("안냥"),
		Disassemble("안녕하세요"),
	}
	target := Disassemble("안녕")
	want := []bool{true, false, false}

	for i, item := range input {
		if got := item.Equals(target); got != want[i] {
			t.Errorf("Daneo.Equals() = %v, want %v", got, want[i])
		}
	}
}

func TestDaneo_GetChoseong(t *testing.T) {
	input := Disassemble("안녕")
	want := "ㅇㄴ"

	if got := input.GetChoseong(); got != want {
		t.Errorf("Daneo.GetChoseong() = %v, want %v", got, want)
	}
}

func TestDaneo_String(t *testing.T) {
	input := Disassemble("안녕")
	want := "ㅇㅏㄴㄴㅕㅇ"

	if got := input.String(); got != want {
		t.Errorf("Daneo.String() = %v, want %v", got, want)
	}
}
