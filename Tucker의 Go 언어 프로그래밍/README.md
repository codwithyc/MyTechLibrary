## Tucker의 Go 언어 프로그래밍

![alt text](./img/book_cover.png)

## 목차

- 00. 개발 환경 구축 - 윈도우, MacOS, 리눅스

### 1단계 탄탄하게 기초 문법 다지기

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EC%9E%A5%20Hello%20Go%20World/Hello%20Go%20World.md#hello-go-wolrd">01. Hello Go World</a>    
    - 1.1 Go 역사
    - 1.2 Go 언어 특징
    - 1.3 코드가 실행되기까지
    - 1.4 Hello Go World 코드 뜯어보기

> 1장은 불필요하다고 판단 되는 내용을 제외 하고 필요한 부분만 각색해서 작성하였습니다.  
> Hello Go World만 보면 됩니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/2%EC%9E%A5%20%EB%B3%80%EC%88%98/%EB%B3%80%EC%88%98.md#%EB%B3%80%EC%88%98">02. 변수</a> 
    - 2.1 변수란?
    - 2.2 변수 선언
    - 2.3 변수에 대해 더 알아보기
    - 2.4 변수 선언의 다른 형태
    - 2.5 타입 변환
    - 2.6 변수의 범위
    - 2.7 숫자 표현
    - <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/2%EC%9E%A5%20%EB%B3%80%EC%88%98/go%20mod%20%EC%B6%94%EA%B0%80%20%EC%84%A4%EB%AA%85.md#go-mod-%EC%B6%94%EA%B0%80-%EC%84%A4%EB%AA%85">2.8 mod 추가 설명</a>  
        - 변수 학습을 진행하는 과정에서 `mod`에 대한 설명이 다소 부족하다고 판단해,  
        이번에는 실제 실습 환경을 기준으로 `mod`를 어떻게 설정하고 사용해야 하는지 보다 명확하게 정리해보려 합니다.

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 "변수란?"을 시작으로, 앞서 다루지 않았던 mod 개념과 변수 선언, 변수 선언의 다양한 방식, 타입 변환,  
> 변수의 범위, 숫자 표현 등의 내용을 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/3%EC%9E%A5%20fmt%20%ED%8C%A8%ED%82%A4%EC%A7%80%EB%A5%BC%20%EC%9D%B4%EC%9A%A9%ED%95%9C%20%ED%85%8D%EC%8A%A4%ED%8A%B8%20%EC%9E%85%EC%B6%9C%EB%A0%A5/%ED%85%8D%EC%8A%A4%ED%8A%B8%20%EC%9E%85%EC%B6%9C%EB%A0%A5.md#fmt-%ED%8C%A8%ED%82%A4%EC%A7%80%EB%A5%BC-%EC%9D%B4%EC%9A%A9%ED%95%9C-%ED%85%8D%EC%8A%A4%ED%8A%B8-%EC%9E%85%EC%B6%9C%EB%A0%A5">03. fmt 패키지를 이용한 텍스트 입출력</a>
    - 3.1 표준 입출력
    - 3.2 표준 입력
    - 3.3 키보드 입력과 Scan( ) 함수의 동작 원리
    - 핵심 요약 / 연습문제
    - <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/3%EC%9E%A5%20fmt%20%ED%8C%A8%ED%82%A4%EC%A7%80%EB%A5%BC%20%EC%9D%B4%EC%9A%A9%ED%95%9C%20%ED%85%8D%EC%8A%A4%ED%8A%B8%20%EC%9E%85%EC%B6%9C%EB%A0%A5/%ED%91%9C%EC%A4%80%20%EC%8A%A4%ED%8A%B8%EB%A6%BC%EA%B3%BC%20%ED%84%B0%EB%AF%B8%EB%84%90%20%EC%8B%A4%ED%96%89%20%ED%9D%90%EB%A6%84%20%EC%B6%94%EA%B0%80%20%EC%84%A4%EB%AA%85.md#%ED%91%9C%EC%A4%80-%EC%8A%A4%ED%8A%B8%EB%A6%BC%EA%B3%BC-%ED%84%B0%EB%AF%B8%EB%84%90-%EC%8B%A4%ED%96%89-%ED%9D%90%EB%A6%84-%EC%B6%94%EA%B0%80-%EC%84%A4%EB%AA%85">3.4 표준 스트림과 터미널 실행 흐름 추가 설명</a>  
        - `fmt`를 배우는 과정에서 함께 이해하면 좋은 stdin, stdout, stderr, 리다이렉션과 파이프 흐름을 정리했습니다.

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 `fmt` 패키지, 표준 입출력, 출력 형식, `Scan()` 함수의 동작 원리와 입력 처리 흐름을 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/4%EC%9E%A5%20%EC%97%B0%EC%82%B0%EC%9E%90/%EC%97%B0%EC%82%B0%EC%9E%90.md#%EC%97%B0%EC%82%B0%EC%9E%90">04. 연산자</a>
    - 4.1 산술 연산자
    - 4.2 비교 연산자
    - 4.3 실수 오차
    - 4.4 논리 연산자
    - 4.5 대입 연산자
    - 4.6 연산자 우선순위
    - 핵심 요약 / 연습문제
    - <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/4%EC%9E%A5%20%EC%97%B0%EC%82%B0%EC%9E%90/%EC%BB%B4%ED%8C%8C%EC%9D%BC%20%EC%97%90%EB%9F%AC%EC%99%80%20%ED%83%80%EC%9E%85%20%EC%8B%9C%EC%8A%A4%ED%85%9C%20%EC%B6%94%EA%B0%80%20%EC%84%A4%EB%AA%85.md#%EC%BB%B4%ED%8C%8C%EC%9D%BC-%EC%97%90%EB%9F%AC%EC%99%80-%ED%83%80%EC%9E%85-%EC%8B%9C%EC%8A%A4%ED%85%9C-%EC%B6%94%EA%B0%80-%EC%84%A4%EB%AA%85">4.7 컴파일 에러와 타입 시스템 추가 설명</a>  
        - 연산자 학습 중 자주 만나는 타입 불일치와 컴파일 에러를 Go의 타입 시스템 관점에서 읽는 방법을 정리했습니다.

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 산술, 비교, 논리, 대입 연산자와 실수 오차, 연산자 우선순위의 핵심 개념을 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/5%EC%9E%A5%20%ED%95%A8%EC%88%98/%ED%95%A8%EC%88%98.md#%ED%95%A8%EC%88%98">05. 함수</a>
    - 5.1 함수 정의
    - 5.2 함수를 호출하면 생기는 일
    - 5.3 함수는 왜 쓰나?
    - 5.4 재귀 호출
    - 핵심 요약 / 연습문제
    - <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/5%EC%9E%A5%20%ED%95%A8%EC%88%98/Go%20%EB%AC%B8%EC%84%9C%20%EC%9D%BD%EA%B8%B0%EC%99%80%20%ED%95%A8%EC%88%98%20%EC%8B%9C%EA%B7%B8%EB%8B%88%EC%B2%98%20%EC%B6%94%EA%B0%80%20%EC%84%A4%EB%AA%85.md#go-%EB%AC%B8%EC%84%9C-%EC%9D%BD%EA%B8%B0%EC%99%80-%ED%95%A8%EC%88%98-%EC%8B%9C%EA%B7%B8%EB%8B%88%EC%B2%98-%EC%B6%94%EA%B0%80-%EC%84%A4%EB%AA%85">5.5 Go 문서 읽기와 함수 시그니처 추가 설명</a>  
        - 표준 라이브러리 문서를 읽고 함수 시그니처, 반환값, `go doc` 명령을 해석하는 방법을 정리했습니다.

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 함수 정의, 매개변수와 반환값, 함수 호출 흐름, 값 복사, 다중 반환, 재귀 호출의 기본 원리를 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/6%EC%9E%A5%20%EC%83%81%EC%88%98/%EC%83%81%EC%88%98.md#%EC%83%81%EC%88%98">06. 상수</a>
    - 6.1 상수 선언
    - 6.2 상수는 언제 사용하나?
    - 6.3 타입 없는 상수
    - 6.4 상수와 리터럴
    - 핵심 요약 / 연습문제
    - <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/6%EC%9E%A5%20%EC%83%81%EC%88%98/gofmt%EC%99%80%20%EC%BD%94%EB%93%9C%20%EC%8A%A4%ED%83%80%EC%9D%BC%20%EC%B6%94%EA%B0%80%20%EC%84%A4%EB%AA%85.md#gofmt%EC%99%80-%EC%BD%94%EB%93%9C-%EC%8A%A4%ED%83%80%EC%9D%BC-%EC%B6%94%EA%B0%80-%EC%84%A4%EB%AA%85">6.5 gofmt와 코드 스타일 추가 설명</a>  
        - `const` 블록을 다루며 함께 익히면 좋은 `gofmt`, 저장 시 자동 포맷, Go 코드 스타일 기준을 정리했습니다.

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 `const`, `iota`, 타입 없는 상수, 리터럴과 상수를 사용하는 기준을 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/7%EC%9E%A5%20if%EB%AC%B8/if%EB%AC%B8.md#if%EB%AC%B8">07. if문</a>
    - 7.1 if문 기본 사용법
    - 7.2 그리고 &&, 또는 ||
    - 7.3 중첩 if
    - 7.4 if 초기문; 조건문
    - 핵심 요약 / 연습문제
    - <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/7%EC%9E%A5%20if%EB%AC%B8/Go%EC%9D%98%20%EC%97%90%EB%9F%AC%20%EC%B2%98%EB%A6%AC%20%EC%82%AC%EA%B3%A0%EB%B0%A9%EC%8B%9D%20%EC%B6%94%EA%B0%80%20%EC%84%A4%EB%AA%85.md#go%EC%9D%98-%EC%97%90%EB%9F%AC-%EC%B2%98%EB%A6%AC-%EC%82%AC%EA%B3%A0%EB%B0%A9%EC%8B%9D-%EC%B6%94%EA%B0%80-%EC%84%A4%EB%AA%85">7.5 Go의 에러 처리 사고방식 추가 설명</a>  
        - `if`를 배우며 함께 익히면 좋은 `error` 값, `err != nil`, 에러를 무시하지 않는 습관을 정리했습니다.

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 if문의 기본 구조, `else if`, 논리 연산자를 활용한 조건 조합, 중첩 if, if 초기문을 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/8%EC%9E%A5%20switch%EB%AC%B8/switch%EB%AC%B8.md#switch%EB%AC%B8">08. switch문</a>
    - 8.1 switch문 동작 원리
    - 8.2 switch문을 언제 쓰는가?
    - 8.3 다양한 switch문 형태
    - 8.4 const 열거값과 switch
    - 8.5 break와 fallthrough 키워드
    - 핵심 요약 / 연습문제
    - <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/8%EC%9E%A5%20switch%EB%AC%B8/%EA%B0%92%20%EB%B6%84%EB%A5%98%EC%99%80%20%EC%97%B4%EA%B1%B0%ED%98%95%20%EC%97%86%EB%8A%94%20Go%20%EC%B6%94%EA%B0%80%20%EC%84%A4%EB%AA%85.md#%EA%B0%92-%EB%B6%84%EB%A5%98%EC%99%80-%EC%97%B4%EA%B1%B0%ED%98%95-%EC%97%86%EB%8A%94-go-%EC%B6%94%EA%B0%80-%EC%84%A4%EB%AA%85">8.6 값 분류와 열거형 없는 Go 추가 설명</a>  
        - `switch` 학습과 함께 알아두면 좋은 Go의 enum 부재, 사용자 정의 타입, 값 유효성 검사 흐름을 정리했습니다.

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 switch문의 동작 원리, 다양한 switch 형태, 상수 열거값과 switch, `break`와 `fallthrough`의 차이를 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/9%EC%9E%A5%20for%EB%AC%B8/for%EB%AC%B8.md#for%EB%AC%B8">09. for문</a>
    - 9.1 for문 동작 원리
    - 9.2 continue와 break
    - 9.3 중첩 for문
    - 9.4 중첩 for문과 break, 레이블
    - 핵심 요약 / 연습문제
    - <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/9%EC%9E%A5%20for%EB%AC%B8/%EC%9E%91%EC%9D%80%20%EC%98%88%EC%A0%9C%EB%A5%BC%20go%20test%EB%A1%9C%20%ED%99%95%EC%9D%B8%ED%95%98%EA%B8%B0%20%EC%B6%94%EA%B0%80%20%EC%84%A4%EB%AA%85.md#%EC%9E%91%EC%9D%80-%EC%98%88%EC%A0%9C%EB%A5%BC-go-test%EB%A1%9C-%ED%99%95%EC%9D%B8%ED%95%98%EA%B8%B0-%EC%B6%94%EA%B0%80-%EC%84%A4%EB%AA%85">9.5 작은 예제를 go test로 확인하기 추가 설명</a>  
        - 반복문 예제를 `fmt.Println`으로만 확인하지 않고 `go test`, `_test.go`, table-driven test로 검증하는 흐름을 정리했습니다.

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 Go의 유일한 반복문인 for문, `continue`, `break`, 조건만 있는 반복문, 중첩 반복문과 레이블을 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/10%EC%9E%A5%20%EB%B0%B0%EC%97%B4/%EB%B0%B0%EC%97%B4.md#%EB%B0%B0%EC%97%B4">10. 배열</a>
    - 10.1 배열
    - 10.2 배열 사용법
    - 10.3 배열은 연속된 메모리
    - 10.4 다중 배열
    - 핵심 요약 / 연습문제
    - <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/10%EC%9E%A5%20%EB%B0%B0%EC%97%B4/%EA%B0%92%20%EB%B3%B5%EC%82%AC%EC%99%80%20%EC%B0%B8%EC%A1%B0%EC%B2%98%EB%9F%BC%20%EB%B3%B4%EC%9D%B4%EB%8A%94%20%EA%B0%92%20%EC%B6%94%EA%B0%80%20%EC%84%A4%EB%AA%85.md#%EA%B0%92-%EB%B3%B5%EC%82%AC%EC%99%80-%EC%B0%B8%EC%A1%B0%EC%B2%98%EB%9F%BC-%EB%B3%B4%EC%9D%B4%EB%8A%94-%EA%B0%92-%EC%B6%94%EA%B0%80-%EC%84%A4%EB%AA%85">10.5 값 복사와 참조처럼 보이는 값 추가 설명</a>  
        - 배열 학습을 발판으로 슬라이스, 맵, 함수 전달에서 계속 등장하는 값 복사 감각을 정리했습니다.

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 배열 선언, 인덱스, `len`과 `range`, 연속된 메모리 구조, 값 타입으로서의 배열과 다중 배열을 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/11%EC%9E%A5%20%EA%B5%AC%EC%A1%B0%EC%B2%B4/%EA%B5%AC%EC%A1%B0%EC%B2%B4.md#%EA%B5%AC%EC%A1%B0%EC%B2%B4">11. 구조체</a>
    - 11.1 선언 및 기본 사용
    - 11.2 구조체 변수 초기화
    - 11.3 구조체를 포함하는 구조체
    - 11.4 구조체 크기
    - 11.5 프로그래밍에서 구조체의 역할
    - 핵심 요약 / 연습문제
    - <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/11%EC%9E%A5%20%EA%B5%AC%EC%A1%B0%EC%B2%B4/%EA%B3%B5%EA%B0%9C%20API%EC%99%80%20%ED%8C%A8%ED%82%A4%EC%A7%80%20%EA%B2%BD%EA%B3%84%20%EC%B6%94%EA%B0%80%20%EC%84%A4%EB%AA%85.md#%EA%B3%B5%EA%B0%9C-api%EC%99%80-%ED%8C%A8%ED%82%A4%EC%A7%80-%EA%B2%BD%EA%B3%84-%EC%B6%94%EA%B0%80-%EC%84%A4%EB%AA%85">11.6 공개 API와 패키지 경계 추가 설명</a>  
        - 구조체 필드의 대소문자에서 이어지는 공개 API, 패키지 경계, 주석 문서화 기준을 정리했습니다.

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 구조체 선언, 필드 접근, 구조체 초기화, 구조체를 포함하는 구조체, 구조체 크기와 역할을 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/12%EC%9E%A5%20%ED%8F%AC%EC%9D%B8%ED%84%B0/%ED%8F%AC%EC%9D%B8%ED%84%B0.md#%ED%8F%AC%EC%9D%B8%ED%84%B0">12. 포인터</a>
    - 12.1 포인터란?
    - 12.2 포인터는 왜 쓰나?
    - 12.3 인스턴스
    - 12.4 스택 메모리와 힙 메모리
    - 핵심 요약 / 연습문제

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 주소와 역참조, 포인터가 필요한 이유, 함수와 포인터, nil 포인터, 인스턴스, 스택과 힙 메모리의 기본 개념을 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/13%EC%9E%A5%20%EB%AC%B8%EC%9E%90%EC%97%B4/%EB%AC%B8%EC%9E%90%EC%97%B4.md#%EB%AC%B8%EC%9E%90%EC%97%B4">13. 문자열</a>
    - 13.1 문자열
    - 13.2 문자열 순회
    - 13.3 문자열 합치기
    - 13.4 문자열 구조
    - 13.5 문자열은 불변이다
    - 핵심 요약 / 연습문제

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 문자열과 바이트, `rune`, `len`, 문자열 순회, 문자열 합치기, 내부 구조와 불변성을 담았습니다.

