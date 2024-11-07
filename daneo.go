package gohangul

import "strings"

type Daneo []Eumjeol

// Equals 단어가 같은지 확인합니다.
func (d Daneo) Equals(target Daneo) bool {
	if len(d) != len(target) {
		return false
	}

	for i := range d {
		if !d[i].Equals(target[i]) {
			return false
		}
	}
	return true
}

func (d Daneo) String() string {
	var sb strings.Builder

	for _, item := range d {
		if !item.Choseong.Empty() {
			sb.WriteRune(item.Choseong.toHangulLetterSios().Rune())
		}
		if !item.Jungseong.Empty() {
			v, ok := complexJungseongReversedMap[item.Jungseong]
			if ok {
				sb.WriteString(v)
			} else {
				sb.WriteRune(item.Jungseong.toHangulLetterSios().Rune())
			}
		}
		if !item.Jongseong.Empty() {
			v, ok := complexJongseongReversedMap[item.Jongseong]
			if ok {
				sb.WriteString(v)
			} else {
				sb.WriteRune(item.Jongseong.toHangulLetterSios().Rune())
			}
		}
	}
	return sb.String()
}

// Assemble 단어를 조합합니다.
func (d Daneo) Assemble() string {
	var sb strings.Builder

	for _, item := range d {
		sb.WriteString(item.String())
	}
	return sb.String()
}

// At 단어에서 i번째 음절을 가져옵니다.
func (d Daneo) At(i int) Eumjeol {
	if i < 0 || i >= len(d) {
		return Eumjeol{}
	}
	return d[i]
}

// Each 단어의 각 음절에 대해 함수를 실행합니다.
func (d Daneo) Each(f func(int, Eumjeol)) {
	for i := range d {
		f(i, d[i])
	}
}

// GetChoseong 단어에서 초성만 분리합니다.
func (d Daneo) GetChoseong() string {
	var sb strings.Builder

	for i := range d {
		sb.WriteString(d[i].Choseong.toHangulLetterSios().String())
	}
	return sb.String()
}
