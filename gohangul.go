package gohangul

import (
	"strings"
	"time"
	"unicode/utf8"
)

const (
	baseHangul    = 0xAC00 // 첫 번째 한글
	lastHangul    = 0xD7A3 // 마지막 한글
	baseChoseong  = 0x1100 // 초성 시작 위치
	baseJungseong = 0x1161 // 중성 시작 위치
	baseJongseong = 0x11A7 // 종성 시작 위치
	numChoseong   = 19     // 초성 개수
	numJungseong  = 21     // 중성 개수
	numJongseong  = 28     // 종성 개수
)

var (
	// Hangul Choseong Sios -> Hangul Letter Sios
	choseongToLetterSiosMap = map[Jamo]Jamo{
		0x1100: 0x3131, // ㄱ (U+1100) -> ㄱ (U+3131)
		0x1101: 0x3132, // ㄲ (U+1101) -> ㄲ (U+3132)
		0x1102: 0x3134, // ㄴ (U+1102) -> ㄴ (U+3134)
		0x1103: 0x3137, // ㄷ (U+1103) -> ㄷ (U+3137)
		0x1104: 0x3138, // ㄸ (U+1104) -> ㄸ (U+3138)
		0x1105: 0x3139, // ㄹ (U+1105) -> ㄹ (U+3139)
		0x1106: 0x3141, // ㅁ (U+1106) -> ㅁ (U+3141)
		0x1107: 0x3142, // ㅂ (U+1107) -> ㅂ (U+3142)
		0x1108: 0x3143, // ㅃ (U+1108) -> ㅃ (U+3143)
		0x1109: 0x3145, // ㅅ (U+1109) -> ㅅ (U+3145)
		0x110A: 0x3146, // ㅆ (U+110A) -> ㅆ (U+3146)
		0x110B: 0x3147, // ㅇ (U+110B) -> ㅇ (U+3147)
		0x110C: 0x3148, // ㅈ (U+110C) -> ㅈ (U+3148)
		0x110D: 0x3149, // ㅉ (U+110D) -> ㅉ (U+3149)
		0x110E: 0x314A, // ㅊ (U+110E) -> ㅊ (U+314A)
		0x110F: 0x314B, // ㅋ (U+110F) -> ㅋ (U+314B)
		0x1110: 0x314C, // ㅌ (U+1110) -> ㅌ (U+314C)
		0x1111: 0x314D, // ㅍ (U+1111) -> ㅍ (U+314D)
		0x1112: 0x314E, // ㅎ (U+1112) -> ㅎ (U+314E)
		0x1161: 0x314F, // ㅏ (U+1161) -> ㅏ (U+314F)
		0x1162: 0x3150, // ㅐ (U+1162) -> ㅐ (U+3150)
		0x1163: 0x3151, // ㅑ (U+1163) -> ㅑ (U+3151)
		0x1164: 0x3152, // ㅒ (U+1164) -> ㅒ (U+3152)
		0x1165: 0x3153, // ㅓ (U+1165) -> ㅓ (U+3153)
		0x1166: 0x3154, // ㅔ (U+1166) -> ㅔ (U+3154)
		0x1167: 0x3155, // ㅕ (U+1167) -> ㅕ (U+3155)
		0x1168: 0x3156, // ㅖ (U+1168) -> ㅖ (U+3156)
		0x1169: 0x3157, // ㅗ (U+1169) -> ㅗ (U+3157)
		0x116A: 0x3158, // ㅘ (U+116A) -> ㅘ (U+3158)
		0x116B: 0x3159, // ㅙ (U+116B) -> ㅙ (U+3159)
		0x116C: 0x315A, // ㅚ (U+116C) -> ㅚ (U+315A)
		0x116D: 0x315B, // ㅛ (U+116D) -> ㅛ (U+315B)
		0x116E: 0x315C, // ㅜ (U+116E) -> ㅜ (U+315C)
		0x116F: 0x315D, // ㅝ (U+116F) -> ㅝ (U+315D)
		0x1170: 0x315E, // ㅞ (U+1170) -> ㅞ (U+315E)
		0x1171: 0x315F, // ㅟ (U+1171) -> ㅟ (U+315F)
		0x1172: 0x3160, // ㅠ (U+1172) -> ㅠ (U+3160)
		0x1173: 0x3161, // ㅡ (U+1173) -> ㅡ (U+3161)
		0x1174: 0x3162, // ㅢ (U+1174) -> ㅢ (U+3162)
		0x1175: 0x3163, // ㅣ (U+1175) -> ㅣ (U+3163)
		0x11A8: 0x3131, // ㄱ (U+11A8) -> ㄱ (U+3131)
		0x11A9: 0x3132, // ㄲ (U+11A9) -> ㄲ (U+3132)
		0x11AA: 0x3133, // ㄳ (U+11AA) -> ㄳ (U+3133)
		0x11AB: 0x3134, // ㄴ (U+11AB) -> ㄴ (U+3134)
		0x11AC: 0x3135, // ㄵ (U+11AC) -> ㄵ (U+3135)
		0x11AD: 0x3136, // ㄶ (U+11AD) -> ㄶ (U+3136)
		0x11AE: 0x3137, // ㄷ (U+11AE) -> ㄷ (U+3137)
		0x11AF: 0x3139, // ㄹ (U+11AF) -> ㄹ (U+3139)
		0x11B0: 0x313A, // ㄺ (U+11B0) -> ㄺ (U+313A)
		0x11B1: 0x313B, // ㄻ (U+11B1) -> ㄻ (U+313B)
		0x11B2: 0x313C, // ㄼ (U+11B2) -> ㄼ (U+313C)
		0x11B3: 0x313D, // ㄽ (U+11B3) -> ㄽ (U+313D)
		0x11B4: 0x313E, // ㄾ (U+11B4) -> ㄾ (U+313E)
		0x11B5: 0x313F, // ㄿ (U+11B5) -> ㄿ (U+313F)
		0x11B6: 0x3140, // ㅀ (U+11B6) -> ㅀ (U+3140)
		0x11B7: 0x3141, // ㅁ (U+11B7) -> ㅁ (U+3141)
		0x11B8: 0x3142, // ㅂ (U+11B8) -> ㅂ (U+3142)
		0x11B9: 0x3144, // ㅄ (U+11B9) -> ㅄ (U+3144)
		0x11BA: 0x3145, // ㅅ (U+11BA) -> ㅅ (U+3145)
		0x11BB: 0x3146, // ㅆ (U+11BB) -> ㅆ (U+3146)
		0x11BC: 0x3147, // ㅇ (U+11BC) -> ㅇ (U+3147)
		0x11BD: 0x3148, // ㅈ (U+11BD) -> ㅈ (U+3148)
		0x11BE: 0x314A, // ㅊ (U+11BE) -> ㅊ (U+314A)
		0x11BF: 0x314B, // ㅋ (U+11BF) -> ㅋ (U+314B)
		0x11C0: 0x314C, // ㅌ (U+11C0) -> ㅌ (U+314C)
		0x11C1: 0x314D, // ㅍ (U+11C1) -> ㅍ (U+314D)
		0x11C2: 0x314E, // ㅎ (U+11C2) -> ㅎ (U+314E)
	}

	// Hangul Letter Sios -> Hangul Choseong Sios
	letterToChoseongSiosMap = map[Jamo]Jamo{
		0x3131: 0x1100, // ㄱ (U+3131) -> ㄱ (U+1100)
		0x3132: 0x1101, // ㄲ (U+3132) -> ㄲ (U+1101)
		0x3133: 0x11AA, // ㄳ (U+3133) -> ㄳ (U+11AA)
		0x3134: 0x1102, // ㄴ (U+3134) -> ㄴ (U+1102)
		0x3135: 0x11AC, // ㄵ (U+3135) -> ㄵ (U+11AC)
		0x3136: 0x11AD, // ㄶ (U+3136) -> ㄶ (U+11AD)
		0x3137: 0x1103, // ㄷ (U+3137) -> ㄷ (U+1103)
		0x3138: 0x1104, // ㄸ (U+3138) -> ㄸ (U+1104)
		0x3139: 0x1105, // ㄹ (U+3139) -> ㄹ (U+1105)
		0x313A: 0x11B0, // ㄺ (U+313A) -> ㄺ (U+11B0)
		0x313B: 0x11B1, // ㄻ (U+313B) -> ㄻ (U+11B1)
		0x313C: 0x11B2, // ㄼ (U+313C) -> ㄼ (U+11B2)
		0x313D: 0x11B3, // ㄽ (U+313D) -> ㄽ (U+11B3)
		0x313E: 0x11B4, // ㄾ (U+313E) -> ㄾ (U+11B4)
		0x313F: 0x11B5, // ㄿ (U+313F) -> ㄿ (U+11B5)
		0x3140: 0x11B6, // ㅀ (U+3140) -> ㅀ (U+11B6)
		0x3141: 0x1106, // ㅁ (U+3141) -> ㅁ (U+1106)
		0x3142: 0x1107, // ㅂ (U+3142) -> ㅂ (U+1107)
		0x3143: 0x1108, // ㅃ (U+3143) -> ㅃ (U+1108)
		0x3144: 0x11B9, // ㅄ (U+3144) -> ㅄ (U+11B9)
		0x3145: 0x1109, // ㅅ (U+3145) -> ㅅ (U+1109)
		0x3146: 0x110A, // ㅆ (U+3146) -> ㅆ (U+110A)
		0x3147: 0x110B, // ㅇ (U+3147) -> ㅇ (U+110B)
		0x3148: 0x110C, // ㅈ (U+3148) -> ㅈ (U+110C)
		0x3149: 0x110D, // ㅉ (U+3149) -> ㅉ (U+110D)
		0x314A: 0x110E, // ㅊ (U+314A) -> ㅊ (U+110E)
		0x314B: 0x110F, // ㅋ (U+314B) -> ㅋ (U+110F)
		0x314C: 0x1110, // ㅌ (U+314C) -> ㅌ (U+1110)
		0x314D: 0x1111, // ㅍ (U+314D) -> ㅍ (U+1111)
		0x314E: 0x1112, // ㅎ (U+314E) -> ㅎ (U+1112)
		0x314F: 0x1161, // ㅏ (U+314F) -> ㅏ (U+1161)
		0x3150: 0x1162, // ㅐ (U+3150) -> ㅐ (U+1162)
		0x3151: 0x1163, // ㅑ (U+3151) -> ㅑ (U+1163)
		0x3152: 0x1164, // ㅒ (U+3152) -> ㅒ (U+1164)
		0x3153: 0x1165, // ㅓ (U+3153) -> ㅓ (U+1165)
		0x3154: 0x1166, // ㅔ (U+3154) -> ㅔ (U+1166)
		0x3155: 0x1167, // ㅕ (U+3155) -> ㅕ (U+1167)
		0x3156: 0x1168, // ㅖ (U+3156) -> ㅖ (U+1168)
		0x3157: 0x1169, // ㅗ (U+3157) -> ㅗ (U+1169)
		0x3158: 0x116A, // ㅘ (U+3158) -> ㅘ (U+116A)
		0x3159: 0x116B, // ㅙ (U+3159) -> ㅙ (U+116B)
		0x315A: 0x116C, // ㅚ (U+315A) -> ㅚ (U+116C)
		0x315B: 0x116D, // ㅛ (U+315B) -> ㅛ (U+116D)
		0x315C: 0x116E, // ㅜ (U+315C) -> ㅜ (U+116E)
		0x315D: 0x116F, // ㅝ (U+315D) -> ㅝ (U+116F)
		0x315E: 0x1170, // ㅞ (U+315E) -> ㅞ (U+1170)
		0x315F: 0x1171, // ㅟ (U+315F) -> ㅟ (U+1171)
		0x3160: 0x1172, // ㅠ (U+3160) -> ㅠ (U+1172)
		0x3161: 0x1173, // ㅡ (U+3161) -> ㅡ (U+1173)
		0x3162: 0x1174, // ㅢ (U+3162) -> ㅢ (U+1174)
		0x3163: 0x1175, // ㅣ (U+3163) -> ㅣ (U+1175)
	}

	// 초성 -> 종성
	choseongToJongseongMap = map[Jamo]Jamo{
		0x1100: 0x11A8, // ㄱ
		0x1101: 0x11A9, // ㄲ
		0x1102: 0x11AB, // ㄴ
		0x1103: 0x11AC, // ㄷ
		0x1105: 0x11AF, // ㄹ
		0x1106: 0x11B7, // ㅁ
		0x1107: 0x11B8, // ㅂ
		0x1109: 0x11BA, // ㅅ
		0x110A: 0x11BB, // ㅆ
		0x110B: 0x11BC, // ㅇ
		0x110C: 0x11BD, // ㅈ
		0x110E: 0x11BE, // ㅊ
		0x110F: 0x11BF, // ㅋ
		0x1110: 0x11C0, // ㅌ
		0x1111: 0x11C1, // ㅍ
		0x1112: 0x11C2, // ㅎ
	}

	// 종성 -> 초성
	jongseongToChoseong = map[Jamo]Jamo{
		0x11A8: 0x1100, // ㄱ
		0x11A9: 0x1101, // ㄲ
		0x11AB: 0x1102, // ㄴ
		0x11AC: 0x1103, // ㄷ
		0x11AF: 0x1105, // ㄹ
		0x11B7: 0x1106, // ㅁ
		0x11B8: 0x1107, // ㅂ
		0x11BA: 0x1109, // ㅅ
		0x11BB: 0x110A, // ㅆ
		0x11BC: 0x110B, // ㅇ
		0x11BD: 0x110C, // ㅈ
		0x11BE: 0x110E, // ㅊ
		0x11BF: 0x110F, // ㅋ
		0x11C0: 0x1110, // ㅌ
		0x11C1: 0x1111, // ㅍ
		0x11C2: 0x1112, // ㅎ
	}

	// 복합 초성 -> 중성
	complexJungseongMap = map[string]Jamo{
		"ㅗㅏ": 0x116A, // ㅘ
		"ㅗㅐ": 0x116B, // ㅙ
		"ㅗㅣ": 0x116C, // ㅚ
		"ㅜㅓ": 0x116F, // ㅝ
		"ㅜㅔ": 0x1170, // ㅞ
		"ㅜㅣ": 0x1171, // ㅟ
		"ㅡㅣ": 0x1174, // ㅢ
		"ㅕㅣ": 0x1168, // ㅖ
		"ㅏㅣ": 0x1162, // ㅐ
		"ㅑㅣ": 0x1164, // ㅒ
		"ㅓㅣ": 0x1166, // ㅔ
	}

	// 복합 중성 -> 초성
	complexJungseongReversedMap = map[Jamo]string{
		0x116A: "ㅗㅏ", // ㅘ
		0x116B: "ㅗㅐ", // ㅙ
		0x116C: "ㅗㅣ", // ㅚ
		0x116F: "ㅜㅓ", // ㅝ
		0x1170: "ㅜㅔ", // ㅞ
		0x1171: "ㅜㅣ", // ㅟ
		0x1174: "ㅡㅣ", // ㅢ
		0x1168: "ㅕㅣ", // ㅖ
		0x1162: "ㅏㅣ", // ㅐ
		0x1164: "ㅑㅣ", // ㅒ
		0x1166: "ㅓㅣ", // ㅔ
	}

	// 복합 초성 -> 종성
	complexJongseongMap = map[string]Jamo{
		"ㄱㄱ": 0x11A9, // ㄲ
		"ㄱㅅ": 0x11AA, // ㄳ
		"ㄴㅈ": 0x11AD, // ㄵ
		"ㄴㅎ": 0x11AE, // ㄶ
		"ㄹㄱ": 0x11B0, // ㄺ
		"ㄹㅁ": 0x11B1, // ㄻ
		"ㄹㅂ": 0x11B2, // ㄼ
		"ㄹㅅ": 0x11B3, // ㄽ
		"ㄹㅌ": 0x11B4, // ㄾ
		"ㄹㅍ": 0x11B5, // ㄿ
		"ㄹㅎ": 0x11B6, // ㅀ
		"ㅂㅅ": 0x11B9, // ㅄ
		"ㅅㅅ": 0x11BB, // ㅆ
	}

	// 복합 종성 -> 초성
	complexJongseongReversedMap = map[Jamo]string{
		0x11A9: "ㄱㄱ", // ㄲ
		0x11AA: "ㄱㅅ", // ㄳ
		0x11AD: "ㄴㅈ", // ㄵ
		0x11AE: "ㄴㅎ", // ㄶ
		0x11B0: "ㄹㄱ", // ㄺ
		0x11B1: "ㄹㅁ", // ㄻ
		0x11B2: "ㄹㅂ", // ㄼ
		0x11B3: "ㄹㅅ", // ㄽ
		0x11B4: "ㄹㅌ", // ㄾ
		0x11B5: "ㄹㅍ", // ㄿ
		0x11B6: "ㄹㅎ", // ㅀ
		0x11B9: "ㅂㅅ", // ㅄ
		0x11BB: "ㅅㅅ", // ㅆ
	}

	// 복합 받침
	doubleBatchim = map[rune]bool{
		2:  true, // ㄲ
		3:  true, // ㄳ
		5:  true, // ㄵ
		6:  true, // ㄶ
		9:  true, // ㄺ
		10: true, // ㄻ
		11: true, // ㄼ
		12: true, // ㄽ
		13: true, // ㄾ
		14: true, // ㄿ
		15: true, // ㅀ
		18: true, // ㅄ
		20: true, // ㅆ
	}

	// 초성 로마자
	choseongRomaja = map[string]string{
		"ㄱ": "g",
		"ㄲ": "kk",
		"ㄴ": "n",
		"ㄷ": "d",
		"ㄸ": "tt",
		"ㄹ": "r",
		"ㅁ": "m",
		"ㅂ": "b",
		"ㅃ": "pp",
		"ㅅ": "s",
		"ㅆ": "ss",
		"ㅇ": "",
		"ㅈ": "j",
		"ㅉ": "jj",
		"ㅊ": "ch",
		"ㅋ": "k",
		"ㅌ": "t",
		"ㅍ": "p",
		"ㅎ": "h",
	}

	// 중성 로마자
	jungseongRomaja = map[string]string{
		"ㅏ": "a",
		"ㅐ": "ae",
		"ㅑ": "ya",
		"ㅒ": "yae",
		"ㅓ": "eo",
		"ㅔ": "e",
		"ㅕ": "yeo",
		"ㅖ": "ye",
		"ㅗ": "o",
		"ㅘ": "wa",
		"ㅙ": "wae",
		"ㅚ": "oe",
		"ㅛ": "yo",
		"ㅜ": "u",
		"ㅝ": "wo",
		"ㅞ": "we",
		"ㅟ": "wi",
		"ㅠ": "yu",
		"ㅡ": "eu",
		"ㅢ": "ui",
		"ㅣ": "i",
	}

	// 종성 로마자
	jongseongRomaja = map[string]string{
		"":  "",
		"ㄱ": "k",
		"ㄲ": "k",
		"ㄳ": "ks",
		"ㄴ": "n",
		"ㄵ": "nj",
		"ㄶ": "nh",
		"ㄷ": "t",
		"ㄹ": "l",
		"ㄺ": "lk",
		"ㄻ": "lm",
		"ㄼ": "lb",
		"ㄽ": "ls",
		"ㄾ": "lt",
		"ㄿ": "lp",
		"ㅀ": "lh",
		"ㅁ": "m",
		"ㅂ": "p",
		"ㅄ": "ps",
		"ㅅ": "t",
		"ㅆ": "t",
		"ㅇ": "ng",
		"ㅈ": "t",
		"ㅊ": "t",
		"ㅋ": "k",
		"ㅌ": "t",
		"ㅍ": "p",
		"ㅎ": "h",
	}
)

