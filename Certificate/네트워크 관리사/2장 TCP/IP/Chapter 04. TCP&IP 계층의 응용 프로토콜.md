# TCP/IP 계층의 응용 프로토콜

## 4-1. 응용 프로토콜

응용 프로토콜은 TCP/IP 계층에서 **사용자가 직접 이용하는 네트워크 서비스**를 제공하기 위한 프로토콜입니다.
웹 접속, 파일 전송, 원격 접속, 이메일 송수신, 도메인 조회, 네트워크 관리, IP 자동 할당 등이 모두 응용 프로토콜을 통해 동작합니다.

| 프로토콜   | 전체 이름                               | 주요 역할             | 기본 포트        |
| ------ | ----------------------------------- | ----------------- | ------------ |
| Telnet | Teletype Network                    | 원격 접속             | TCP 23       |
| SSH    | Secure Shell                        | 보안 원격 접속          | TCP 22       |
| FTP    | File Transfer Protocol              | 파일 전송             | TCP 20, 21   |
| HTTP   | HyperText Transfer Protocol         | 웹 문서 전송           | TCP 80       |
| HTTPS  | HyperText Transfer Protocol Secure  | 암호화된 웹 문서 전송      | TCP 443      |
| DNS    | Domain Name System                  | 도메인 이름을 IP 주소로 변환 | TCP/UDP 53   |
| SMTP   | Simple Mail Transfer Protocol       | 이메일 송신            | TCP 25       |
| POP3   | Post Office Protocol version 3      | 이메일 수신            | TCP 110      |
| IMAP   | Internet Message Access Protocol    | 이메일 서버 동기화        | TCP 143      |
| SNMP   | Simple Network Management Protocol  | 네트워크 장비 관리        | UDP 161, 162 |
| DHCP   | Dynamic Host Configuration Protocol | IP 주소 자동 할당       | UDP 67, 68   |

#### 한 줄 정리

```text
응용 프로토콜 = 사용자가 이용하는 네트워크 서비스를 제공하는 프로토콜
```

## 1) Telnet과 SSH

Telnet과 SSH는 모두 원격지에 있는 컴퓨터나 서버에 접속하여 명령을 실행하기 위한 프로토콜입니다.
둘의 핵심 차이는 **보안 여부**입니다.

### Telnet

Telnet은 원격 컴퓨터에 접속하여 명령어를 실행할 수 있게 해주는 원격 접속 프로토콜입니다.

1. Telnet의 특징

   * Teletype Network의 약자입니다.
   * 원격 서버에 접속하여 명령을 실행할 수 있습니다.
   * 기본 포트는 TCP 23번입니다.
   * 데이터가 암호화되지 않고 평문으로 전송됩니다.
   * 보안에 취약하여 현재는 거의 사용하지 않습니다.

| 구분    | 내용               |
| ----- | ---------------- |
| 전체 이름 | Teletype Network |
| 목적    | 원격 접속            |
| 기본 포트 | TCP 23           |
| 보안    | 암호화 없음           |
| 특징    | 평문 전송으로 보안에 취약   |

### SSH

SSH는 Telnet의 보안 문제를 보완한 원격 접속 프로토콜입니다.

1. SSH의 특징

   * Secure Shell의 약자입니다.
   * 원격 서버에 안전하게 접속할 수 있습니다.
   * 기본 포트는 TCP 22번입니다.
   * 통신 내용을 암호화합니다.
   * 사용자 인증, 명령 실행, 파일 전송 등에 사용됩니다.
   * 서버 관리에서 가장 많이 사용되는 원격 접속 방식입니다.

| 구분    | 내용           |
| ----- | ------------ |
| 전체 이름 | Secure Shell |
| 목적    | 보안 원격 접속     |
| 기본 포트 | TCP 22       |
| 보안    | 암호화 지원       |
| 특징    | 안전한 원격 접속 제공 |

### Telnet과 SSH 비교

