[![Go Reference](https://pkg.go.dev/badge/github.com/yms2772/gohangul.svg)](https://pkg.go.dev/github.com/yms2772/gohangul)
[![codecov](https://codecov.io/github/yms2772/gohangul/graph/badge.svg?token=NALASBSYU3)](https://codecov.io/github/yms2772/gohangul)

# gohangul
gohangul은 한글을 처리하는 Go 패키지입니다.

한글을 초성, 중성, 종성으로 분리하거나 초성, 중성, 종성을 합쳐 문자로 만들 수 있습니다.

## 구조
### Jamo (자모)
* 기본적으로 `rune` 타입으로 처리합니다.
```go
type Jamo rune
```
### Eumjeol (음절)
* 초성, 중성, 종성을 `Jamo`로 하여 음절로 표현합니다. 
```go
type Eumjeol struct {
	Choseong  Jamo
	Jungseong Jamo
	Jongseong Jamo
}
```
### Daneo (단어)
* 음절의 집합체로 단어를 표현합니다.
```go
type Daneo []Eumjeol
```

## 사용법
* 테스트 코드에 모든 예시가 포함되어있습니다.
### 문자열 분리
```go
package main

import (
	"fmt"

	"github.com/yms2772/gohangul"
)

func main() {
	item := gohangul.Disassemble("안녕하세요")

	fmt.Println(item.At(0).Choseong)  // ㅇ
	fmt.Println(item.At(0).Jungseong) // ㅏ
	fmt.Println(item.At(0).Jongseong) // ㄴ
	fmt.Println(item.String())        // ㅇㅏㄴㄴㅕㅇㅎㅏㅅㅓㅣㅇㅛ
	fmt.Println(item.Assemble())      // 안녕하세요
}

```
### 문자열 병합
```go
package main

import (
	"fmt"

	"github.com/yms2772/gohangul"
)

func main() {
	item := gohangul.Assemble("ㅇㅏㄴㄴㅕㅇㅎㅏㅅㅓㅣㅇㅛ")

	fmt.Println(item) // 안녕하세요
}

```
### 조사 구분
```go
package main

import (
	"fmt"

	"github.com/yms2772/gohangul"
)

func main() {
	item := gohangul.Josa("생각", "을/를")

	fmt.Println(item) // 생각을
}
```
### 로마자 변환
```go
package main

import (
	"fmt"

	"github.com/yms2772/gohangul"
)

func main() {
	item := gohangul.Romanize("안녕하세요")

	fmt.Println(item) // annyeonghaseyo
}
```

## 벤치마크
```shell
BenchmarkDisassemble
BenchmarkDisassemble-10         	 8318858	       127.8 ns/op
BenchmarkAssemble
BenchmarkAssemble-10            	 7490484	       158.3 ns/op
BenchmarkRomanize
BenchmarkRomanize-10            	 3312685	       363.4 ns/op
BenchmarkCanBeChoseong
BenchmarkCanBeChoseong-10       	50878087	        23.46 ns/op
BenchmarkCanBeJungseong
BenchmarkCanBeJungseong-10      	42416503	        28.31 ns/op
BenchmarkCanBeJongseong
BenchmarkCanBeJongseong-10      	34962742	        34.45 ns/op
BenchmarkCombineCharacter
BenchmarkCombineCharacter-10    	 9039068	       133.4 ns/op
BenchmarkCombineVowels
BenchmarkCombineVowels-10       	60781804	        19.79 ns/op
BenchmarkDays
BenchmarkDays-10                	1000000000	         0.3196 ns/op
BenchmarkGetChoseong
BenchmarkGetChoseong-10         	11115223	       108.2 ns/op
BenchmarkHasBatchim
BenchmarkHasBatchim-10          	193764027	         6.183 ns/op
BenchmarkJosa
BenchmarkJosa-10                	61659937	        19.47 ns/op
BenchmarkJosaPick
BenchmarkJosaPick-10            	148084624	         8.089 ns/op
BenchmarkNumberToHangul
BenchmarkNumberToHangul-10      	 5481848	       218.8 ns/op
BenchmarkWeekday
BenchmarkWeekday-10             	1000000000	         0.3200 ns/op
```

## 라이센스
[MIT](https://github.com/yms2772/gohangul/blob/main/LICENSE)