var (
	digitsHangul = [...]string{
		"", "십", "백", "천",
	}
	digitsHangul2 = [...]string{
		"", "만", "억", "조", "경",
		"해", "자", "양", "구", "간",
		"정", "재", "극", "항하사", "아승기",
		"나유타", "불가사의", "무량대수",
	}
	numberHangul = map[int32]string{
		'0': "영",
		'1': "일",
		'2': "이",
		'3': "삼",
		'4': "사",
		'5': "오",
		'6': "육",
		'7': "칠",
		'8': "팔",
		'9': "구",
	}
	daysHangul = map[int]string{
		1:  "하루",
		2:  "이틀",
		3:  "사흘",
		4:  "나흘",
		5:  "닷새",
		6:  "엿새",
		7:  "이레",
		8:  "여드레",
		9:  "아흐레",
		10: "열흘",
		11: "열하루",
		12: "열이틀",
		13: "열사흘",
		14: "열나흘",
		15: "보름",
		16: "열엿새",
		17: "열이레",
		18: "열여드레",
		19: "열아흐레",
		20: "스무날",
		21: "스무하루",
		22: "스무이틀",
		23: "스무사흘",
		24: "스무나흘",
		25: "스무닷새",
		26: "스무엿새",
		27: "스무이레",
		28: "스무여드레",
		29: "스무아흐레",
		30: "서른날",
	}
	weekdayHangul = map[time.Weekday]string{
		time.Sunday:    "일",
		time.Monday:    "월",
		time.Tuesday:   "화",
		time.Wednesday: "수",
		time.Thursday:  "목",
		time.Friday:    "금",
		time.Saturday:  "토",
	}
)