| 구분     | Telnet           | SSH          |
| ------ | ---------------- | ------------ |
| 전체 이름  | Teletype Network | Secure Shell |
| 기본 포트  | TCP 23           | TCP 22       |
| 암호화    | 지원하지 않음          | 지원함          |
| 보안성    | 낮음               | 높음           |
| 데이터 전송 | 평문 전송            | 암호화 전송       |
| 현재 사용  | 거의 사용하지 않음       | 서버 관리에 널리 사용 |

#### 한 줄 정리

```text
Telnet = 암호화 없는 원격 접속
SSH = 암호화된 보안 원격 접속
```

## 2) FTP

### 개요

FTP는 File Transfer Protocol의 약자로, 네트워크를 통해 파일을 업로드하거나 다운로드하기 위한 프로토콜입니다.

1. FTP의 특징

   * 파일 전송을 위한 응용 계층 프로토콜입니다.
   * TCP 기반으로 동작합니다.
   * 제어 연결과 데이터 연결을 분리하여 사용합니다.
   * 기본적으로 인증을 위해 사용자 ID와 비밀번호를 사용합니다.
   * 암호화되지 않은 FTP는 보안에 취약할 수 있습니다.

| 구분     | 내용                     |
| ------ | ---------------------- |
| 전체 이름  | File Transfer Protocol |
| 목적     | 파일 업로드 및 다운로드          |
| 전송 계층  | TCP                    |
| 제어 포트  | TCP 21                 |
| 데이터 포트 | TCP 20 또는 동적 포트        |
| 특징     | 제어 연결과 데이터 연결 분리       |

### FTP 연결 구조

FTP는 일반적인 프로토콜과 달리 **제어 연결**과 **데이터 연결**을 따로 사용합니다.

| 구분     | 역할                 | 포트              |
| ------ | ------------------ | --------------- |
| 제어 연결  | 로그인, 명령어 전달, 응답 처리 | TCP 21          |
| 데이터 연결 | 실제 파일 데이터 전송       | TCP 20 또는 동적 포트 |

```text
FTP = 제어 연결 + 데이터 연결
```

### FTP 종류

| 구분            | 전체 이름                            | 특징                    |
| ------------- | -------------------------------- | --------------------- |
| FTP           | File Transfer Protocol           | 기본 파일 전송 프로토콜, 암호화 없음 |
| Anonymous FTP | Anonymous File Transfer Protocol | 익명 사용자 접속 허용          |
| FTPS          | FTP Secure                       | FTP에 SSL/TLS 보안 기능 추가 |
| SFTP          | SSH File Transfer Protocol       | SSH 기반 파일 전송 프로토콜     |

#### FTP와 SFTP 주의점

FTP와 SFTP는 이름이 비슷하지만 구조가 다릅니다.

| 구분      | FTP      | SFTP            |
| ------- | -------- | --------------- |
| 기반 프로토콜 | FTP      | SSH             |
| 기본 포트   | TCP 21   | TCP 22          |
| 암호화     | 기본적으로 없음 | 지원              |
| 특징      | 파일 전송 전용 | SSH 기반 보안 파일 전송 |

```text
FTPS = FTP + SSL/TLS
SFTP = SSH 기반 파일 전송
```

### Active mode와 Passive mode

FTP는 데이터 연결을 만드는 방식에 따라 Active Mode와 Passive Mode로 나뉩니다.

### Active Mode

Active Mode는 클라이언트가 서버에게 자신의 포트를 알려주면, 서버가 클라이언트로 데이터 연결을 시도하는 방식입니다.

1. Active Mode의 특징

   * 클라이언트가 서버의 21번 포트로 제어 연결을 맺습니다.
   * 서버가 자신의 20번 포트를 사용하여 클라이언트에게 데이터 연결을 시도합니다.
   * 방화벽이나 NAT 환경에서는 클라이언트 쪽으로 들어오는 연결이 차단될 수 있습니다.

| 구분     | 내용                          |
| ------ | --------------------------- |
| 제어 연결  | 클라이언트 → 서버 TCP 21           |
| 데이터 연결 | 서버 TCP 20 → 클라이언트           |
| 특징     | 서버가 클라이언트에게 데이터 연결 시도       |
| 단점     | 클라이언트 방화벽/NAT 환경에서 문제 발생 가능 |

