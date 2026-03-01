## switch문

### switch문이란?
---
`switch`문은 여러 조건 중 하나를 선택해 실행할 때 사용하는 문법입니다.  
`if`, `else if`, `else`로도 여러 조건을 처리할 수 있지만, 조건이 많아지면 코드가 길고 복잡해집니다.

예를 들어 요일 번호에 따라 요일 이름을 출력한다고 해보겠습니다.

```go
day := 3

if day == 1 {
	fmt.Println("월요일")
} else if day == 2 {
	fmt.Println("화요일")
} else if day == 3 {
	fmt.Println("수요일")
} else {
	fmt.Println("알 수 없는 요일")
}
```

이 코드는 `switch`문으로 더 읽기 좋게 바꿀 수 있습니다.

```go
switch day {
case 1:
	fmt.Println("월요일")
case 2:
	fmt.Println("화요일")
case 3:
	fmt.Println("수요일")
default:
	fmt.Println("알 수 없는 요일")
}
```

`switch`문은 하나의 값 또는 조건을 기준으로 여러 갈래 중 하나를 선택할 때 유용합니다.

### switch문 동작 원리
---
기본 형태는 다음과 같습니다.

```go
switch 비교할값 {
case 값1:
	실행할 코드
case 값2:
	실행할 코드
default:
	어느 case에도 해당하지 않을 때 실행할 코드
}
```

`switch`문은 위에서 아래로 `case`를 비교합니다.  
일치하는 `case`를 찾으면 그 블록을 실행하고 `switch`문을 빠져나옵니다.

Go의 `switch`문은 다른 언어와 달리 각 `case` 끝에 `break`를 쓰지 않아도 됩니다.  
기본적으로 하나의 `case`만 실행하고 자동으로 끝납니다.

```go
grade := "B"

switch grade {
case "A":
	fmt.Println("아주 좋습니다.")
case "B":
	fmt.Println("좋습니다.")
case "C":
	fmt.Println("조금 더 노력해봅시다.")
default:
	fmt.Println("알 수 없는 등급입니다.")
}
```

### switch문을 언제 쓰는가?
---
`switch`문은 다음 상황에서 특히 좋습니다.

- 하나의 값을 여러 후보와 비교할 때
- 상태값에 따라 처리 흐름이 나뉠 때
- `if else`가 길게 이어져 읽기 어려울 때
- 여러 값을 같은 처리로 묶고 싶을 때

여러 값을 하나의 `case`에 묶을 수도 있습니다.

```go
month := 1

switch month {
case 12, 1, 2:
	fmt.Println("겨울")
case 3, 4, 5:
	fmt.Println("봄")
case 6, 7, 8:
	fmt.Println("여름")
case 9, 10, 11:
	fmt.Println("가을")
default:
	fmt.Println("잘못된 월입니다.")
}
```

조건이 단순히 `==` 비교라면 `switch`문이 `if else`보다 읽기 쉬운 경우가 많습니다.

### 다양한 switch문 형태
---
Go의 `switch`문은 비교할 값을 생략할 수도 있습니다.  
이 경우 각 `case`에 조건식을 직접 적습니다.

```go
score := 87

switch {
case score >= 90:
	fmt.Println("A")
case score >= 80:
	fmt.Println("B")
case score >= 70:
	fmt.Println("C")
default:
	fmt.Println("D")
}
```

이 형태는 `if else`와 비슷하지만, 조건이 여러 단계로 나뉠 때 더 정돈된 느낌을 줄 수 있습니다.

`switch`에도 초기문을 사용할 수 있습니다.

```go
switch score := 87; {
case score >= 90:
	fmt.Println("A")
case score >= 80:
	fmt.Println("B")
default:
	fmt.Println("C 이하")
}
```

초기문에서 선언한 변수는 `switch`문 안에서만 사용할 수 있습니다.  
필요한 범위를 좁게 유지할 수 있다는 점에서 `if 초기문`과 비슷합니다.

### const 열거값과 switch
---
`switch`문은 상수와 함께 사용할 때 더 강력합니다.  
상태값을 숫자나 문자열로 직접 비교하는 대신, 의미 있는 이름을 가진 상수로 표현할 수 있습니다.

```go
const (
	statusReady = iota
	statusRunning
	statusDone
	statusFailed
)
```

이제 상태에 따른 처리를 `switch`문으로 작성할 수 있습니다.

```go
status := statusRunning

switch status {
case statusReady:
	fmt.Println("준비 중입니다.")
case statusRunning:
	fmt.Println("실행 중입니다.")
case statusDone:
	fmt.Println("완료되었습니다.")
case statusFailed:
	fmt.Println("실패했습니다.")
}
```

이 방식은 숫자 `0`, `1`, `2`, `3`을 직접 비교하는 것보다 훨씬 읽기 쉽습니다.  
코드를 읽는 사람은 값 자체보다 값의 의미를 먼저 이해할 수 있습니다.

### break와 fallthrough 키워드
---
Go의 `switch`문은 기본적으로 `case` 하나를 실행하면 자동으로 끝납니다.  
따라서 일반적인 상황에서는 `break`가 필요하지 않습니다.

```go
switch number {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
}
```

그래도 `case` 안에서 중간에 빠져나가고 싶다면 `break`를 사용할 수 있습니다.

```go
switch number {
case 1:
	if skip {
		break
	}
	fmt.Println("one")
}
```

다음 `case`까지 이어서 실행하고 싶다면 `fallthrough`를 사용할 수 있습니다.

```go
switch number {
case 1:
	fmt.Println("one")
	fallthrough
case 2:
	fmt.Println("two")
}
```

하지만 `fallthrough`는 코드 흐름을 헷갈리게 만들 수 있으므로 꼭 필요한 경우에만 사용하는 것이 좋습니다.  
대부분의 Go 코드에서는 `fallthrough`를 자주 쓰지 않습니다.

### 실습 코드
---
이번 장의 실습 파일은 아래 위치에 작성했습니다.

```text
8장 switch문/실습 파일/switch_statement.go
```

실행 방법은 다음과 같습니다.

```bash
go run switch_statement.go
```

### 핵심 요약
---
- `switch`문은 여러 조건 중 하나를 선택해 실행하는 문법입니다.
- Go에서는 `case`마다 `break`를 쓰지 않아도 자동으로 빠져나옵니다.
- 여러 값을 하나의 `case`에 묶을 수 있습니다.
- 비교할 값을 생략하면 `case`에 조건식을 직접 적을 수 있습니다.
- 상수와 `switch`를 함께 쓰면 상태별 분기 코드를 읽기 쉽게 만들 수 있습니다.
- `fallthrough`는 다음 `case`까지 실행을 이어가지만, 자주 쓰는 문법은 아닙니다.