// GetChoseong 문자열을 받아서 초성 단위로 분리하여 반환합니다.
func GetChoseong(word string) string {
	return Disassemble(word).GetChoseong()
}

// NumberToHangul 숫자를 한글로 변환합니다.
func NumberToHangul(number string) string {
	var sb strings.Builder

	for _, ch := range strings.TrimSpace(number) {
		if (ch >= '0' && ch <= '9') || ch == '.' {
			sb.WriteRune(ch)
		}
	}

	fields := strings.Split(sb.String(), ".")

	sb.Reset()

	for i, ch := range fields[0] {
		digitNumber := (len(fields[0]) - i - 1) % len(digitsHangul)
		digitNumber2 := (len(fields[0]) - i - 1) / len(digitsHangul)
		if digitNumber2 >= len(digitsHangul2) {
			return ""
		}

		if ch != '0' {
			sb.WriteString(numberHangul[ch])
			sb.WriteString(digitsHangul[digitNumber])
		}
		if digitNumber == 0 {
			sb.WriteString(digitsHangul2[digitNumber2])
		}
	}

	if len(fields) > 1 {
		sb.WriteString("점")
		for _, ch := range fields[1] {
			sb.WriteString(numberHangul[ch])
		}
	}
	return sb.String()
}

// HasBatchim 받침이 있는지 판단합니다.
func HasBatchim(str string, onlyDouble ...bool) bool {
	lastRune, _ := utf8.DecodeLastRuneInString(str)
	if lastRune < baseHangul || lastRune > lastHangul {
		return false
	}

	final := (lastRune - baseHangul) % numJongseong
	if final == 0 {
		return false
	}

	if len(onlyDouble) > 0 {
		if onlyDouble[0] {
			return doubleBatchim[final]
		} else {
			return !doubleBatchim[final]
		}
	}
	return true
}

