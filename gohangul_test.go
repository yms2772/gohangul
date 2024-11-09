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

func BenchmarkCanBeChoseong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CanBeChoseong("ㅇ")
	}
}

func BenchmarkCanBeJungseong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CanBeJungseong("ㅡ")
	}
}

func BenchmarkCanBeJongseong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CanBeJongseong("ㅇ")
	}
}

func BenchmarkCombineCharacter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CombineCharacter("ㄱ", "ㅏ")
	}
}

func BenchmarkCombineVowels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CombineVowels("ㅏ", "ㅓ")
	}
}

func BenchmarkDays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Days(14)
	}
}

func BenchmarkGetChoseong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetChoseong("라면")
	}
}

func BenchmarkHasBatchim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HasBatchim("갂")
	}
}

func BenchmarkJosa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Josa("사과", "이/가")
	}
}

func BenchmarkJosaPick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JosaPick("사과", "이/가")
	}
}

func BenchmarkNumberToHangul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberToHangul("1234567890")
	}
}

func BenchmarkWeekday(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Weekday(time.Sunday)
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
	tests := []struct {
		word     string
		josaType string
		expected string
	}{
		// "사과"는 종성이 없는 단어
		{"사과", "이/가", "가"},
		{"사과", "을/를", "를"},
		{"사과", "은/는", "는"},
		{"사과", "으로/로", "로"},
		{"사과", "와/과", "와"},
		{"사과", "이나/나", "나"},
		{"사과", "이란/란", "란"},
		{"사과", "아/야", "야"},
		{"사과", "이랑/랑", "랑"},
		{"사과", "이에요/예요", "예요"},
		{"사과", "으로서/로서", "로서"},
		{"사과", "으로써/로써", "로써"},
		{"사과", "으로부터/로부터", "로부터"},
		{"사과", "이라/라", "라"},
		{"사과", "?/?", "?/?"},

		// "귤"은 종성이 있는 단어
		{"귤", "이/가", "이"},
		{"귤", "을/를", "을"},
		{"귤", "은/는", "은"},
		{"귤", "으로/로", "으로"},
		{"귤", "와/과", "과"},
		{"귤", "이나/나", "이나"},
		{"귤", "이란/란", "이란"},
		{"귤", "아/야", "아"},
		{"귤", "이랑/랑", "이랑"},
		{"귤", "이에요/예요", "이에요"},
		{"귤", "으로서/로서", "으로서"},
		{"귤", "으로써/로써", "으로써"},
		{"귤", "으로부터/로부터", "으로부터"},
		{"귤", "이라/라", "이라"},
		{"귤", "?/?", "?/?"},
	}

	for _, test := range tests {
		result := JosaPick(test.word, test.josaType)
		if result != test.expected {
			t.Errorf("JosaPick(%q, %q) = %q; want %q", test.word, test.josaType, result, test.expected)
		}
	}
}

func TestJosa(t *testing.T) {
	tests := []struct {
		word     string
		josaType string
		expected string
	}{
		{"사과", "이/가", "사과가"},
		{"귤", "이/가", "귤이"},
		{"사과", "을/를", "사과를"},
		{"귤", "을/를", "귤을"},
	}

	for _, test := range tests {
		result := Josa(test.word, test.josaType)
		if result != test.expected {
			t.Errorf("Josa(%q, %q) = %q; want %q", test.word, test.josaType, result, test.expected)
		}
	}
}

func TestCanBeChoseong(t *testing.T) {
	input := []string{"", "ㄱ", "ㅎ", "ㅃ", "ㄱㄱ", "ㅘ", "ㅜ"}
	want := []bool{false, true, true, true, false, false, false}

	for i, v := range input {
		output := CanBeChoseong(v)
		if output != want[i] {
			t.Errorf("CanBeChoseong(%q) = %t; want %t", v, output, want[i])
		}
	}
}

func TestCanBeJungseong(t *testing.T) {
	input := []string{"", "ㄱ", "ㅃ", "ㄱㄱ", "ㅗ", "ㅘ", "ㅡㅣ", "ㅡㅣㅑ"}
	want := []bool{false, false, false, false, true, true, true, false}

	for i, v := range input {
		output := CanBeJungseong(v)
		if output != want[i] {
			t.Errorf("CanBeJungseong(%q) = %t; want %t", v, output, want[i])
		}
	}
}