### Passive Mode

Passive Mode는 클라이언트가 서버에게 데이터 연결용 포트를 요청하고, 클라이언트가 서버의 해당 포트로 접속하는 방식입니다.

1. Passive Mode의 특징

   * 클라이언트가 서버의 21번 포트로 제어 연결을 맺습니다.
   * 서버가 데이터 연결에 사용할 포트를 알려줍니다.
   * 클라이언트가 서버의 해당 포트로 데이터 연결을 시도합니다.
   * NAT와 방화벽 환경에서 Active Mode보다 사용하기 쉽습니다.

| 구분     | 내용                    |
| ------ | --------------------- |
| 제어 연결  | 클라이언트 → 서버 TCP 21     |
| 데이터 연결 | 클라이언트 → 서버의 동적 포트     |
| 특징     | 클라이언트가 서버에게 데이터 연결 시도 |
| 장점     | 방화벽/NAT 환경에서 비교적 유리   |

### Active Mode와 Passive Mode 비교

| 구분         | Active Mode   | Passive Mode  |
| ---------- | ------------- | ------------- |
| 데이터 연결 시작  | 서버가 클라이언트로 접속 | 클라이언트가 서버로 접속 |
| 서버 데이터 포트  | TCP 20        | 서버가 알려준 동적 포트 |
| 방화벽/NAT 환경 | 불리할 수 있음      | 상대적으로 유리      |
| 핵심         | 서버 → 클라이언트    | 클라이언트 → 서버    |

#### 한 줄 정리

```text
FTP = 파일 전송 프로토콜
FTP 제어 연결 = 21번 포트
FTP 데이터 연결 = 20번 또는 동적 포트

Active Mode = 서버가 클라이언트로 데이터 연결
Passive Mode = 클라이언트가 서버로 데이터 연결
```

## 3) HTTP

HTTP는 HyperText Transfer Protocol의 약자로, 웹 브라우저와 웹 서버가 데이터를 주고받기 위해 사용하는 프로토콜입니다.

1. HTTP의 특징

   * 웹 문서를 전송하기 위한 응용 계층 프로토콜입니다.
   * 클라이언트와 서버 구조를 사용합니다.
   * 요청과 응답 방식으로 동작합니다.
   * 기본 포트는 TCP 80번입니다.
   * 상태를 유지하지 않는 Stateless 특성을 가집니다.
   * HTTPS는 HTTP에 TLS 보안 기능을 추가한 방식입니다.

| 구분       | 내용                          |
| -------- | --------------------------- |
| 전체 이름    | HyperText Transfer Protocol |
| 목적       | 웹 문서 및 웹 데이터 전송             |
| 기본 포트    | TCP 80                      |
| 보안 버전    | HTTPS                       |
| HTTPS 포트 | TCP 443                     |
| 특징       | 요청/응답 구조, Stateless         |

### HTTP 요청과 응답

HTTP는 클라이언트가 서버에 요청을 보내고, 서버가 응답을 반환하는 방식으로 동작합니다.

```text
클라이언트 → 서버 : HTTP Request
서버 → 클라이언트 : HTTP Response
```

| 구분       | 의미                  |
| -------- | ------------------- |
| Request  | 클라이언트가 서버에 보내는 요청   |
| Response | 서버가 클라이언트에게 반환하는 응답 |

### HTTP Method

| 메서드     | 역할              |
| ------- | --------------- |
| GET     | 데이터 조회          |
| POST    | 데이터 생성 또는 전송    |
| PUT     | 전체 데이터 수정       |
| PATCH   | 일부 데이터 수정       |
| DELETE  | 데이터 삭제          |
| HEAD    | 응답 헤더만 요청       |
| OPTIONS | 서버가 지원하는 메서드 확인 |

### HTTP 상태 코드