// JosaPick 단어와 조사를 받아 적절한 조사를 반환합니다.
// 지원하는 조사: 이/가, 을/를, 은/는, 으로/로, 와/과,
// 이나/나, 이란/란, 아/야, 이랑/랑, 이에요/예요,
// 으로서/로서, 으로써/로써, 으로부터/로부터, 이라/라
func JosaPick(word, josaType string) string {
	hasBatchim := HasBatchim(word)

	switch josaType {
	case "이/가":
		if hasBatchim {
			return "이"
		} else {
			return "가"
		}
	case "을/를":
		if hasBatchim {
			return "을"
		} else {
			return "를"
		}
	case "은/는":
		if hasBatchim {
			return "은"
		} else {
			return "는"
		}
	case "으로/로":
		if hasBatchim {
			return "으로"
		} else {
			return "로"
		}
	case "와/과":
		if hasBatchim {
			return "과"
		} else {
			return "와"
		}
	case "이나/나":
		if hasBatchim {
			return "이나"
		} else {
			return "나"
		}
	case "이란/란":
		if hasBatchim {
			return "이란"
		} else {
			return "란"
		}
	case "아/야":
		if hasBatchim {
			return "아"
		} else {
			return "야"
		}
	case "이랑/랑":
		if hasBatchim {
			return "이랑"
		} else {
			return "랑"
		}
	case "이에요/예요":
		if hasBatchim {
			return "이에요"
		} else {
			return "예요"
		}
	case "으로서/로서":
		if hasBatchim {
			return "으로서"
		} else {
			return "로서"
		}
	case "으로써/로써":
		if hasBatchim {
			return "으로써"
		} else {
			return "로써"
		}
	case "으로부터/로부터":
		if hasBatchim {
			return "으로부터"
		} else {
			return "로부터"
		}
	case "이라/라":
		if hasBatchim {
			return "이라"
		} else {
			return "라"
		}
	}
	return josaType
}