- <a href="https://github.com/codwithyc/MyTechLibrary/blob/main/Tucker%EC%9D%98%20Go%20%EC%96%B8%EC%96%B4%20%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D/1%EB%8B%A8%EA%B3%84%20%ED%83%84%ED%83%84%ED%95%98%EA%B2%8C%20%EA%B8%B0%EC%B4%88%20%EB%AC%B8%EB%B2%95%20%EB%8B%A4%EC%A7%80%EA%B8%B0/14%EC%9E%A5%20%ED%8C%A8%ED%82%A4%EC%A7%80/%ED%8C%A8%ED%82%A4%EC%A7%80.md#%ED%8C%A8%ED%82%A4%EC%A7%80">14. 패키지</a>
    - 14.1 패키지
    - 14.2 패키지 사용하기
    - 14.3 Go 모듈
    - 14.4 패키지명과 패키지 외부 공개
    - 14.5 패키지 초기화
    - 핵심 요약 / 연습문제

> 1장과 마찬가지로, 불필요하다고 판단한 내용은 제외하고 해당 목차에 포함된 내용을 새롭게 재구성해 작성했습니다.  
> 이번 장에는 패키지의 역할, `import`, Go 모듈, 외부 공개 규칙, `init()` 함수와 패키지 분리 기준을 담았습니다.