| 상태 코드 | 의미       | 설명             |
| ----- | -------- | -------------- |
| 1xx   | 정보 응답    | 요청을 처리 중       |
| 2xx   | 성공       | 요청이 정상 처리됨     |
| 3xx   | 리다이렉션    | 다른 위치로 이동 필요   |
| 4xx   | 클라이언트 오류 | 요청에 문제가 있음     |
| 5xx   | 서버 오류    | 서버 처리 중 문제가 발생 |

### HTTP와 HTTPS 비교

| 구분    | HTTP                        | HTTPS                              |
| ----- | --------------------------- | ---------------------------------- |
| 전체 이름 | HyperText Transfer Protocol | HyperText Transfer Protocol Secure |
| 기본 포트 | TCP 80                      | TCP 443                            |
| 암호화   | 없음                          | TLS 기반 암호화                         |
| 보안성   | 낮음                          | 높음                                 |
| 사용 예시 | 일반 웹 통신                     | 로그인, 결제, 개인정보 처리                   |

#### 한 줄 정리

```text
HTTP = 웹 브라우저와 웹 서버가 데이터를 주고받는 프로토콜
HTTPS = HTTP에 암호화를 추가한 보안 웹 프로토콜
```

## 4) DNS

DNS는 Domain Name System의 약자로, 사람이 이해하기 쉬운 도메인 이름을 컴퓨터가 사용하는 IP 주소로 변환하는 시스템입니다.

1. DNS의 특징

   * 도메인 이름을 IP 주소로 변환합니다.
   * 기본 포트는 53번입니다.
   * 일반적인 질의는 UDP 53번을 사용합니다.
   * 큰 응답이나 Zone Transfer는 TCP 53번을 사용할 수 있습니다.
   * 계층적인 분산 데이터베이스 구조를 가집니다.

| 구분            | 내용                 |
| ------------- | ------------------ |
| 전체 이름         | Domain Name System |
| 목적            | 도메인 이름을 IP 주소로 변환  |
| 기본 포트         | TCP/UDP 53         |
| 일반 질의         | UDP 53             |
| Zone Transfer | TCP 53             |
| 특징            | 계층적, 분산형 이름 해석 시스템 |

### DNS 동작 흐름

```text
사용자 입력: www.example.com
→ DNS 질의
→ IP 주소 응답
→ 해당 IP 주소로 서버 접속
```

### DNS 계층 구조

| 구분                       | 역할                           |
| ------------------------ | ---------------------------- |
| Root DNS Server          | 최상위 DNS 서버                   |
| TLD DNS Server           | .com, .net, .kr 등 최상위 도메인 관리 |
| Authoritative DNS Server | 특정 도메인의 실제 레코드 보유            |
| Recursive DNS Server     | 클라이언트 대신 DNS 질의를 수행          |

### DNS 주요 레코드

| 레코드   | 의미                  | 역할                        |
| ----- | ------------------- | ------------------------- |
| A     | Address Record      | 도메인을 IPv4 주소로 변환          |
| AAAA  | IPv6 Address Record | 도메인을 IPv6 주소로 변환          |
| CNAME | Canonical Name      | 도메인의 별칭 지정                |
| MX    | Mail Exchange       | 메일 서버 지정                  |
| NS    | Name Server         | 도메인의 네임서버 지정              |
| PTR   | Pointer             | IP 주소를 도메인 이름으로 역변환       |
| TXT   | Text                | 텍스트 정보 저장, SPF/DKIM 등에 사용 |
| SOA   | Start of Authority  | 도메인 영역의 기본 정보 저장          |

### 정방향 조회와 역방향 조회

| 구분     | 의미                 | 예시                            |
| ------ | ------------------ | ----------------------------- |
| 정방향 조회 | 도메인 이름을 IP 주소로 변환  | `example.com → 93.184.216.34` |
| 역방향 조회 | IP 주소를 도메인 이름으로 변환 | `93.184.216.34 → example.com` |

#### 한 줄 정리

```text
DNS = 도메인 이름을 IP 주소로 변환
A 레코드 = 도메인 → IPv4
AAAA 레코드 = 도메인 → IPv6
PTR 레코드 = IP → 도메인
MX 레코드 = 메일 서버 지정
```

## 5) SMTP, POP 및 IMAP