// Josa 단어와 조사를 받아 적절한 조사를 붙여 반환합니다.
// 지원하는 조사: 이/가, 을/를, 은/는, 으로/로, 와/과,
// 이나/나, 이란/란, 아/야, 이랑/랑, 이에요/예요,
// 으로서/로서, 으로써/로써, 으로부터/로부터, 이라/라
func Josa(word, josaType string) string {
	return word + JosaPick(word, josaType)
}

// Assemble 문자열을 받아서 적절하게 합쳐서 반환합니다.
func Assemble(str string) string {
	result := make(Daneo, 0, len(str))
	index := -1

	eumjeolList := Disassemble(str)

	for _, eumjeol := range eumjeolList {
		if (eumjeol.isHangul()) && index >= 0 {
			if !result[index].Choseong.Empty() && result[index].Jungseong.Empty() &&
				eumjeol.Choseong.Empty() && !eumjeol.Jungseong.Empty() && eumjeol.Jongseong.Empty() {
				result[index].Jungseong = eumjeol.Jungseong
				continue
			}
			if !result[index].Jongseong.Empty() &&
				eumjeol.Choseong.Empty() && !eumjeol.Jungseong.Empty() && eumjeol.Jongseong.Empty() {
				eumjeol.Choseong = result[index].Jongseong.toChoseong()
				result[index].Jongseong = 0
				index++
				result = append(result, eumjeol)
				continue
			}
			if !result[index].Choseong.Empty() && !result[index].Jungseong.Empty() && result[index].Jongseong.Empty() &&
				!eumjeol.Choseong.Empty() && eumjeol.Jungseong.Empty() && eumjeol.Jongseong.Empty() {
				result[index].Jongseong = eumjeol.Choseong.toJongseong()
				continue
			}
			if !result[index].Choseong.Empty() && !result[index].Jungseong.Empty() && result[index].Jongseong.Empty() &&
				eumjeol.Choseong.Empty() && !eumjeol.Jungseong.Empty() && eumjeol.Jongseong.Empty() {
				cplx := result[index].Jungseong.toHangulLetterSios().toChoseong().String() +
					eumjeol.Jungseong.toHangulLetterSios().String()
				if v, ok := complexJungseongMap[cplx]; ok {
					result[index].Jungseong = v
					continue
				}
			}
			if !result[index].Choseong.Empty() && !result[index].Jungseong.Empty() && !result[index].Jongseong.Empty() &&
				!eumjeol.Choseong.Empty() && eumjeol.Jungseong.Empty() && eumjeol.Jongseong.Empty() {
				cplx := result[index].Jongseong.toHangulLetterSios().toChoseong().String() +
					eumjeol.Choseong.toHangulLetterSios().String()
				if v, ok := complexJongseongMap[cplx]; ok {
					result[index].Jongseong = v
					continue
				}
			}
		}

		index++
		result = append(result, eumjeol)
	}
	return result.Assemble()
}