- 15. [Project] 숫자 맞추기 게임 만들기 ★☆☆☆ - 정리 생략
    - 프로젝트 실습 장이므로 개념 정리 대상에서는 제외합니다.
    - 15.1 해법
    - 15.2 사전지식
    - 15.3 랜덤한 숫자 생성하기
    - 15.4 숫자값 입력받기
    - 15.5 숫자 맞추기 완성하기
    - 핵심 요약 / 연습문제

### 2단계 고급 기법으로 Go 레벨업하기

- <a href="">16. 슬라이스</a>
    - <a href="">16.1 슬라이스</a>
    - <a href="">16.2 슬라이스 동작 원리</a>
    - <a href="">16.3 슬라이싱</a>
    - <a href="">16.4 유용한 슬라이싱 기능 활용</a>
    - <a href="">16.5 슬라이스 정렬</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">17. 메서드</a>
    - <a href="">17.1 메서드 선언</a>
    - <a href="">17.2 메서드는 왜 필요한가?</a>
    - <a href="">17.3 포인터 메서드 vs 값 타입 메서드</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">18. 인터페이스</a>
    - <a href="">18.1 인터페이스</a>
    - <a href="">18.2 인터페이스 왜 쓰나?</a>
    - <a href="">18.3 덕 타이핑</a>
    - <a href="">18.4 인터페이스 기능 더 알기</a>
    - <a href="">18.5 인터페이스 변환하기</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">19. 함수 고급편</a>
    - <a href="">19.1 가변 인수 함수</a>
    - <a href="">19.2 defer 지연 실행</a>
    - <a href="">19.3 함수 타입 변수</a>
    - <a href="">19.4 함수 리터럴</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">20. 자료구조</a>
    - <a href="">20.1 리스트</a>
    - <a href="">20.2 링</a>
    - <a href="">20.3 맵</a>
    - <a href="">20.4 맵의 원리</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">21. 에러 핸들링</a>
    - <a href="">21.1 에러 반환</a>
    - <a href="">21.2 에러 타입</a>
    - <a href="">21.3 패닉</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">22. 고루틴과 동시성 프로그래밍</a>
    - <a href="">22.1 스레드란?</a>
    - <a href="">22.2 고루틴 사용</a>
    - <a href="">22.3 고루틴의 동작 방법</a>
    - <a href="">22.4 동시성 프로그래밍 주의점</a>
    - <a href="">22.5 뮤텍스를 이용한 동시성 문제 해결</a>
    - <a href="">22.6 뮤텍스와 데드락</a>
    - <a href="">22.7 또 다른 자원 관리 기법</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">23. 채널과 컨텍스트</a>
    - <a href="">23.1 채널 사용하기</a>
    - <a href="">23.2 컨텍스트 사용하기</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">24. 제네릭 프로그래밍</a>
    - <a href="">24.1 제네릭 프로그래밍 소개</a>
    - <a href="">24.2 제네릭 함수</a>
    - <a href="">24.3 제네릭 타입</a>
    - <a href="">24.4 언제 제네릭 프로그래밍을 사용해야 하는가?</a>
    - <a href="">24.5 제네릭을 사용해 만든 유용한 기본 패키지</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">25. [Project] 단어 검색 프로그램 만들기 ★★☆☆</a>
    - <a href="">25.1 해법</a>
    - <a href="">25.2 사전 지식</a>
    - <a href="">25.3 실행 인수 읽고 파일 목록 가져오기</a>
    - <a href="">25.4 파일을 열어서 라인 읽기</a>
    - <a href="">25.5 파일 검색 프로그램 완성하기</a>
    - <a href="">25.6 개선하기</a>
    - <a href="">핵심 요약 / 연습문제</a>