SMTP, POP3, IMAP은 이메일 송수신에 사용되는 대표적인 프로토콜입니다.

| 프로토콜 | 전체 이름                            | 역할           | 기본 포트   |
| ---- | -------------------------------- | ------------ | ------- |
| SMTP | Simple Mail Transfer Protocol    | 이메일 송신       | TCP 25  |
| POP3 | Post Office Protocol version 3   | 이메일 수신       | TCP 110 |
| IMAP | Internet Message Access Protocol | 이메일 수신 및 동기화 | TCP 143 |

### SMTP

SMTP는 이메일을 보내기 위한 프로토콜입니다.

1. SMTP의 특징

   * Simple Mail Transfer Protocol의 약자입니다.
   * 이메일을 송신할 때 사용합니다.
   * 메일 클라이언트에서 메일 서버로 보낼 때 사용됩니다.
   * 메일 서버 간 메일 전달에도 사용됩니다.
   * 기본 포트는 TCP 25번입니다.

| 구분    | 내용                                 |
| ----- | ---------------------------------- |
| 전체 이름 | Simple Mail Transfer Protocol      |
| 역할    | 이메일 송신                             |
| 기본 포트 | TCP 25                             |
| 특징    | 메일 클라이언트 → 메일 서버, 메일 서버 → 메일 서버 전송 |

### POP3

POP3는 메일 서버에 있는 이메일을 클라이언트로 내려받기 위한 프로토콜입니다.

1. POP3의 특징

   * Post Office Protocol version 3의 약자입니다.
   * 이메일 수신에 사용합니다.
   * 메일을 클라이언트로 다운로드하는 방식입니다.
   * 기본적으로 다운로드 후 서버에서 메일을 삭제하는 방식으로 사용될 수 있습니다.
   * 여러 기기에서 동기화하기에는 IMAP보다 불리합니다.

| 구분    | 내용                             |
| ----- | ------------------------------ |
| 전체 이름 | Post Office Protocol version 3 |
| 역할    | 이메일 수신                         |
| 기본 포트 | TCP 110                        |
| 특징    | 메일을 클라이언트로 내려받는 방식             |

### IMAP

IMAP은 메일을 서버에 보관하면서 여러 기기에서 동기화하여 확인할 수 있는 프로토콜입니다.

1. IMAP의 특징

   * Internet Message Access Protocol의 약자입니다.
   * 이메일 수신과 동기화에 사용합니다.
   * 메일을 서버에 보관합니다.
   * 여러 기기에서 같은 메일함 상태를 유지하기 좋습니다.
   * 기본 포트는 TCP 143번입니다.

| 구분    | 내용                               |
| ----- | -------------------------------- |
| 전체 이름 | Internet Message Access Protocol |
| 역할    | 이메일 수신 및 동기화                     |
| 기본 포트 | TCP 143                          |
| 특징    | 서버 보관, 여러 기기 동기화에 유리             |

### SMTP, POP3, IMAP 비교

| 구분        | SMTP    | POP3      | IMAP          |
| --------- | ------- | --------- | ------------- |
| 역할        | 이메일 송신  | 이메일 수신    | 이메일 수신 및 동기화  |
| 기본 포트     | TCP 25  | TCP 110   | TCP 143       |
| 메일 저장 위치  | 전송 중심   | 클라이언트 중심  | 서버 중심         |
| 여러 기기 동기화 | 해당 없음   | 불리함       | 유리함           |
| 핵심        | 보낼 때 사용 | 내려받을 때 사용 | 서버와 동기화할 때 사용 |

### 보안 포트

| 프로토콜            | 보안 방식             | 포트      |
| --------------- | ----------------- | ------- |
| SMTPS           | SMTP over SSL/TLS | TCP 465 |
| SMTP Submission | 메일 제출             | TCP 587 |
| POP3S           | POP3 over SSL/TLS | TCP 995 |
| IMAPS           | IMAP over SSL/TLS | TCP 993 |

#### 한 줄 정리

```text
SMTP = 메일 송신
POP3 = 메일을 내려받아 수신
IMAP = 서버에 메일을 보관하며 동기화
```