func TestCanBeJongseong(t *testing.T) {
	input := []string{"", "ㄱ", "ㅎ", "ㅃ", "ㄱㄱ", "ㅅㅅ", "ㅘ", "ㅜ", "ㅂㅂㅂ"}
	want := []bool{false, true, true, false, true, true, false, false, false}

	for i, v := range input {
		output := CanBeJongseong(v)
		if output != want[i] {
			t.Errorf("CanBeJongseong(%q) = %t; want %t", v, output, want[i])
		}
	}
}

func TestCombineCharacter(t *testing.T) {
	input := [][3]string{{"ㅇ", "ㅡㅣ"}, {"ㄱ", "ㅏ"}}
	want := []string{"의", "가"}

	for i, v := range input {
		output := CombineCharacter(v[0], v[1])
		if output != want[i] {
			t.Errorf("CombineCharacter(%q) = %q; want %q", v, output, want[i])
		}
	}

	input = [][3]string{{"ㄱ", "ㅏ", "ㅆ"}, {"ㅃ", "ㅜㅓ", "ㄹㄱ"}}
	want = []string{"갔", "뿱"}

	for i, v := range input {
		output := CombineCharacter(v[0], v[1], v[2])
		if output != want[i] {
			t.Errorf("CombineCharacter(%q) = %q; want %q", v, output, want[i])
		}
	}
}

func TestCombineVowels(t *testing.T) {
	input := [][]string{{"ㅏ", "ㅓ"}, {"ㅡ", "ㅣ"}, {"ㅜ", "ㅓ"}, {"ㅗ", "ㅏ"}}
	want := []string{"ㅏㅓ", "ㅢ", "ㅝ", "ㅘ"}

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
	input := []string{".", "감", "갃", "강", "가"}
	want := []bool{false, true, true, true, false}

	for i, v := range input {
		output := HasBatchim(v)
		if output != want[i] {
			t.Errorf("HasBatchim(%q) = %t; want %t", v, output, want[i])
		}
	}

	input = []string{"갂", "가", "각", "갔"}
	want = []bool{true, false, false, true}

	for i, v := range input {
		output := HasBatchim(v, true)
		if output != want[i] {
			t.Errorf("HasBatchim(%q) = %t; want %t", v, output, want[i])
		}
	}

	input = []string{"갂", "가", "각", "갔"}
	want = []bool{false, false, true, false}

	for i, v := range input {
		output := HasBatchim(v, false)
		if output != want[i] {
			t.Errorf("HasBatchim(%q) = %t; want %t", v, output, want[i])
		}
	}
}

func TestNumberToHangul(t *testing.T) {
	input := []string{"", "12", "0123456", "7890", "1000000", "1000000.1023", "102001030", "100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}
	want := []string{"", "일십이", "일십이만삼천사백오십육", "칠천팔백구십", "일백만", "일백만점일영이삼", "일억이백만일천삼십", ""}

	for i, v := range input {
		output := NumberToHangul(v)
		if output != want[i] {
			t.Errorf("NumberToHangul(%s) = %q; want %q", v, output, want[i])
		}
	}
}

func TestRomanize(t *testing.T) {
	input := []string{"", "안녕하세요", "반갑습니다", "한글로", "로마자로"}
	want := []string{"", "annyeonghaseyo", "bangapseupnida", "hangeulro", "romajaro"}

	for i, v := range input {
		output := Romanize(v)
		if output != want[i] {
			t.Errorf("Romanize(%q) = %q; want %q", v, output, want[i])
		}
	}
}

func TestWeekday(t *testing.T) {
	input := []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}
	want := []string{"일", "월", "화", "수", "목", "금", "토"}

	for i, v := range input {
		output := Weekday(v)
		if output != want[i] {
			t.Errorf("Weekday(%q) = %q; want %q", v, output, want[i])
		}
	}

	want = []string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"}

	for i, v := range input {
		output := Weekday(v, true)
		if output != want[i] {
			t.Errorf("Weekday(%q) = %q; want %q", v, output, want[i])
		}
	}
}
