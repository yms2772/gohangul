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
BenchmarkDisassemble-10    	 9122798	       128.0 ns/op
BenchmarkAssemble
BenchmarkAssemble-10       	 7166037	       162.8 ns/op
BenchmarkRomanize
BenchmarkRomanize-10       	 3180787	       372.6 ns/op
```

## 라이센스
[MIT](https://github.com/yms2772/gohangul/blob/main/LICENSE)