// CanBeChoseong 문자열이 초성으로 사용될 수 있는지 판단합니다.
func CanBeChoseong(character string) bool {
	if len(character) == 0 {
		return false
	}

	eumjeolList := Disassemble(character)
	if len(eumjeolList) != 1 {
		return false
	}
	return (eumjeolList[0].Choseong >= baseChoseong && eumjeolList[0].Choseong < baseChoseong+numChoseong) &&
		(eumjeolList[0].Jungseong.Empty() && eumjeolList[0].Jongseong.Empty())
}

// CanBeJungseong 문자열이 중성으로 사용될 수 있는지 판단합니다.
func CanBeJungseong(character string) bool {
	if len(character) == 0 {
		return false
	}

	eumjeolList := Disassemble(character)
	switch len(eumjeolList) {
	case 1:
		return (eumjeolList[0].Jungseong >= baseJungseong && eumjeolList[0].Jungseong < baseJungseong+numJungseong) &&
			(eumjeolList[0].Choseong.Empty() && eumjeolList[0].Jongseong.Empty())
	case 2:
		cplx := eumjeolList[0].Jungseong.toHangulLetterSios().toChoseong().String() +
			eumjeolList[1].Jungseong.toHangulLetterSios().String()
		_, ok := complexJungseongMap[cplx]
		return ok
	}
	return false
}

