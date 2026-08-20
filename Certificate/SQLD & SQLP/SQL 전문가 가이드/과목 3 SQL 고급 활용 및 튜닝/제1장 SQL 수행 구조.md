## 제1장 SQL 수행 구조
> SQL을 받으면 -> 실행방법을 정하고 -> 데이터를 메모리/디스크에서 읽고 -> 변경된 내용은 안전하게 저장한다.

Oracle의 실제 SQL 처리도 크게 Parsing -> Optimization -> Row Source Generation -> Execution 순서입니다.

> SQL 수행 구조를 이해하기 위한 필수 개념
- Buffer = 데이터를 처리하는 동안 잠시 담아두는 메모리 공간
- Cache = 다시 사용할 가능성이 있는 데이터를 빠른 곳에 보관해 재사용하는 것
- DB Buffer Cache
    - Data File에서 읽은 Data Block을 메모리에 보관
    - 같은 Block이 다시 필요하면 Disk 대신 Memory에서 재사용
    - 핵심: Disk I/O 감소
- 서버 프로세스 = 사용자의 SQL과 명령을 직접 처리
- 백그라운드 프로세스 = 뒤에서 저장·복구·정리 같은 관리 작업 수행
Optimizer = 최적의 실행 방법을 고르는 DBMS 엔진
Optimization = Optimizer가 최적의 실행 방법을 찾는 과정

### 제1절 데이터베이스 아키텍처 → 데이터는 어디에 있고 누가 처리하지?
- Oracle과 SQL Server → 저장 원리는 비슷하지만 공간 관리 구조에는 세부 차이가 있음
    - 쉽게: 둘 다 데이터를 파일에 저장하지만, 저장 공간을 나누고 관리하는 이름과 구조가 조금 다름
    - 키워드: Oracle = Block / SQL Server = Page

- 대량 ORDER BY, GROUP BY, Hash Join 같은 작업을 메모리만으로 처리하기 어려울 때 디스크의 임시 파일을 사용
    - 쉽게: 메모리 작업대가 부족하면 TEMP라는 디스크 공간을 임시 작업대로 사용
    - 키워드: PGA = 주 작업공간 / TEMP = 메모리 부족 시 사용

- 로그 파일은 DB 버퍼 캐시의 변경사항을 기록하고 복구에 사용
    - 쉽게 : 데이터가 바뀔 때 변경 내용을 작업일지처럼 기록해 두고, 장애가 나면 그 기록으로 복구
    - 복구: Checkpoint/SCN을 기준으로 미반영된 Redo를 찾아 다시 적용한다.
    - Redo = 변경사항 재현을 위한 기록
        → Roll Forward = 변경사항 다시 적용
    - Undo → Roll Back = COMMIT 안 된 작업 취소
    - Checkpoint = "어디까지 Data File에 반영됐는지 알려주는 복구 기준점"
        - Redo를 처음부터 다시 읽지 않도록 찍어주는 복구 책깔피
    - SCN = "DB 변경의 논리적인 순서를 나타내는 번호"
        - SCN이 필요한 이유는 Redo라는 작업일지를 원래 발생할 순서대로 재현해서 정확한 상태로 복구하기 위해서
        - 실제 Oracle의 SCN은 SQL마다 하나씩 붙이는 단순한 번호표라기보다 DB의 일관됨 시점을 나타내는 논리적 시간값에 가깝습니다.
    - 키워드: Checkpoint = 복구 시작 기준 / SCN = 변경 순서

