## 제1절 예상 실행계획
- Oracle은 Explaing plan 명령으로 실행계획을 수집한다.
- Oracle은 AutoTrace 명령어로 예상 실행계획뿐 아니라 여러 가지 실행통계도 확인할 수 있다.
- DBMS_XPLAN 패키지를 이용해 수집된 실행계획을 출력한다.
- SQL Server는 set showplan_text on 명령으로 예상 실행계획을 확인할 수 있다.
- SQL Server는 set showplan_all on 명령문을 통해 PhysicalOp(물리 연산자), LogicalOp(논리 연산자), EstimateRows(예상 로우 수) 등을 확인할 수 있다.

## 제2절 SQL 트레이스
- Oracle은 sql_trace 파라미터를 활성화함으로써 SQL 트레이스를 수집할 수 있다.
- Oracle은 수집된 트레이스 정보를 tkprof 명령어를 통해 리포트를 출력한다.
- Oracle에서 gather_plan_statistics 힌트를 이용하면 트레이스 정보를 SGA 메모리에 수집할 수 있고, 그 정보를 DBMS_XPLAN 패키지로 포맷팅할 수 있다.
- SQL Server는 statics profile, statistics io, statistics time 옵션들을 활성화(on)함으로써 SQL 트레이스 정보를 확인할 수 있다.

## 제3절 응답 시간 분석
- DBMS 내부에서 활동하는 수많은 프로세스 간에는 상호작용이 필요하며, 그 과정에서 다른 프로세스가 일을 마칠때까지 기다려야만 하는 상황이 자주 발생한다.
- DBMS는 프로세스가 OS에 CPU를 반환하고 대기할 때마다 로그를 남긴다. 오라클은 이를 '대기 이벤트(Wait Event)'라고 부르고, SQL Server에서는 '대기 유형(Wait Tyhpe)'이라고 부른다.
- 응답 시간(Response Time)을 Service Time과 Wait Time의 합으로 정의하고, 대기(Wait) 원인을 분석함으로써 병목을 해소해 나가는 성능 관리 방법론을 '응답 시간 분석(Response Time Analysis)'이라고 한다.
- 응답 시간 분석 방법론을 지원하는 많은 성능 관리 도구가 개발됐고, AWR은 이를 지원하는 Oracle 표준 도구이다.