// CanBeJongseong 문자열이 종성으로 사용될 수 있는지 판단합니다.
func CanBeJongseong(character string) bool {
	if len(character) == 0 {
		return false
	}

	eumjeolList := Disassemble(character)
	switch len(eumjeolList) {
	case 1:
		return (eumjeolList[0].Choseong.toJongseong() >= baseJongseong && eumjeolList[0].Choseong.toJongseong() < baseJongseong+numJongseong) &&
			(eumjeolList[0].Jungseong.Empty() && eumjeolList[0].Jongseong.Empty())
	case 2:
		cplx := eumjeolList[0].Choseong.toHangulLetterSios().toChoseong().String() +
			eumjeolList[1].Choseong.toHangulLetterSios().String()
		_, ok := complexJongseongMap[cplx]
		return ok
	}
	return false
}

// CombineCharacter 초성, 중성, 종성을 합쳐 반환합니다.
func CombineCharacter(choseong, jungseong string, jongseong ...string) string {
	if len(jongseong) > 0 {
		return Assemble(choseong + jungseong + jongseong[0])
	}
	return Assemble(choseong + jungseong)
}

// CombineVowels 두 모음을 합쳐 반환합니다.
func CombineVowels(vowel1, vowel2 string) string {
	cplx := vowel1 + vowel2
	v, ok := complexJungseongMap[cplx]
	if ok {
		return v.toHangulLetterSios().String()
	}
	return cplx
}