## 6) SNMP

SNMP는 Simple Network Management Protocol의 약자로, 네트워크 장비를 관리하고 상태를 모니터링하기 위한 프로토콜입니다.

1. SNMP의 특징

   * 네트워크 장비 관리에 사용됩니다.
   * 라우터, 스위치, 서버, 프린터 등의 상태를 확인할 수 있습니다.
   * 장비의 CPU 사용률, 메모리 사용률, 트래픽, 장애 상태 등을 수집할 수 있습니다.
   * UDP 기반으로 동작합니다.
   * Manager와 Agent 구조를 사용합니다.

| 구분       | 내용                                 |
| -------- | ---------------------------------- |
| 전체 이름    | Simple Network Management Protocol |
| 목적       | 네트워크 장비 관리 및 모니터링                  |
| 전송 계층    | UDP                                |
| 요청/응답 포트 | UDP 161                            |
| Trap 포트  | UDP 162                            |
| 관리 대상    | 라우터, 스위치, 서버, 프린터 등                |

### SNMP 구성 요소

| 구성 요소   | 역할                                        |
| ------- | ----------------------------------------- |
| Manager | 관리 시스템, 장비 상태를 조회하고 제어                    |
| Agent   | 관리 대상 장비에 설치되어 정보를 제공                     |
| MIB     | Management Information Base, 관리 정보 데이터베이스 |
| OID     | Object Identifier, 관리 항목을 식별하는 번호         |
| Trap    | Agent가 Manager에게 이벤트를 알리는 메시지             |

### SNMP 동작 방식

| 동작       | 의미                           |
| -------- | ---------------------------- |
| Get      | Manager가 Agent에게 정보를 요청      |
| Set      | Manager가 Agent의 값을 변경        |
| Response | Agent가 요청에 응답                |
| Trap     | Agent가 장애나 이벤트를 Manager에게 알림 |

### SNMP 버전

| 버전      | 특징                 |
| ------- | ------------------ |
| SNMPv1  | 초기 버전, 보안 기능 약함    |
| SNMPv2c | 성능 개선, 커뮤니티 문자열 기반 |
| SNMPv3  | 인증과 암호화 지원, 보안 강화  |

#### 한 줄 정리

```text
SNMP = 네트워크 장비 상태를 관리하고 모니터링하는 프로토콜
Manager = 관리 주체
Agent = 관리 대상 장비
MIB/OID = 관리 정보 식별 체계
Trap = 장애나 이벤트 알림
```

## 7) DHCP

DHCP는 Dynamic Host Configuration Protocol의 약자로, 네트워크에 접속한 장비에게 IP 주소와 네트워크 설정 정보를 자동으로 할당하는 프로토콜입니다.

1. DHCP의 특징

   * IP 주소를 자동으로 할당합니다.
   * 서브넷 마스크, 기본 게이트웨이, DNS 서버 정보를 함께 제공할 수 있습니다.
   * UDP 기반으로 동작합니다.
   * DHCP 서버는 UDP 67번 포트를 사용합니다.
   * DHCP 클라이언트는 UDP 68번 포트를 사용합니다.
   * IP 주소를 일정 기간 임대하는 Lease 방식을 사용합니다.

| 구분       | 내용                                  |
| -------- | ----------------------------------- |
| 전체 이름    | Dynamic Host Configuration Protocol |
| 목적       | IP 주소와 네트워크 설정 자동 할당                |
| 전송 계층    | UDP                                 |
| 서버 포트    | UDP 67                              |
| 클라이언트 포트 | UDP 68                              |
| 제공 정보    | IP 주소, 서브넷 마스크, 게이트웨이, DNS 서버       |
| 특징       | 임대, Lease 방식 사용                     |

### DHCP 제공 정보

| 제공 정보      | 설명              |
| ---------- | --------------- |
| IP 주소      | 클라이언트에게 할당되는 주소 |
| 서브넷 마스크    | 네트워크 범위 구분      |
| 기본 게이트웨이   | 외부 네트워크로 나가는 경로 |
| DNS 서버     | 도메인 이름 해석 서버    |
| Lease Time | IP 주소 임대 시간     |