- DB가 사용하는 메모리는 모두 같이 쓰는 공간과, 각 프로세스가 혼자 쓰는 공간으로 나뉜다
    - 쉽게 : DB에는 모두 같이 쓰는 메모리와 각 프로세스가 혼자 쓰는 메모리가 따로 있다.
    - 이렇게 나누는 이유는 공유해야 하는 정보와 굳이 공유할 필요가 없는 작업이 다르기 때문입니다.
      - 여러 프로세스가 공통으로 봐야 하는 데이터나 SQL 정보는 **SGA 같은 공유 메모리**에 둡니다.
      - 한 프로세스만 잠깐 쓰는 정렬, 해시, 개인 작업 정보는 **PGA 같은 전용 메모리**에 둡니다.
    - 이렇게 나누는 이유는 크게 두 가지 이유입니다.
      1. **공유하면 효율적이기 떄문**
        - 같은 데이터 블록이나 같은 SQL 정보를 여러 프로세스가 각자 따로 가지고 있으면 메모리가 낭비됩니다.
        - 그래서 공통으로 필요한 건 한 곳에 두고 같이 씁니다.
      2. **개인 작업까지 공유하면 오히려 복잡해지기 때문**
        - 정렬이나 해시 같은 작업은 특정 프로세스만 필요합니다.
        - 이런 것까지 공유하면 서로 충돌하지 않도록 Lock이나 Latch 같은 제어가 더 많이 필요해집니다.
        - 그래서 개인 작업은 아예 각자 PGA에서 처리하는 게 더 단순하고 빠릅니다.
    - SGA(System Global Area) = 여러 프로세스가 공유하는 메모리
      - 특징
          - 여러 프로세스가 같이 사용하므로 동시 접근 충돌을 막아야 함
          - 액세스 직렬화 메커니즘 = 공유 자원에 대한 접근 순서를 제어하는 장치
              - Latch = 공유 메모리 구조를 짧게 보호
              - Buffer Lock = Buffer 접근 충돌 방지
              - Library Cache Lock/Pin = SQL/공유 객체의 사용·변경 충돌 방지
      - 구성 요소
          - DB Buffer Cache
              - Data File에서 읽은 Data Block을 메모리에 보관
              - Buffer 상태
                  - Free = 새로운 Block을 담는 데 사용 가능
                  - Clean = Memory와 Disk 내용이 같음
                  - Dirty = Memory에서 변경됐지만 Disk에는 아직 미반영
                  - Pinned = 현재 프로세스가 사용 중
              - Dirty Buffer
                  - DBWn이 Data File에 기록
                  - 기록 후 Clean 상태가 되고 필요하면 재사용 가능
              - Buffer Pool
                  - DEFAULT = 일반적인 Block 관리
                  - KEEP = 자주 사용하는 Block을 오래 유지
                  - RECYCLE = 재사용 가능성이 낮은 Block을 빨리 교체
          - Shared Pool
              - SQL, 실행계획, 메타데이터 등을 저장
              - 같은 SQL이 있으면 재사용 가능
              - Library Cache Lock/Pin으로 공유 객체 사용·변경 충돌 방지
          - Redo Log Buffer
              - DML(SELECT, INSERT, DELETE, UPDATE)로 발생한 변경사항(Redo)을 임시 저장
              - LGWR(Log Write)가 Redo Log File에 기록
      - 사용 상황
          - 데이터 읽기
          - SQL/실행계획 재사용
          - 변경 기록 관리
    - PGA(Program Global Area) = 개별 서버 프로세스만 사용하는 전용 메모리
      - 특징
          - 다른 프로세스와 공유하지 않는 개인 작업 공간
          - 공유 메모리 보호를 위한 Latch 같은 액세스 직렬화가 필요 없음
          - 버퍼 캐시에서 블록을 읽을 때보다 훨씬 빠름
      - 주요 용도
          - Sort = ORDER BY 등 정렬 작업
          - Hash = Hash Join, Hash 작업
          - 프로세스별 SQL 실행 정보
          - 세션/커서 상태 정보
      - 메모리가 부족하면
          - PGA에서 작업을 끝내지 못하고 TEMP를 사용할 수 있음
          - TEMP 사용 → Disk I/O 발생 → 성능 저하 가능
    - 키워드
        - SGA = 공유 / PGA = 개인
        - PGA = Sort + Hash + 프로세스별 작업
        - PGA 부족 → TEMP

- Process
    - Server Process
        - 사용자 프로세스와 통신
        - 사용자가 보낸 SQL과 명령을 직접 처리
    - Background Process
        - DB 뒤에서 저장·복구·정리 작업 수행
        - DBWn = Dirty Buffer → Data File
        - LGWR(Log Writer) = Redo Log Buffer → Redo Log File
        - SMON(System Monitor) = 인스턴스 복구
        - PMON(Process Monitor) = 비정상 프로세스 정리

- Wait Event(대기 이벤트)
    - 프로세스가 다른 작업이나 자원을 기다릴 때 그 이유를 기록
    - 쉽게: "DB 프로세스가 왜 멈춰 있었는지 남기는 기록"
    - 예
        - Disk I/O 대기
        - Lock 대기
        - Latch 대기
        - Network 대기
    - 용도
        - 성능 병목 원인 분석
    - 키워드
        - Wait Event = 기다림의 이유 = 병목의 흔적

### 제2절 SQL 처리 과정 → SQL을 받으면 어떻게 실행 방법을 결정하지?
- SQL의 특징
    - Structured(구조적)
        - 정해진 SQL 문법과 구조에 따라 작성

    - Set-based(집합적)
        - 데이터를 한 건씩이 아니라 집합 단위로 처리

    - Declarative(선언적)
        - "무엇을 원하는지(WHAT)는 말하지만, 어떻게 처리할지(HOW)는 직접 지시하지 않는다."
        - 사용자는 원하는 결과(WHAT)를 SQL로 작성
        - 실제 처리 방법(HOW)은 DBMS의 Optimizer가 결정
        - Optimizer = SQL의 최적 실행 방법을 선택하는 DBMS 엔진
        - 핵심: WHAT = 사용자 / HOW = Optimizer

