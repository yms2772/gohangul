package gohangul

// Jamo 한글 자모
type Jamo rune

// Empty 자모가 비어있는지 확인합니다.
func (j Jamo) Empty() bool {
	return j == 0
}

// Equals 자모가 같은지 확인합니다.
func (j Jamo) Equals(target Jamo) bool {
	return j == target ||
		j.toLetter() == target.toLetter() ||
		j.toChoseong() == target.toChoseong() ||
		j.toJongseong() == target.toJongseong()
}

// String 자모를 문자열로 변환합니다.
func (j Jamo) String() string {
	return string(j)
}

// IsHangul 한글인지 확인합니다.
func (j Jamo) IsHangul() bool {
	return (j >= baseChoseong && j <= baseJongseong+numJongseong) ||
		(j >= baseJungseong && j <= baseJungseong+numJungseong) ||
		(j >= baseJongseong && j <= baseJongseong+numJongseong)
}

// toLetter 자모를 한글 문자로 변환합니다.
func (j Jamo) toLetter() Jamo {
	if v, ok := toLetterMap[j]; ok {
		return v
	}
	return j
}

// toChoseong 한글 문자를 자모 초성으로 변환합니다.
func (j Jamo) toChoseong() Jamo {
	if v, ok := toChoseongMap[j]; ok {
		return v
	}
	return j
}

// toJongseong 중성으로 변환합니다.
func (j Jamo) toJongseong() Jamo {
	if v, ok := toJongseongMap[j]; ok {
		return v
	}
	return j
}

// complexJungseongToChoseong 복합 중성을 초성으로 변환합니다.
func (j Jamo) complexJungseongToChoseong() string {
	if v, ok := complexJungseongReversedMap[j]; ok {
		return v
	}
	return j.toLetter().String()
}

// complexJongseongToChoseong 복합 종성을 초성으로 변환합니다.
func (j Jamo) complexJongseongToChoseong() string {
	if v, ok := complexJongseongReversedMap[j]; ok {
		return v
	}
	return j.toLetter().String()
}