### 3단계 네트워크 서버 개발 및 성능 개선 기법 익히기

- <a href="">26. 테스트와 벤치마크하기</a>
    - <a href="">26.1 테스트 코드</a>
    - <a href="">26.2 테스트 주도 개발</a>
    - <a href="">26.3 벤치마크</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">27. 프로파일링으로 성능 개선하기</a>
    - <a href="">27.1 특정 구간 프로파일링</a>
    - <a href="">27.2 서버에서 프로파일링</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">28. [Project] HTTP 웹 서버 만들기 ★★☆☆</a>
    - <a href="">28.1 HTTP 웹 서버 만들기</a>
    - <a href="">28.2 HTTP 동작 원리</a>
    - <a href="">28.3 HTTP 쿼리 인수 사용하기</a>
    - <a href="">28.4 ServeMux 인스턴스 이용하기</a>
    - <a href="">28.5 파일 서버</a>
    - <a href="">28.6 웹 서버 테스트 코드 만들기</a>
    - <a href="">28.7 JSON 데이터 전송</a>
    - <a href="">28.8 HTTPS 웹 서버 만들기</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">29. [Project] RESTful API 서버 만들기 ★★★☆</a>
    - <a href="">29.1 해법</a>
    - <a href="">29.2 사전 지식 : RESTful API</a>
    - <a href="">29.3 RESTful API 서버 만들기</a>
    - <a href="">29.4 테스트 코드 작성하기</a>
    - <a href="">29.5 특정 학생 데이터 반환하기</a>
    - <a href="">29.6 학생 데이터 추가/삭제하기</a>
    - <a href="">29.7 RESTful API로의 발전</a>
    - <a href="">29.8 Gin으로 서버 만들기</a>
    - <a href="">핵심 요약 / 연습문제</a>