- SQL 처리 과정
    - Parsing
        - SQL 문법, 객체, 권한 등을 확인
        - 쉽게: "SQL이 제대로 됐나?"

    - Optimization
        - Optimizer가 SQL을 어떻게 처리할지 결정하는 단계
        - 여러 실행 방법의 비용을 비교하여 가장 효율적이라고 판단한 경로를 선택
        - 쉽게: "어떤 길로 갈까?"

        - Execution Plan(실행계획)
            - Optimization의 결과
            - Optimizer가 선택한 SQL 처리 절차를 트리 구조로 표현한 것
            - 사용자는 실행계획을 보고 DB가 어떤 방법으로 SQL을 처리하려는지 확인할 수 있음
            - 예: Full Scan / Index Scan / Join 순서 / Join 방식 등
            - 쉽게: "Optimizer가 최종적으로 선택한 작전표"

        - Optimizer Hint(옵티마이저 힌트)
            - 통계정보가 부정확하거나 다른 이유로 Optimizer가 좋지 않은 실행계획을 선택할 수 있음
            - 이때 개발자가 Hint를 이용해 특정 실행 방법을 선택하도록 Optimizer를 유도할 수 있음
            - 쉽게: "Optimizer에게 이 길도 한번 써봐라고 방향을 알려주는 것"

    - Row Source Generation
        - 선택된 실행계획을 실제 실행 가능한 작업 구조로 변환
        - 쉽게: "작전표를 실제 일할 순서로 준비"

    - Execution
        - 준비된 작업 순서대로 실제 데이터를 처리
        - 쉽게: "실제로 실행"

### 제3절 데이터베이스 I/O 메커니즘 → 결정된 실행 방법으로 데이터를 실제 어떻게 읽지?
- DBMS의 I/O는 Block(Page) 단위로 이루어짐
    - Oracle = Block
    - SQL Server = Page
    - 필요한 데이터가 한 행뿐이어도 그 행이 들어 있는 Block 전체를 읽음
    - 쉽게: "데이터 하나가 필요해도 그 데이터가 들어 있는 상자(Block)째 가져옴"
    - 키워드: Oracle = Block / SQL Server = Page

- I/O 튜닝의 목표
    - Disk I/O를 최소화
    - Buffer Cache를 효율적으로 사용
    - 하지만 가장 근본적인 방법은 불필요한 Block 요청 자체를 줄이는 것
    - 쉽게: "빠르게 읽는 것보다 애초에 적게 읽는 것이 더 중요"
    - 키워드: I/O 튜닝 = 읽는 Block 수 최소화

- Logical I/O
    - Buffer Cache에 있는 Data Block을 요청하고 읽는 작업
    - Memory에서 처리되므로 Physical I/O보다 상대적으로 빠름
    - 하지만 Logical I/O가 많으면 그만큼 많은 Block을 읽고 있다는 뜻
    - 쉽게: "메모리에서 Block 읽기"
    - 키워드: Logical I/O = Buffer Cache / Memory

- Physical I/O
    - 필요한 Block이 Buffer Cache에 없어서 Disk의 Data File에서 읽어오는 작업
    - Disk 접근이 필요하므로 Logical I/O보다 비용이 큼
    - 쉽게: "메모리에 없어서 디스크까지 가지러 감"
    - 키워드: Physical I/O = Disk

- Logical I/O 요청 횟수를 줄이는 것이 중요한 이유
    - Logical I/O가 많다는 것은 많은 Block을 계속 요청하고 있다는 뜻
    - Block 요청 자체가 줄어들면 Physical I/O가 발생할 가능성도 같이 줄어듦
    - 같은 결과를 얻더라도 적은 Block을 읽는 SQL이 더 효율적
    - 쉽게: "결과가 몇 건인지보다 그 결과를 얻으려고 몇 Block을 읽었는지가 중요"
    - 키워드: Logical I/O 최소화 = SQL 튜닝 핵심

- Sequential Access
    - 레코드의 논리적 또는 물리적인 순서를 따라 차례대로 읽는 방식
    - 쉽게: "1 → 2 → 3 → 4처럼 순서대로 쭉 읽기"
    - 대량 데이터를 연속적으로 읽을 때 유리할 수 있음
    - 키워드: Sequential = 순서대로

