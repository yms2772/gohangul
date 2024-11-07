package gohangul

type Jamo rune

// Empty 자모가 비어있는지 확인합니다.
func (j Jamo) Empty() bool {
	return j == 0
}

// Equals 자모가 같은지 확인합니다.
func (j Jamo) Equals(target Jamo) bool {
	return j == target ||
		j.toHangulLetterSios() == target.toHangulLetterSios() ||
		j.toHangulChoseongSios() == target.toHangulChoseongSios() ||
		j.toChoseong() == target.toChoseong() ||
		j.toJongseong() == target.toJongseong()
}

// String 자모를 문자열로 변환합니다.
func (j Jamo) String() string {
	if j.Empty() {
		return ""
	}
	return string(j)
}

// Rune 자모를 룬 코드로 변환합니다.
func (j Jamo) Rune() rune {
	return rune(j)
}

// isHangul 한글인지 확인합니다.
func (j Jamo) isHangul() bool {
	return (j >= baseChoseong && j <= baseJongseong+numJongseong) ||
		(j >= baseJungseong && j <= baseJungseong+numJungseong) ||
		(j >= baseJongseong && j <= baseJongseong+numJongseong)
}

// toHangulLetterSios Hangul Letter Sios로 변환합니다.
func (j Jamo) toHangulLetterSios() Jamo {
	if v, ok := choseongToLetterSiosMap[j]; ok {
		return v
	}
	return j
}

// toHangulChoseongSios Hangul Choseong Sios로 변환합니다.
func (j Jamo) toHangulChoseongSios() Jamo {
	if v, ok := letterToChoseongSiosMap[j]; ok {
		return v
	}
	return j
}

// toChoseong 초성으로 변환합니다.
func (j Jamo) toChoseong() Jamo {
	if v, ok := jongseongToChoseong[j]; ok {
		return v
	}
	return j
}

// toJongseong 중성으로 변환합니다.
func (j Jamo) toJongseong() Jamo {
	if v, ok := choseongToJongseongMap[j]; ok {
		return v
	}
	return j
}