- <a href="">30. [Project] gnet과 gRPC으로 채팅 앱 만들기 ★★★★</a>
    - <a href="">30.1 gnet을 이용해서 echo 서버 제작</a>
    - <a href="">30.2 클라이언트 제작</a>
    - <a href="">30.3 채팅 서버 제작</a>
    - <a href="">30.4 gRPC란?</a>
    - <a href="">30.5 gRPC를 이용한 채팅 프로그램</a>
    - <a href="">핵심 요약 / 연습문제</a>

### Tucker 노트

- <a href="">노트 A. Go 문법 보충 수업</a>
    - <a href="">A.1 배열과 슬라이스</a>
    - <a href="">A.2 for range</a>
    - <a href="">A.3 입출력 처리</a>
    - <a href="">A.4 알아두면 유용한 go 명령어</a>
    - <a href="">A.5 cgo로 C 언어 호출하기</a>
    - <a href="">A.6 go doc</a>
    - <a href="">A.7 Embed</a>

- <a href="">노트 B. 생각하는 프로그래밍</a>
    - <a href="">B.1 Go는 객체지향 언어인가?</a>
    - <a href="">B.2 구조체에 생성자를 둘 수 있나?</a>
    - <a href="">B.3 포인터를 사용해도 복사가 일어나나?</a>
    - <a href="">B.4 값 타입을 쓸 것인가? 포인터를 쓸 것인가?</a>
    - <a href="">B.5 구체화된 객체와 관계하라고?</a>
    - <a href="">B.6 Go 언어 가비지 컬렉터</a>