### DHCP 동작 과정

DHCP는 보통 DORA 과정으로 동작합니다.

| 단계 | 메시지         | 의미                     |
| -- | ----------- | ---------------------- |
| 1  | Discover    | 클라이언트가 DHCP 서버를 찾음     |
| 2  | Offer       | DHCP 서버가 사용 가능한 IP를 제안 |
| 3  | Request     | 클라이언트가 제안받은 IP 사용 요청   |
| 4  | Acknowledge | DHCP 서버가 IP 할당 승인      |

```text
DORA = Discover → Offer → Request → Acknowledge
```

### DHCP 동작 흐름

```text
1. 클라이언트 → 전체 네트워크 : DHCP Discover
2. 서버 → 클라이언트 : DHCP Offer
3. 클라이언트 → 서버 : DHCP Request
4. 서버 → 클라이언트 : DHCP Acknowledge
```

### DHCP 장점

1. IP 주소 관리가 편리합니다.
2. 사용자가 직접 IP를 설정하지 않아도 됩니다.
3. IP 충돌을 줄일 수 있습니다.
4. 네트워크 설정을 중앙에서 관리할 수 있습니다.
5. 이동이 많은 환경에서 효율적입니다.

#### 한 줄 정리

```text
DHCP = IP 주소와 네트워크 설정을 자동으로 할당하는 프로토콜
DORA = Discover → Offer → Request → Acknowledge
서버 포트 = UDP 67
클라이언트 포트 = UDP 68
```

## 전체 핵심 요약 표

| 프로토콜   | 전체 이름                               | 역할         | 기본 포트        | 핵심 암기        |
| ------ | ----------------------------------- | ---------- | ------------ | ------------ |
| Telnet | Teletype Network                    | 원격 접속      | TCP 23       | 암호화 없음       |
| SSH    | Secure Shell                        | 보안 원격 접속   | TCP 22       | 암호화 지원       |
| FTP    | File Transfer Protocol              | 파일 전송      | TCP 20, 21   | 제어/데이터 연결 분리 |
| HTTP   | HyperText Transfer Protocol         | 웹 문서 전송    | TCP 80       | 웹 통신         |
| HTTPS  | HyperText Transfer Protocol Secure  | 보안 웹 문서 전송 | TCP 443      | TLS 암호화      |
| DNS    | Domain Name System                  | 도메인 이름 변환  | TCP/UDP 53   | 도메인 → IP     |
| SMTP   | Simple Mail Transfer Protocol       | 이메일 송신     | TCP 25       | 메일 보내기       |
| POP3   | Post Office Protocol version 3      | 이메일 수신     | TCP 110      | 메일 내려받기      |
| IMAP   | Internet Message Access Protocol    | 이메일 수신/동기화 | TCP 143      | 서버 보관/동기화    |
| SNMP   | Simple Network Management Protocol  | 네트워크 관리    | UDP 161, 162 | 장비 모니터링      |
| DHCP   | Dynamic Host Configuration Protocol | IP 자동 할당   | UDP 67, 68   | DORA         |

## 시험용 최종 암기

```text
Telnet = 원격 접속, TCP 23, 암호화 없음
SSH = 보안 원격 접속, TCP 22
FTP = 파일 전송, TCP 20/21
HTTP = 웹, TCP 80
HTTPS = 보안 웹, TCP 443
DNS = 도메인 → IP, TCP/UDP 53
SMTP = 메일 송신, TCP 25
POP3 = 메일 수신, TCP 110
IMAP = 메일 동기화, TCP 143
SNMP = 네트워크 관리, UDP 161/162
DHCP = IP 자동 할당, UDP 67/68
```

```text
FTP Active = 서버가 클라이언트로 데이터 연결
FTP Passive = 클라이언트가 서버로 데이터 연결

DNS A = IPv4 주소
DNS AAAA = IPv6 주소
DNS MX = 메일 서버
DNS CNAME = 별칭

DHCP DORA = Discover → Offer → Request → Acknowledge
```
