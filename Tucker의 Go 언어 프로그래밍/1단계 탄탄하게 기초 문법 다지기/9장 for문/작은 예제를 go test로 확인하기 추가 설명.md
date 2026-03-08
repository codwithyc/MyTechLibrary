## 작은 예제를 go test로 확인하기 추가 설명

반복문을 공부하면 합계 구하기, 최댓값 찾기, 특정 조건 세기 같은 작은 함수를 많이 작성하게 됩니다.  
이때 결과를 매번 `fmt.Println()`으로 확인할 수도 있지만, Go를 제대로 익히려면 `go test`를 일찍 경험하는 것이 좋습니다.

이번 추가 설명은 `for` 문법 자체가 아니라, 반복 로직을 테스트로 확인하는 기본 흐름을 정리합니다.

### 출력 확인과 테스트는 다르다
---

처음에는 다음처럼 결과를 직접 출력해 확인합니다.

```go
func Sum(numbers []int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}
```

```go
fmt.Println(Sum([]int{1, 2, 3}))
```

이 방식은 빠르지만, 매번 눈으로 결과를 확인해야 합니다.  
테스트를 작성하면 기대값을 코드로 남길 수 있습니다.

### 테스트 파일은 _test.go로 끝난다
---

Go 테스트 파일은 보통 같은 패키지 안에 만들고, 이름을 `_test.go`로 끝냅니다.

```text
sum.go
sum_test.go
```

테스트 함수는 `Test`로 시작하고, `*testing.T`를 매개변수로 받습니다.

```go
func TestSum(t *testing.T) {
    got := Sum([]int{1, 2, 3})
    want := 6

    if got != want {
        t.Fatalf("got %d, want %d", got, want)
    }
}
```

실행은 다음 명령어로 합니다.

```bash
go test
```

### 반복문 로직은 테스트와 잘 어울린다
---

반복문은 조건 하나가 바뀌면 결과가 달라지기 쉽습니다.  
예를 들어 빈 슬라이스, 원소가 하나인 슬라이스, 음수가 섞인 슬라이스를 모두 확인하고 싶을 수 있습니다.

이럴 때는 여러 케이스를 표로 정리하는 방식이 좋습니다.

```go
func TestSum(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {name: "positive", nums: []int{1, 2, 3}, want: 6},
        {name: "empty", nums: []int{}, want: 0},
        {name: "negative", nums: []int{-1, 1}, want: 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Sum(tt.nums)
            if got != tt.want {
                t.Fatalf("got %d, want %d", got, tt.want)
            }
        })
    }
}
```

이런 방식을 table-driven test라고 부릅니다.  
Go에서 매우 자주 보이는 테스트 스타일입니다.

### go test ./...는 하위 패키지까지 검사한다
---

현재 패키지만 테스트하려면 다음처럼 실행합니다.

```bash
go test
```

프로젝트 전체의 모든 하위 패키지를 테스트하려면 다음 명령을 사용합니다.

```bash
go test ./...
```

`./...`는 현재 디렉터리 아래의 모든 패키지를 의미합니다.  
프로젝트가 커지면 커밋하기 전에 `go test ./...`로 전체 확인을 하는 습관이 좋습니다.

### 정리
---

반복문을 배울 때 테스트를 함께 익히면, 결과 확인 방식이 훨씬 단단해집니다.

정리하면 다음과 같습니다.

- `fmt.Println()`은 빠른 확인에는 좋지만 기대값을 코드로 남기지 못합니다.
- Go 테스트 파일은 `_test.go`로 끝납니다.
- 테스트 함수는 `Test`로 시작하고 `*testing.T`를 받습니다.
- 반복문 로직은 table-driven test로 확인하기 좋습니다.
- 프로젝트 전체 테스트는 `go test ./...`로 실행할 수 있습니다.

작은 반복문 예제부터 테스트로 확인하는 습관을 들이면, 이후 자료구조와 알고리즘을 공부할 때도 큰 도움이 됩니다.