// Days 일자를 한글로 변환합니다.
func Days(day int) string {
	return daysHangul[day]
}

// Weekday 요일을 한글로 변환합니다.
func Weekday(weekday time.Weekday, full ...bool) string {
	if len(full) > 0 && full[0] {
		return weekdayHangul[weekday] + "요일"
	}
	return weekdayHangul[weekday]
}

// Disassemble 문자열을 받아서 분해하여 Daneo 로 반환합니다.
func Disassemble(str string) Daneo {
	result := make(Daneo, 0, utf8.RuneCountInString(str))

	for _, ch := range str {
		var eumjeol Eumjeol

		if ch >= baseHangul && ch <= baseHangul+numChoseong*numJungseong*numJongseong-1 {
			ch -= baseHangul
			cho := ch / (numJungseong * numJongseong)
			jung := (ch % (numJungseong * numJongseong)) / numJongseong
			jong := ch % numJongseong

			eumjeol.Choseong = Jamo(baseChoseong + cho).toHangulChoseongSios()
			eumjeol.Jungseong = Jamo(baseJungseong + jung).toHangulChoseongSios()
			if jong != 0 {
				eumjeol.Jongseong = Jamo(baseJongseong + jong).toHangulChoseongSios()
			}
		} else {
			j := Jamo(ch).toHangulChoseongSios()
			if j >= baseJungseong && j <= baseJungseong+numJungseong {
				eumjeol.Jungseong = j
			} else {
				eumjeol.Choseong = j
			}
		}

		result = append(result, eumjeol)
	}
	return result
}

// Romanize 로마자로 변환합니다.
func Romanize(str string) string {
	eumjeolList := Disassemble(str)
	if len(eumjeolList) == 0 {
		return ""
	}

	var sb strings.Builder

	for _, eumjeol := range eumjeolList {
		sb.WriteString(choseongRomaja[eumjeol.Choseong.toHangulLetterSios().String()])
		sb.WriteString(jungseongRomaja[eumjeol.Jungseong.toHangulLetterSios().String()])
		sb.WriteString(jongseongRomaja[eumjeol.Jongseong.toHangulLetterSios().String()])
	}
	return sb.String()
}
