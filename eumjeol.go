package gohangul

// Eumjeol 초성, 중성, 종성으로 이루어진 음절
type Eumjeol struct {
	Choseong  Jamo
	Jungseong Jamo
	Jongseong Jamo
}

// Empty 음절이 비어있는지 확인합니다.
func (e Eumjeol) Empty() bool {
	return e.Choseong.Empty() &&
		e.Jungseong.Empty() &&
		e.Jongseong.Empty()
}

// Equals 음절이 같은지 확인합니다.
func (e Eumjeol) Equals(target Eumjeol) bool {
	return e.Choseong.Equals(target.Choseong) &&
		e.Jungseong.Equals(target.Jungseong) &&
		e.Jongseong.Equals(target.Jongseong)
}

// String 음절을 합쳐서 한글 문자로 반환합니다.
func (e Eumjeol) String() string {
	if e.Empty() {
		return ""
	}
	if !e.Choseong.Empty() && e.Jungseong.Empty() && e.Jongseong.Empty() {
		return e.Choseong.toLetter().String()
	}

	result := Jamo(baseHangul)
	if !e.Choseong.Empty() {
		result += (e.Choseong.toChoseong() - baseChoseong) * numJungseong * numJongseong
	}
	if !e.Jungseong.Empty() {
		result += (e.Jungseong.toChoseong() - baseJungseong) * numJongseong
	}
	if !e.Jongseong.Empty() {
		result += e.Jongseong.toChoseong().toJongseong() - baseJongseong
	}
	return result.String()
}

// isHangul 한글인지 확인합니다.
func (e Eumjeol) isHangul() bool {
	return e.Choseong.IsHangul() || e.Jungseong.IsHangul() || e.Jongseong.IsHangul()
}
