package gohangul

import (
	"testing"
	"time"
)

func BenchmarkDisassemble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Disassemble("안녕하세요")
	}
}

func BenchmarkAssemble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Disassemble("ㅇㅏㄴㅕㅇㅎㅏㅅㅔㅇㅛ")
	}
}

func BenchmarkRomanize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Romanize("안녕하세요")
	}
}

func TestAssemble(t *testing.T) {
	input := "ㅇㅏㄴㄴㅕㅇㅎㅏㅅㅓㅣㅇㅛ. ㅂㅗㄱㅎㅏㅂ ㅈㅜㅇㅅㅓㅇ (ㅇㅡㅣㅅㅏ, ㅇㅗㅣㄱㅗㅏ) ㅂㅗㄱㅎㅏㅂ ㅈㅗㅇㅅㅓㅇ (ㅅㅏㄹㅁ, ㄷㅏㄹㄱ)"
	want := "안녕하세요. 복합 중성 (의사, 외과) 복합 종성 (삶, 닭)"
	output := Assemble(input)

	if output != want {
		t.Errorf("Assemble(%q) = %q; want %q", input, output, want)
	}
}

func TestDisassemble(t *testing.T) {
	input := "안녕하세요. 복합 중성 (의사, 외과) 복합 종성 (삶, 닭)"
	want := "ㅇㅏㄴㄴㅕㅇㅎㅏㅅㅓㅣㅇㅛ. ㅂㅗㄱㅎㅏㅂ ㅈㅜㅇㅅㅓㅇ (ㅇㅡㅣㅅㅏ, ㅇㅗㅣㄱㅗㅏ) ㅂㅗㄱㅎㅏㅂ ㅈㅗㅇㅅㅓㅇ (ㅅㅏㄹㅁ, ㄷㅏㄹㄱ)"
	output := Disassemble(input).String()

	if output != want {
		t.Errorf("Disassemble(%q) = %q; want %q", input, output, want)
	}
}

func TestJosaPick(t *testing.T) {
	input := []string{"사과", "귤", "바나나", "멜론", "딸기"}
	want := []string{"는", "은", "는", "은", "는"}

	for i, v := range input {
		output := JosaPick(v, "은/는")
		if output != want[i] {
			t.Errorf("JosaPick(%q) = %q; want %q", v, output, want[i])
		}
	}
}

func TestJosa(t *testing.T) {
	input := []string{"사과", "귤", "바나나", "멜론", "딸기"}
	want := []string{"사과는", "귤은", "바나나는", "멜론은", "딸기는"}

	for i, v := range input {
		output := Josa(v, "은/는")
		if output != want[i] {
			t.Errorf("Josa(%q) = %q; want %q", v, output, want[i])
		}
	}
}

func TestCanBeChoseong(t *testing.T) {
	input := []string{"ㄱ", "ㅎ", "ㅃ", "ㄱㄱ", "ㅘ", "ㅜ"}
	want := []bool{true, true, true, false, false, false}

	for i, v := range input {
		output := CanBeChoseong(v)
		if output != want[i] {
			t.Errorf("CanBeChoseong(%q) = %t; want %t", v, output, want[i])
		}
	}
}

func TestCanBeJungseong(t *testing.T) {
	input := []string{"ㄱ", "ㅃ", "ㄱㄱ", "ㅗ", "ㅘ", "ㅡㅣ"}
	want := []bool{false, false, false, true, true, true}

	for i, v := range input {
		output := CanBeJungseong(v)
		if output != want[i] {
			t.Errorf("CanBeJungseong(%q) = %t; want %t", v, output, want[i])
		}
	}
}

func TestCanBeJongseong(t *testing.T) {
	input := []string{"ㄱ", "ㅎ", "ㅃ", "ㄱㄱ", "ㅅㅅ", "ㅘ", "ㅜ"}
	want := []bool{true, true, false, true, true, false, false}

	for i, v := range input {
		output := CanBeJongseong(v)
		if output != want[i] {
			t.Errorf("CanBeJongseong(%q) = %t; want %t", v, output, want[i])
		}
	}
}

func TestCombineCharacter(t *testing.T) {
	input := [][3]string{{"ㄱ", "ㅏ", "ㅆ"}, {"ㅇ", "ㅡㅣ"}, {"ㅃ", "ㅜㅓ", "ㄹㄱ"}}
	want := []string{"갔", "의", "뿱"}

	for i, v := range input {
		output := CombineCharacter(v[0], v[1], v[2])
		if output != want[i] {
			t.Errorf("CombineCharacter(%q) = %q; want %q", v, output, want[i])
		}
	}
}

func TestCombineVowels(t *testing.T) {
	input := [][]string{{"ㅡ", "ㅣ"}, {"ㅜ", "ㅓ"}, {"ㅗ", "ㅏ"}}
	want := []string{"ㅢ", "ㅝ", "ㅘ"}

	for i, v := range input {
		output := CombineVowels(v[0], v[1])
		if output != want[i] {
			t.Errorf("CombineVowels(%q) = %q; want %q", v, output, want[i])
		}
	}
}

func TestDays(t *testing.T) {
	input := []int{14, 2, 29}
	want := []string{"열나흘", "이틀", "스무아흐레"}

	for i, v := range input {
		output := Days(v)
		if output != want[i] {
			t.Errorf("Days(%d) = %q; want %q", v, output, want[i])
		}
	}
}

func TestGetChoseong(t *testing.T) {
	input := []string{"라면", "안녕하세요", "뽀로로", "까마귀"}
	want := []string{"ㄹㅁ", "ㅇㄴㅎㅅㅇ", "ㅃㄹㄹ", "ㄲㅁㄱ"}

	for i, v := range input {
		output := GetChoseong(v)
		if output != want[i] {
			t.Errorf("GetChoseong(%q) = %q; want %q", v, output, want[i])
		}
	}
}

func TestHasBatchim(t *testing.T) {
	input := []string{"감", "갃", "강", "가"}
	want := []bool{true, true, true, false}

	for i, v := range input {
		output := HasBatchim(v)
		if output != want[i] {
			t.Errorf("HasBatchim(%q) = %t; want %t", v, output, want[i])
		}
	}
}

func TestNumberToHangul(t *testing.T) {
	input := []string{"0123456", "7890", "1000000", "1000000.1023", "102001030"}
	want := []string{"일십이만삼천사백오십육", "칠천팔백구십", "일백만", "일백만점일영이삼", "일억이백만일천삼십"}

	for i, v := range input {
		output := NumberToHangul(v)
		if output != want[i] {
			t.Errorf("NumberToHangul(%s) = %q; want %q", v, output, want[i])
		}
	}
}

func TestRomanize(t *testing.T) {
	input := []string{"안녕하세요", "반갑습니다", "한글로", "로마자로"}
	want := []string{"annyeonghaseyo", "bangapseupnida", "hangeulro", "romajaro"}

	for i, v := range input {
		output := Romanize(v)
		if output != want[i] {
			t.Errorf("Romanize(%q) = %q; want %q", v, output, want[i])
		}
	}
}

func TestWeekday(t *testing.T) {
	input := []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}
	want := []string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"}

	for i, v := range input {
		output := Weekday(v, true)
		if output != want[i] {
			t.Errorf("Weekday(%q) = %q; want %q", v, output, want[i])
		}
	}
}
