## 값 분류와 열거형 없는 Go 추가 설명

`switch`문은 여러 값이나 조건을 분류할 때 자주 사용됩니다.  
이때 다른 언어의 `enum`을 기대하면 Go가 조금 낯설게 느껴질 수 있습니다.

Go에는 별도의 `enum` 키워드가 없습니다.  
대신 상수, 사용자 정의 타입, `iota`, `switch`를 조합해 값의 종류를 표현합니다.

### 값의 종류를 타입으로 표현할 수 있다
---

단순히 문자열이나 정수로 상태를 표현할 수도 있습니다.

```go
status := "ready"
```

하지만 상태가 여러 곳에서 사용된다면 별도 타입으로 만들 수 있습니다.

```go
type Status int

const (
    StatusReady Status = iota
    StatusRunning
    StatusDone
)
```

이렇게 하면 상태 값이라는 의미가 타입에 드러납니다.

### switch는 값 분류를 읽기 좋게 만든다
---

정의한 상태 타입은 `switch`와 잘 어울립니다.

```go
switch status {
case StatusReady:
    fmt.Println("준비")
case StatusRunning:
    fmt.Println("실행 중")
case StatusDone:
    fmt.Println("완료")
default:
    fmt.Println("알 수 없는 상태")
}
```

값의 종류가 늘어날수록 `if else`보다 `switch`가 더 읽기 쉬운 경우가 많습니다.

### Go의 상수 묶음은 enum과 완전히 같지는 않다
---

Go에서 상수 묶음으로 enum처럼 표현할 수 있지만, 다른 언어의 enum과 완전히 같지는 않습니다.

```go
var status Status = 100
```

위처럼 정의하지 않은 숫자도 타입 변환을 통해 들어올 수 있습니다.  
그래서 외부 입력을 상태 값으로 바꿀 때는 유효성 검사를 따로 해주는 것이 좋습니다.

```go
func IsValidStatus(status Status) bool {
    switch status {
    case StatusReady, StatusRunning, StatusDone:
        return true
    default:
        return false
    }
}
```

Go에서는 타입을 만들었다고 해서 모든 잘못된 값이 자동으로 막히는 것은 아닙니다.

### String 메서드를 붙이면 출력이 좋아진다
---

상태 값을 출력하면 기본적으로 숫자가 보입니다.

```go
fmt.Println(StatusReady) // 0
```

사람이 읽기 좋은 이름으로 출력하고 싶다면 `String()` 메서드를 만들 수 있습니다.

```go
func (s Status) String() string {
    switch s {
    case StatusReady:
        return "ready"
    case StatusRunning:
        return "running"
    case StatusDone:
        return "done"
    default:
        return "unknown"
    }
}
```

이제 `fmt.Println(status)`를 호출했을 때 더 의미 있는 문자열이 출력됩니다.

### default는 방어 코드가 될 수 있다
---

`switch`에서 `default`는 단순히 나머지 처리만 의미하지 않습니다.  
예상하지 못한 값이 들어왔을 때 방어하는 역할도 합니다.

```go
switch status {
case StatusReady, StatusRunning, StatusDone:
    fmt.Println(status)
default:
    return fmt.Errorf("알 수 없는 상태: %d", status)
}
```

Go에서는 값의 유효성을 개발자가 명시적으로 다루는 경우가 많습니다.  
`default`는 그런 방어 지점을 만들 때 유용합니다.

### 정리
---

`switch`문을 배우는 시점에 Go의 값 분류 방식을 함께 익히면 좋습니다.

정리하면 다음과 같습니다.

- Go에는 별도의 `enum` 키워드가 없습니다.
- 상수, 사용자 정의 타입, `iota`로 enum과 비슷한 표현을 만들 수 있습니다.
- `switch`는 값의 종류를 분류할 때 읽기 좋습니다.
- 정의하지 않은 값이 들어올 수 있으므로 유효성 검사가 필요할 수 있습니다.
- `String()` 메서드를 만들면 상태 값을 사람이 읽기 좋게 출력할 수 있습니다.

Go에서는 언어 기능이 적은 대신, 작은 기능을 조합해서 의도를 분명히 드러내는 방식을 자주 사용합니다.