- Random Access
    - 순서를 따르지 않고 필요한 Block을 하나씩 찾아가는 방식
    - Index를 통해 ROWID로 Table Block을 찾아갈 때 대표적으로 발생
    - 쉽게: "여기 갔다가 저기 갔다가 필요한 곳만 찾아가기"
    - 소량 조회에는 효율적일 수 있지만 대상이 많으면 반복 접근 비용이 커질 수 있음
    - 키워드: Random = Index / 여기저기 접근

- Single Block I/O
    - 한 번의 I/O Call에 하나의 Data Block만 읽는 방식
    - Index를 이용해 소량 데이터를 읽을 때 효율적
    - 쉽게: "한 번에 Block 하나"
    - 키워드: Single Block I/O = Index / 소량

- MultiBlock I/O
    - 한 번의 I/O Call에 인접한 여러 Block을 함께 읽는 방식
    - Table Full Scan처럼 대량 데이터를 읽을 때 효율적
    - 쉽게: "한 번에 Block 여러 개"
    - 키워드: MultiBlock I/O = Full Scan / 대량

- Sequential/Random과 Single/MultiBlock은 서로 다른 기준
    - Sequential / Random
        - "어떤 순서로 Block에 접근하나?"
    - Single / MultiBlock
        - "한 번의 I/O Call에 몇 개의 Block을 읽나?"
    - 쉽게:
        - Sequential / Random = 어디를 어떤 순서로 갈까?
        - Single / Multi = 한 번 갈 때 몇 개 가져올까?

- 소량 데이터를 조회할 때
    - Index Scan이 유리할 수 있음
    - Single Block I/O가 주로 사용
    - Random Access가 발생할 수 있음
    - 쉽게: "필요한 것만 정확하게 하나씩 찾아가기"

- 대량 데이터를 조회할 때
    - Table Full Scan이 유리할 수 있음
    - MultiBlock I/O로 여러 Block을 한 번에 읽음
    - 쉽게: "하나씩 찾지 말고 여러 Block을 한꺼번에 쭉 읽기"

- Index가 항상 빠른 것은 아님
    - 소량 조회에서는 Index가 효율적일 수 있음
    - 하지만 대량 조회에서 Index를 이용하면 Random Access가 너무 많이 발생해 오히려 느릴 수 있음
    - 이 경우 Full Scan + MultiBlock I/O가 더 효율적일 수 있음
    - 쉽게: "책에서 3쪽 찾기 = 목차(Index), 책 전체 읽기 = 처음부터 쭉"

- Optimizer와 I/O
    - Optimizer는 Index Scan, Full Scan 등 여러 Access Path 중 하나를 선택
    - 선택한 Access Path에 따라 읽는 Block 수와 I/O 방식이 달라짐
    - 좋은 실행계획을 위해 정확한 통계정보 등 Optimizing Factor가 필요
    - 필요하면 Hint를 이용해 더 적절한 Access Path로 유도할 수 있음

- I/O 튜닝의 최종 핵심
    - 필요한 최소 Block만 읽도록 SQL을 작성
    - Index를 사용하는 것 자체가 목적이 아니라 가장 적은 I/O로 원하는 결과를 얻는 것이 목적
    - 핵심: "최대한 적게 읽고 원하는 결과를 얻어라"

- Buffer Cache Hit Ratio
    - 요청한 Block 중 Disk까지 가지 않고 Buffer Cache에서 찾은 비율
    - Logical I/O = db block gets + consistent gets
    - Physical I/O = physical reads
    - 계산식
        - Hit Ratio = 1 - (Physical Reads / Logical Reads)
        - Hit Ratio = 1 - [physical reads / (db block gets + consistent gets)]
        - %로 계산할 때 × 100
    - 쉽게: "Block 요청 중 몇 %를 메모리에서 해결했나?"
    - 예
        - Logical I/O = 10,000
        - Physical I/O = 1,000
        - Hit Ratio = 1 - (1,000 / 10,000) = 90%
    - 키워드
        - Logical = db block gets + consistent gets
        - Physical = physical reads
        - Hit Ratio = 1 - Physical / Logical

- Oracle Wait Event 주의
    - db file sequential read
        - Single Block I/O
        - Index를 이용한 Random Access에서 주로 발생

    - db file scattered read
        - MultiBlock I/O
        - Full Scan에서 주로 발생

- Index만으로 처리 가능한 SQL
    - WHERE 조건 컬럼
    - SELECT 컬럼
    - ORDER BY에 필요한 컬럼
      등이 Index에 모두 있으면
    - Table Access를 생략할 수 있음
    - Table Block I/O 감소
