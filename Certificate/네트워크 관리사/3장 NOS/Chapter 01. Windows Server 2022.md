---

# Windows Server 2022 핵심 요약본

## 0. 전체 구조 한눈에 보기

Windows Server 2022는 Microsoft의 서버용 운영체제이며, 기업 환경에서 다음과 같은 역할을 수행한다.

| 구분         | 핵심 역할                |
| ---------- | -------------------- |
| 사용자/계정 관리  | Active Directory     |
| 웹 서버 운영    | IIS                  |
| 가상화        | Hyper-V              |
| 파일/스토리지 관리 | NTFS, ReFS           |
| 서버 기능 설치   | 역할 기반 설치             |
| 원격/중앙 관리   | Windows Admin Center |
| 자동화/관리 명령  | PowerShell           |

네트워크관리사 2급 관점에서는 **“서버 운영체제가 네트워크 안에서 어떤 서비스를 제공하는가”**를 중심으로 이해하면 된다.

---

# 1. Windows Server 개요

## 1.1 Windows Server란?

**Windows Server**는 Microsoft에서 제공하는 서버용 운영체제다.
일반 PC용 Windows와 달리, 여러 사용자가 접속하는 네트워크 환경에서 다음 기능을 제공한다.

* 사용자 계정 및 권한 관리
* 파일 공유
* 프린터 공유
* 웹 서버 운영
* 도메인 관리
* 가상 머신 운영
* 원격 서버 관리
* 보안 정책 적용
* 기업 내부 네트워크 서비스 제공

즉, Windows Server는 단순히 “서버에 설치하는 Windows”가 아니라, **조직의 네트워크 자원과 사용자를 중앙에서 관리하기 위한 서버 플랫폼**이다.

Microsoft는 Windows Server 2022의 핵심 개선 방향을 **보안, Azure 하이브리드 통합 및 관리, 애플리케이션 플랫폼**으로 설명한다. ([Microsoft Learn][1])

---

## 1.2 일반 Windows와 Windows Server의 차이

| 구분    | 일반 Windows        | Windows Server                     |
| ----- | ----------------- | ---------------------------------- |
| 목적    | 개인용 PC 사용         | 기업/조직 서버 운영                        |
| 사용자 수 | 주로 1인 또는 소수       | 다수 사용자 접속                          |
| 주요 기능 | 문서 작업, 브라우징, 앱 실행 | AD, DNS, DHCP, IIS, 파일 서버, Hyper-V |
| 관리 방식 | 로컬 중심             | 중앙 관리 중심                           |
| 안정성   | 일반 사용자 환경 기준      | 장시간 운영, 서비스 지속성 중시                 |
| 보안 정책 | 개인 PC 보안 중심       | 조직 단위 보안 정책 적용                     |
| 서버 역할 | 제한적               | 역할 기반으로 다양한 서버 기능 제공               |

---

## 1.3 Windows Server 2016 vs 2022 비교

| 구분       | Windows Server 2016                                                  | Windows Server 2022                                   |
| -------- | -------------------------------------------------------------------- | ----------------------------------------------------- |
| 출시 방향    | 클라우드와 가상화 기반 강화                                                      | 보안, 하이브리드 클라우드, 최신 애플리케이션 플랫폼 강화                      |
| 보안       | Shielded VM, Credential Guard, Just Enough Administration 등 보안 기능 강화 | Secured-core server, TLS 1.3, SMB 보안 강화 등 다층 보안 강화    |
| 가상화      | Hyper-V 개선, Shielded VM 도입                                           | Hyper-V 유지 및 보안/관리 기능 강화                              |
| 컨테이너     | Windows 컨테이너, Nano Server 도입                                         | 컨테이너 이미지 크기 개선, Kubernetes 관련 개선                      |
| 네트워크     | SDN, 네트워크 컨트롤러, Hyper-V 스위치 개선                                       | SMB over QUIC, Azure Extended Network 등 하이브리드 네트워크 강화 |
| 관리       | Server Manager, PowerShell 중심                                        | Windows Admin Center와 Azure Arc 연계 관리 강화              |
| 스토리지     | Storage Spaces Direct, Storage Replica 등 도입                          | Storage Migration Service, SMB 압축 등 개선                |
| 클라우드 연계  | 초기 하이브리드 기능 강화                                                       | Azure와의 하이브리드 통합이 더 강조됨                               |
| 시험 관점 핵심 | 가상화, 보안, 컨테이너, SDN 도입                                                | 보안 강화, Azure 연계, WAC, 최신 서버 관리                        |

Windows Server 2016은 Hyper-V, SDN, Failover Cluster, Storage Spaces Direct, Shielded VM 등 기업 인프라 기능이 대폭 강화된 버전이다. ([Microsoft Learn][2])
Windows Server 2022는 2019 기반 위에서 보안, Azure 하이브리드 통합, 애플리케이션 플랫폼을 중심으로 개선된 버전이다. ([Microsoft Learn][1])

---

## 1.4 Windows Server 2022 주요 신기능 핵심 요소

| 핵심 요소                     | 설명                                 |
| ------------------------- | ---------------------------------- |
| Secured-core server       | 펌웨어, 하드웨어, OS 수준의 보안을 결합한 서버 보안 기능 |
| TLS 1.3 지원                | 네트워크 통신 암호화 보안 강화                  |
| SMB 보안 개선                 | 파일 공유 프로토콜 SMB의 암호화 및 보안 강화        |
| SMB over QUIC             | VPN 없이도 안전한 SMB 파일 접근을 지원하는 기능     |
| Azure Arc 연계              | 온프레미스 서버를 Azure에서 관리할 수 있도록 연결     |
| Windows Admin Center 강화   | 웹 기반 서버 관리 도구로 로컬/원격 서버 관리         |
| 컨테이너 개선                   | Windows 컨테이너 성능 및 이미지 크기 개선        |
| Storage Migration Service | 기존 서버의 파일/스토리지 데이터를 새 서버로 이전 지원    |
| SMB 압축                    | 파일 전송 시 압축을 통해 전송 효율 향상            |

시험에서는 세부 구현보다 **“Windows Server 2022는 보안과 하이브리드 관리가 강화된 서버 OS”**라는 방향을 잡는 것이 중요하다.

---

# 2. 에디션별 분류

Windows Server 2022는 주로 다음 에디션으로 구분된다.

| 에디션                       | 용도                       |
| ------------------------- | ------------------------ |
| Essentials                | 소규모 조직용                  |
| Standard                  | 일반적인 서버 환경               |
| Datacenter                | 대규모 가상화, 데이터센터 환경        |
| Datacenter: Azure Edition | Azure 및 하이브리드 클라우드 특화 환경 |

Microsoft의 Windows Server 2022 수명 주기 문서에 따르면 Windows Server 2022는 **Datacenter, Datacenter: Azure Edition, Essentials, Standard** 에디션에 적용된다. ([Microsoft Learn][3])

---

## 2.1 Essentials

**Essentials**는 소규모 기업을 위한 에디션이다.

### 특징

* 소규모 사용자 환경에 적합
* 간단한 서버 운영 목적
* 대규모 가상화나 고급 데이터센터 기능에는 적합하지 않음
* 소규모 파일 공유, 기본 서버 기능 등에 활용 가능

### 시험 포인트

Essentials는 **소규모 조직용**으로 기억하면 된다.

---

## 2.2 Standard

**Standard**는 일반적인 기업 서버 환경에서 가장 기본적으로 사용하는 에디션이다.

### 특징

* 일반 서버 업무에 적합
* AD, DNS, DHCP, IIS, 파일 서버, Hyper-V 등 주요 역할 사용 가능
* 제한적인 가상화 권한 제공
* 중소규모 서버 환경에 적합

### 사용 예시

* 사내 도메인 컨트롤러
* 파일 서버
* 웹 서버
* 소규모 Hyper-V 서버
* 업무용 애플리케이션 서버

### 시험 포인트

Standard는 **일반 서버 운영용 에디션**이다.

---

## 2.3 Datacenter

**Datacenter**는 대규모 가상화와 데이터센터 환경에 적합한 에디션이다.

Microsoft Evaluation Center 설명에 따르면 Datacenter는 가장 완전한 에디션이며, Shielded Virtual Machines, Storage Spaces Direct, Software-Defined Networking 같은 Datacenter 전용 기능과 무제한 서버 가상화를 포함한다. ([Microsoft][4])

### 특징

* 대규모 가상화 환경에 적합
* 무제한 가상 머신 운영 권한
* 고급 스토리지 기능 지원
* SDN 기능 지원
* 데이터센터, 클라우드, 대규모 서버 인프라에 적합

### 시험 포인트

Datacenter는 **대규모 가상화, 데이터센터, 고급 기능**으로 기억하면 된다.

---

## 2.4 에디션 비교 요약

| 구분      | Essentials | Standard | Datacenter |
| ------- | ---------- | -------- | ---------- |
| 대상      | 소규모 조직     | 일반 기업    | 대규모 데이터센터  |
| 가상화     | 제한적        | 제한적      | 무제한 수준     |
| 고급 스토리지 | 제한적        | 일부 가능    | 강력         |
| SDN     | 제한적        | 제한적      | 지원         |
| 사용 난이도  | 낮음         | 보통       | 높음         |
| 시험 키워드  | 소규모        | 일반 서버    | 대규모 가상화    |

---

# 3. Active Directory, AD

## 3.1 Active Directory란?

**Active Directory, AD**는 Windows Server에서 제공하는 **디렉터리 서비스**다.

디렉터리 서비스란 네트워크 안의 사용자, 컴퓨터, 프린터, 서버, 공유 폴더 같은 자원 정보를 저장하고 관리하는 서비스다.

Microsoft는 AD DS를 네트워크 개체에 대한 정보를 저장하고, 권한 있는 사용자와 관리자가 이 데이터를 사용할 수 있게 하는 디렉터리 서비스라고 설명한다. ([Microsoft Learn][5])

---

## 3.2 AD를 쉽게 이해하기

회사에 직원이 100명 있다고 가정하자.

각 직원의 PC마다 계정을 따로 만들면 다음 문제가 생긴다.

* 직원 계정을 PC마다 따로 만들어야 함
* 비밀번호 변경을 PC마다 해야 함
* 퇴사자 계정을 모든 PC에서 삭제해야 함
* 부서별 접근 권한 관리가 어려움
* 공유 폴더 권한 관리가 복잡함

AD를 사용하면 중앙 서버에서 다음을 관리할 수 있다.

* 사용자 계정
* 비밀번호
* 컴퓨터 계정
* 그룹
* 부서별 권한
* 로그인 정책
* 보안 정책
* 공유 자원 접근 권한

즉, AD는 **기업 네트워크의 사용자와 자원을 중앙에서 관리하는 시스템**이다.

---

## 3.3 AD DS

정확한 명칭은 **Active Directory Domain Services, AD DS**다.

AD DS는 다음 정보를 저장한다.

| 저장 대상 | 예시               |
| ----- | ---------------- |
| 사용자   | 홍길동, 김철수 계정      |
| 컴퓨터   | 사무실 PC, 서버       |
| 그룹    | 인사팀, 개발팀, 관리자 그룹 |
| 프린터   | 네트워크 프린터         |
| 공유 폴더 | 부서별 파일 서버        |
| 정책    | 암호 정책, 로그인 정책    |

---

## 3.4 도메인

**도메인**은 AD에서 중앙 관리되는 논리적 네트워크 단위다.

예를 들어 회사 도메인이 다음과 같다고 하자.

```text
company.local
```

이 도메인 안에 사용자, 컴퓨터, 서버, 그룹이 소속된다.

```text
company.local
 ├─ 사용자: user01, user02
 ├─ 컴퓨터: PC-001, PC-002
 ├─ 서버: FILE-SERVER, WEB-SERVER
 └─ 그룹: 개발팀, 인사팀, 관리자
```

도메인을 사용하면 사용자는 도메인 계정 하나로 여러 네트워크 자원에 접근할 수 있다.

---

## 3.5 도메인 컨트롤러, DC

**도메인 컨트롤러, Domain Controller, DC**는 AD DS가 설치된 서버다.

DC의 역할은 다음과 같다.

* 사용자 로그인 인증
* 사용자 계정 관리
* 그룹 정책 적용
* 도메인 정보 저장
* 권한 검증
* 네트워크 자원 접근 제어

사용자가 회사 PC에 로그인할 때 DC가 다음을 확인한다.

```text
이 사용자가 존재하는가?
비밀번호가 맞는가?
이 사용자가 이 PC에 로그인할 권한이 있는가?
이 사용자가 공유 폴더에 접근할 수 있는가?
```

---

## 3.6 OU

**OU, Organizational Unit**은 AD 안에서 사용자나 컴퓨터를 조직적으로 묶는 단위다.

예시:

```text
company.local
 ├─ OU=개발팀
 │   ├─ user_dev01
 │   └─ PC-DEV01
 ├─ OU=인사팀
 │   ├─ user_hr01
 │   └─ PC-HR01
 └─ OU=관리자
     └─ admin01
```

OU를 사용하는 이유는 다음과 같다.

* 부서별 계정 관리
* 부서별 정책 적용
* 관리 권한 위임
* 컴퓨터 그룹 관리

---

## 3.7 그룹 정책, GPO

**GPO, Group Policy Object**는 도메인 사용자나 컴퓨터에 정책을 적용하는 기능이다.

예시:

| 정책         | 설명                 |
| ---------- | ------------------ |
| 암호 길이 제한   | 비밀번호 최소 8자리 이상     |
| 화면 잠금      | 10분 미사용 시 자동 잠금    |
| USB 차단     | 보안상 USB 저장장치 사용 금지 |
| 제어판 차단     | 일반 사용자의 설정 변경 제한   |
| 특정 프로그램 배포 | 조직 전체에 프로그램 자동 설치  |

시험에서는 GPO를 **도메인 환경에서 사용자/컴퓨터 정책을 중앙 적용하는 기능**으로 기억하면 된다.

---

## 3.8 AD 시험 포인트

| 키워드      | 의미                               |
| -------- | -------------------------------- |
| AD       | 네트워크 자원과 사용자를 중앙 관리하는 디렉터리 서비스   |
| AD DS    | Active Directory Domain Services |
| 도메인      | 중앙 관리되는 논리적 네트워크 범위              |
| DC       | 도메인 컨트롤러, 인증 담당 서버               |
| OU       | 조직 단위                            |
| GPO      | 그룹 정책                            |
| LDAP     | 디렉터리 서비스 접근 프로토콜                 |
| Kerberos | 도메인 인증에 사용되는 인증 프로토콜             |
| DNS      | AD 도메인 환경에서 필수적으로 사용됨            |

---

# 4. Hyper-V

## 4.1 Hyper-V란?

**Hyper-V**는 Microsoft의 가상화 기술이다.

하나의 물리 서버 위에서 여러 개의 가상 머신, VM을 실행할 수 있게 해준다.

Microsoft는 Hyper-V를 Windows Server와 Windows에 기본 제공되는 엔터프라이즈급 하이퍼바이저 기술로 설명하며, 여러 운영체제의 가상 머신을 생성, 관리, 실행할 수 있다고 설명한다. ([Microsoft Learn][6])

---

## 4.2 가상화란?

가상화는 하나의 물리 자원을 여러 개의 논리 자원처럼 나누어 사용하는 기술이다.

예를 들어 물리 서버 1대가 있다.

```text
물리 서버 1대
 ├─ VM 1: Windows Server - AD 서버
 ├─ VM 2: Windows Server - 파일 서버
 ├─ VM 3: Linux - 웹 서버
 └─ VM 4: Linux - DB 서버
```

이렇게 하나의 서버에서 여러 서버를 운영할 수 있다.

---

## 4.3 Hyper-V의 주요 구성 요소

| 구성 요소    | 설명                            |
| -------- | ----------------------------- |
| 호스트      | Hyper-V가 설치된 물리 서버            |
| 게스트      | VM 안에서 실행되는 운영체제              |
| VM       | 가상 머신                         |
| VHD/VHDX | 가상 하드디스크 파일                   |
| 가상 스위치   | VM이 네트워크에 연결되도록 하는 가상 네트워크 장치 |
| 체크포인트    | VM 상태를 특정 시점으로 저장하는 기능        |

---

## 4.4 Hyper-V 가상 스위치 종류

| 종류       | 설명                            |
| -------- | ----------------------------- |
| External | VM이 외부 네트워크와 통신 가능            |
| Internal | 호스트와 VM 간 통신 가능, 외부 네트워크 불가   |
| Private  | VM 간 통신만 가능, 호스트 및 외부 네트워크 불가 |

### External Switch

```text
VM ↔ 호스트 NIC ↔ 외부 네트워크
```

외부 인터넷이나 사내망과 연결할 때 사용한다.

### Internal Switch

```text
VM ↔ 호스트
```

호스트와 VM 사이의 통신만 필요할 때 사용한다.

### Private Switch

```text
VM ↔ VM
```

VM끼리만 통신해야 할 때 사용한다.

---

## 4.5 Hyper-V의 장점

| 장점        | 설명                      |
| --------- | ----------------------- |
| 서버 통합     | 여러 서버를 하나의 물리 서버에 통합 가능 |
| 비용 절감     | 물리 장비 수 감소              |
| 테스트 환경 구축 | 빠르게 VM 생성/삭제 가능         |
| 격리성       | VM 간 독립된 환경 구성          |
| 백업/복구 용이  | VM 단위 백업 가능             |
| 유연한 확장    | 필요 시 VM 추가 가능           |

---

## 4.6 시험 포인트

| 키워드             | 의미               |
| --------------- | ---------------- |
| Hyper-V         | Microsoft 가상화 기술 |
| VM              | 가상 머신            |
| Host            | 물리 서버            |
| Guest           | VM 내부 OS         |
| VHD/VHDX        | 가상 디스크           |
| External Switch | 외부 네트워크 연결       |
| Internal Switch | 호스트-VM 통신        |
| Private Switch  | VM 간 통신          |
| Checkpoint      | 특정 시점 저장         |

---

# 5. IIS, Internet Information Services

## 5.1 IIS란?

**IIS, Internet Information Services**는 Windows Server에서 제공하는 웹 서버 기능이다.

IIS는 웹 사이트, 웹 애플리케이션, FTP 서비스 등을 제공할 수 있다.

Microsoft IIS 공식 사이트는 IIS를 Windows Server용 유연하고 안전하며 관리 가능한 웹 서버라고 설명한다. ([IIS][7])

---

## 5.2 웹 서버란?

웹 서버는 클라이언트의 HTTP/HTTPS 요청을 받아 웹 페이지나 웹 애플리케이션 응답을 제공하는 서버다.

```text
사용자 브라우저 → HTTP/HTTPS 요청 → IIS 서버 → 웹 페이지 응답
```

예시:

```text
https://example.com 접속
→ IIS가 요청을 받음
→ index.html 또는 ASP.NET 애플리케이션 실행
→ 사용자에게 결과 반환
```

---

## 5.3 IIS의 주요 기능

| 기능          | 설명                                |
| ----------- | --------------------------------- |
| 웹 사이트 호스팅   | HTML, CSS, JavaScript 기반 웹 페이지 제공 |
| 웹 애플리케이션 실행 | ASP.NET 등 서버 애플리케이션 실행            |
| FTP 서버      | 파일 업로드/다운로드 서비스 제공                |
| HTTPS 지원    | SSL/TLS 인증서를 이용한 암호화 통신           |
| 인증 기능       | Windows 인증, 기본 인증 등               |
| 로그 관리       | 접속 기록, 오류 기록 저장                   |
| 애플리케이션 풀    | 웹 애플리케이션 실행 환경 분리                 |
| 바인딩         | IP, 포트, 호스트 이름 연결                 |

---

## 5.4 IIS 기본 개념

### 사이트

웹 서비스를 제공하는 단위다.

예:

```text
Default Web Site
company.com
intranet.company.local
```

### 바인딩

사이트가 어떤 주소와 포트로 동작할지 설정하는 것이다.

| 요소     | 예시                                        |
| ------ | ----------------------------------------- |
| IP 주소  | 192.168.0.10                              |
| 포트     | 80, 443                                   |
| 호스트 이름 | [www.company.com](http://www.company.com) |
| 프로토콜   | HTTP, HTTPS                               |

### 애플리케이션 풀

웹 애플리케이션을 실행하는 프로세스 격리 단위다.

여러 웹 사이트를 하나의 IIS에서 운영할 때, 애플리케이션 풀을 분리하면 하나의 사이트 장애가 다른 사이트에 영향을 덜 준다.

---

## 5.5 IIS 기본 포트

| 프로토콜  |  포트 | 설명       |
| ----- | --: | -------- |
| HTTP  |  80 | 일반 웹 통신  |
| HTTPS | 443 | 암호화 웹 통신 |
| FTP   |  21 | 파일 전송    |

---

## 5.6 시험 포인트

| 키워드      | 의미                  |
| -------- | ------------------- |
| IIS      | Windows Server 웹 서버 |
| HTTP     | 80번 포트              |
| HTTPS    | 443번 포트             |
| FTP      | 21번 포트              |
| 바인딩      | IP, 포트, 호스트명 연결     |
| 애플리케이션 풀 | 웹 앱 실행 환경 분리        |
| SSL/TLS  | HTTPS 암호화           |

---

# 6. Windows PowerShell

## 6.1 PowerShell이란?

**Windows PowerShell**은 Windows Server를 명령어 기반으로 관리하고 자동화할 수 있는 셸 및 스크립트 환경이다.

기존 CMD보다 강력하며, Windows Server 관리에 자주 사용된다.

---

## 6.2 PowerShell의 특징

| 특징    | 설명                       |
| ----- | ------------------------ |
| 객체 기반 | 단순 문자열이 아니라 객체를 다룸       |
| 자동화   | 반복 작업을 스크립트로 자동화 가능      |
| 서버 관리 | 역할 설치, 서비스 관리, 사용자 관리 가능 |
| 원격 관리 | 다른 서버에 명령 실행 가능          |
| 파이프라인 | 명령 결과를 다른 명령으로 전달 가능     |

---

## 6.3 CMD와 PowerShell 차이

| 구분     | CMD          | PowerShell            |
| ------ | ------------ | --------------------- |
| 성격     | 전통적인 명령 프롬프트 | 관리 자동화 셸              |
| 데이터 처리 | 문자열 중심       | 객체 중심                 |
| 스크립트   | Batch 파일     | PowerShell Script     |
| 확장성    | 제한적          | 강력                    |
| 서버 관리  | 제한적          | Windows Server 관리에 적합 |

---

## 6.4 PowerShell 명령 구조

PowerShell 명령은 보통 다음 구조를 가진다.

```powershell
동사-명사
```

예시:

```powershell
Get-Service
Get-Process
Restart-Service
New-Item
Remove-Item
Install-WindowsFeature
```

| 명령                     | 의미                      |
| ---------------------- | ----------------------- |
| Get-Service            | 서비스 목록 조회               |
| Get-Process            | 프로세스 목록 조회              |
| Restart-Service        | 서비스 재시작                 |
| New-Item               | 파일/폴더 생성                |
| Remove-Item            | 파일/폴더 삭제                |
| Install-WindowsFeature | Windows Server 역할/기능 설치 |

---

## 6.5 자주 쓰는 PowerShell 예시

### 서비스 목록 확인

```powershell
Get-Service
```

### 특정 서비스 재시작

```powershell
Restart-Service -Name Spooler
```

### IP 설정 확인

```powershell
Get-NetIPConfiguration
```

### 네트워크 어댑터 확인

```powershell
Get-NetAdapter
```

### Windows 기능 목록 확인

```powershell
Get-WindowsFeature
```

### IIS 설치

```powershell
Install-WindowsFeature Web-Server
```

### Hyper-V 설치

```powershell
Install-WindowsFeature Hyper-V
```

---

## 6.6 시험 포인트

| 키워드                    | 의미                   |
| ---------------------- | -------------------- |
| PowerShell             | Windows 관리 자동화 셸     |
| Cmdlet                 | PowerShell 명령        |
| 동사-명사 구조               | Get-Service 같은 명령 형식 |
| Pipeline               | 명령 결과 전달             |
| Install-WindowsFeature | 서버 역할/기능 설치          |
| 원격 관리                  | 다른 서버 관리 가능          |

---

# 7. 파일 시스템, NTFS와 ReFS

## 7.1 파일 시스템이란?

파일 시스템은 저장장치에 파일과 폴더를 저장하고 관리하는 방식이다.

Windows Server에서 중요한 파일 시스템은 다음 두 가지다.

| 파일 시스템 | 설명                          |
| ------ | --------------------------- |
| NTFS   | Windows의 기본 파일 시스템          |
| ReFS   | 안정성과 대용량 스토리지에 초점을 둔 파일 시스템 |

---

## 7.2 NTFS

**NTFS, New Technology File System**는 Windows 계열 운영체제의 기본 파일 시스템이다.

Microsoft는 NTFS가 보안 설명자, 암호화, 디스크 할당량, 풍부한 메타데이터 등 고급 기능을 제공하는 현대 Windows 운영체제의 기본 파일 시스템이라고 설명한다. ([Microsoft Learn][8])

---

## 7.3 NTFS 주요 기능

| 기능        | 설명                    |
| --------- | --------------------- |
| 파일 권한     | 사용자/그룹별 접근 권한 설정      |
| EFS 암호화   | 파일 단위 암호화             |
| 디스크 할당량   | 사용자별 사용 가능한 디스크 용량 제한 |
| 압축        | 파일/폴더 압축              |
| 감사        | 파일 접근 기록 추적           |
| 대용량 볼륨 지원 | 큰 저장공간 관리 가능          |
| 저널링       | 장애 발생 시 복구 가능성 향상     |

---

## 7.4 NTFS 권한

NTFS의 핵심은 **권한 관리**다.

| 권한                   | 설명        |
| -------------------- | --------- |
| Full Control         | 모든 권한     |
| Modify               | 수정, 삭제 가능 |
| Read & Execute       | 읽기 및 실행   |
| List Folder Contents | 폴더 내용 보기  |
| Read                 | 읽기        |
| Write                | 쓰기        |

예를 들어 개발팀 공유 폴더가 있을 때 다음처럼 설정할 수 있다.

```text
개발팀: Modify
인사팀: 접근 불가
관리자: Full Control
```

---

## 7.5 공유 권한과 NTFS 권한 차이

Windows 파일 서버에서 권한은 크게 두 가지가 있다.

| 구분      | 적용 위치        | 설명                |
| ------- | ------------ | ----------------- |
| 공유 권한   | 네트워크 공유 접근 시 | 네트워크를 통해 접근할 때 적용 |
| NTFS 권한 | 파일 시스템 자체    | 로컬/네트워크 접근 모두 적용  |

둘 다 적용되는 경우에는 보통 **더 제한적인 권한이 최종 적용**된다.

예시:

| 공유 권한                 | NTFS 권한           | 최종 권한  |
| --------------------- | ----------------- | ------ |
| Everyone Read         | User Full Control | Read   |
| Everyone Full Control | User Read         | Read   |
| Everyone Modify       | User Modify       | Modify |

---

## 7.6 ReFS

**ReFS, Resilient File System**는 Microsoft가 개발한 최신 파일 시스템이다.

Microsoft는 ReFS가 데이터 손상으로부터 보호하고, 대규모 스토리지 환경을 지원하며, Windows Server 주요 기술과 통합된다고 설명한다. ([Microsoft Learn][9])

---

## 7.7 ReFS 주요 특징

| 특징                | 설명                               |
| ----------------- | -------------------------------- |
| 무결성 강화            | 데이터 손상 감지 및 보호                   |
| 대용량 스토리지 적합       | 대규모 데이터 저장 환경에 적합                |
| 복원력               | 장애나 손상에 대한 대응력 강화                |
| Storage Spaces 연계 | Windows Server 스토리지 기능과 함께 사용 가능 |
| 가상화 환경 적합         | Hyper-V 가상 디스크 저장소에 적합           |

---

## 7.8 NTFS vs ReFS 비교

| 구분         | NTFS              | ReFS                             |
| ---------- | ----------------- | -------------------------------- |
| 기본 용도      | 일반 Windows 파일 시스템 | 대용량/고가용성 스토리지                    |
| 권한 관리      | 강력                | 지원                               |
| 파일 암호화 EFS | 지원                | 제한적 또는 미지원 기능 존재                 |
| 압축         | 지원                | 제한적                              |
| 데이터 무결성    | 기본 제공             | 더 강력                             |
| 대용량 스토리지   | 가능                | 더 적합                             |
| 부팅 볼륨      | 가능                | 일반적으로 부팅 볼륨 용도 아님                |
| 사용 예       | 일반 서버, 파일 서버      | 대용량 데이터, Hyper-V, Storage Spaces |

---

## 7.9 시험 포인트

| 키워드     | 의미                         |
| ------- | -------------------------- |
| NTFS    | Windows 기본 파일 시스템          |
| ReFS    | 복원력과 대용량 스토리지 중심 파일 시스템    |
| EFS     | 파일 암호화                     |
| Quota   | 디스크 사용량 제한                 |
| 공유 권한   | 네트워크 공유 접근 권한              |
| NTFS 권한 | 파일 시스템 권한                  |
| 최종 권한   | 공유 권한과 NTFS 권한 중 더 제한적인 권한 |

---

# 8. 역할 기반 설치

## 8.1 역할 기반 설치란?

Windows Server는 필요한 기능을 **역할(Role)** 과 **기능(Feature)** 단위로 설치한다.

예를 들어 웹 서버로 사용하려면 IIS 역할을 설치하고, 도메인 컨트롤러로 사용하려면 AD DS 역할을 설치한다.

Microsoft 문서에 따르면 Server Manager에서 **Manage → Add Roles and Features**를 선택해 역할 및 기능 추가 마법사를 실행할 수 있다. ([Microsoft Learn][10])

---

## 8.2 Role과 Feature 차이

| 구분      | 설명                  | 예시                                                 |
| ------- | ------------------- | -------------------------------------------------- |
| Role    | 서버가 수행하는 주된 역할      | AD DS, DNS, DHCP, IIS, Hyper-V                     |
| Feature | 역할을 보조하거나 OS 기능을 확장 | .NET Framework, Telnet Client, Failover Clustering |

쉽게 말하면 다음과 같다.

```text
Role = 서버의 직업
Feature = 직업 수행을 도와주는 기능
```

---

## 8.3 주요 서버 역할

| 역할                          | 설명                  |
| --------------------------- | ------------------- |
| AD DS                       | 도메인 사용자/컴퓨터 중앙 관리   |
| DNS Server                  | 도메인 이름을 IP 주소로 변환   |
| DHCP Server                 | 클라이언트에게 IP 주소 자동 할당 |
| Web Server, IIS             | 웹 사이트 및 웹 앱 제공      |
| File and Storage Services   | 파일 공유 및 저장소 관리      |
| Hyper-V                     | 가상 머신 운영            |
| Remote Desktop Services     | 원격 데스크톱 환경 제공       |
| Print and Document Services | 프린터 공유 및 관리         |

---

## 8.4 역할 기반 설치 흐름

일반적인 설치 흐름은 다음과 같다.

```text
Server Manager 실행
→ Manage
→ Add Roles and Features
→ Installation Type 선택
→ 대상 서버 선택
→ Server Roles 선택
→ Features 선택
→ 확인
→ 설치
```

---

## 8.5 설치 방식

| 방식            | 설명                       |
| ------------- | ------------------------ |
| GUI 설치        | Server Manager에서 마법사로 설치 |
| PowerShell 설치 | 명령어로 설치                  |
| 원격 설치         | 다른 서버에 역할 설치 가능          |

예시:

```powershell
Install-WindowsFeature Web-Server
```

위 명령은 IIS 웹 서버 역할을 설치한다.

---

## 8.6 시험 포인트

| 키워드                    | 의미               |
| ---------------------- | ---------------- |
| Role                   | 서버의 주된 기능        |
| Feature                | 보조 기능            |
| Server Manager         | GUI 기반 서버 관리 도구  |
| Add Roles and Features | 역할/기능 추가 마법사     |
| Install-WindowsFeature | PowerShell 기반 설치 |
| AD DS                  | 도메인 서비스          |
| DNS                    | 이름 해석            |
| DHCP                   | IP 자동 할당         |
| IIS                    | 웹 서버             |
| Hyper-V                | 가상화              |

---

# 9. Windows Admin Center, WAC

## 9.1 Windows Admin Center란?

**Windows Admin Center, WAC**는 웹 브라우저 기반의 Windows Server 관리 도구다.

Microsoft는 Windows Admin Center를 물리, 가상, 온프레미스, Azure, 호스팅 환경 어디서나 실행되는 Windows Server용 원격 관리 도구라고 설명한다. ([Microsoft Learn][11])

---

## 9.2 WAC가 필요한 이유

기존 Windows Server 관리는 다음 도구들이 분산되어 있었다.

* Server Manager
* MMC
* Event Viewer
* Services
* Computer Management
* Hyper-V Manager
* Failover Cluster Manager
* PowerShell

WAC는 이런 관리 기능을 웹 기반으로 통합해 제공한다.

즉, 브라우저에서 서버 상태와 기능을 관리할 수 있다.

---

## 9.3 WAC 주요 기능

| 기능            | 설명                           |
| ------------- | ---------------------------- |
| 서버 상태 확인      | CPU, 메모리, 디스크, 네트워크 확인       |
| 이벤트 로그 관리     | 시스템/보안/응용 로그 조회              |
| 서비스 관리        | 서비스 시작, 중지, 재시작              |
| 파일 관리         | 서버 파일 탐색                     |
| 레지스트리 관리      | 레지스트리 확인 및 수정                |
| 방화벽 관리        | Windows Defender Firewall 관리 |
| 업데이트 관리       | Windows Update 확인            |
| 역할/기능 관리      | 설치된 역할 및 기능 확인               |
| Hyper-V 관리    | VM 생성 및 관리                   |
| 스토리지 관리       | 디스크, 볼륨, 파일 공유 관리            |
| PowerShell 실행 | 웹에서 PowerShell 명령 실행         |

---

## 9.4 WAC와 Server Manager 차이

| 구분       | Server Manager    | Windows Admin Center |
| -------- | ----------------- | -------------------- |
| 형태       | Windows 내장 GUI 도구 | 웹 기반 관리 도구           |
| 접근 방식    | 서버 또는 관리 PC에서 실행  | 브라우저로 접속             |
| 관리 범위    | 전통적 서버 관리         | 통합 원격 관리             |
| 확장성      | 제한적               | 확장 기능 지원             |
| Azure 연계 | 제한적               | Azure 하이브리드 관리와 연계   |
| 시험 키워드   | 역할/기능 설치          | 웹 기반 중앙 관리           |

---

## 9.5 WAC 시험 포인트

| 키워드                  | 의미                   |
| -------------------- | -------------------- |
| WAC                  | Windows Admin Center |
| 웹 기반 관리              | 브라우저에서 서버 관리         |
| 원격 관리                | 여러 서버를 중앙에서 관리       |
| 추가 비용 없음             | Microsoft 제공 관리 도구   |
| Azure 연계             | 하이브리드 서버 관리 가능       |
| Server Manager 대체/보완 | 기존 관리 도구를 통합         |

---

# 10. 전체 핵심 암기표

## 10.1 Windows Server 2022 핵심 키워드

| 키워드                 | 핵심 의미             |
| ------------------- | ----------------- |
| Windows Server 2022 | Microsoft 서버 운영체제 |
| AD DS               | 사용자/컴퓨터/자원 중앙 관리  |
| Domain              | 중앙 관리되는 네트워크 단위   |
| DC                  | 도메인 컨트롤러, 인증 서버   |
| GPO                 | 그룹 정책             |
| Hyper-V             | 가상화 기술            |
| VM                  | 가상 머신             |
| IIS                 | 웹 서버              |
| PowerShell          | 관리 자동화 셸          |
| NTFS                | Windows 기본 파일 시스템 |
| ReFS                | 복원력 있는 파일 시스템     |
| Role                | 서버 역할             |
| Feature             | 보조 기능             |
| WAC                 | 웹 기반 서버 관리 도구     |

---

## 10.2 포트 암기

| 서비스         |   포트 | 설명          |
| ----------- | ---: | ----------- |
| HTTP        |   80 | 웹           |
| HTTPS       |  443 | 보안 웹        |
| FTP         |   21 | 파일 전송       |
| DNS         |   53 | 이름 해석       |
| DHCP Server |   67 | IP 할당 서버    |
| DHCP Client |   68 | IP 할당 클라이언트 |
| LDAP        |  389 | 디렉터리 서비스    |
| LDAPS       |  636 | 보안 LDAP     |
| RDP         | 3389 | 원격 데스크톱     |
| SMB         |  445 | 파일 공유       |
| Kerberos    |   88 | AD 인증       |

---

## 10.3 자주 헷갈리는 개념 정리

### AD와 DC 차이

| 구분 | 설명               |
| -- | ---------------- |
| AD | 디렉터리 서비스 개념      |
| DC | AD DS가 설치된 실제 서버 |

AD는 시스템이고, DC는 그 시스템을 운영하는 서버다.

---

### Role과 Feature 차이

| 구분      | 설명              |
| ------- | --------------- |
| Role    | 서버가 담당하는 주된 역할  |
| Feature | 역할 수행을 돕는 추가 기능 |

예:

```text
IIS = Role
.NET Framework = Feature
```

---

### NTFS와 공유 권한 차이

| 구분      | 설명               |
| ------- | ---------------- |
| 공유 권한   | 네트워크로 접근할 때 적용   |
| NTFS 권한 | 파일 시스템 자체에 적용    |
| 최종 권한   | 둘 중 더 제한적인 권한 적용 |

---

### IIS와 웹 브라우저 차이

| 구분   | 설명                   |
| ---- | -------------------- |
| IIS  | 웹 페이지를 제공하는 서버       |
| 브라우저 | 웹 페이지를 요청하고 보는 클라이언트 |

---

### Hyper-V와 VM 차이

| 구분      | 설명                      |
| ------- | ----------------------- |
| Hyper-V | 가상화를 제공하는 기술            |
| VM      | Hyper-V 위에서 실행되는 가상 컴퓨터 |

---

# 11. 네트워크관리사 2급 관점 최종 정리

Windows Server 2022는 네트워크관리사 2급에서 다음 관점으로 정리하면 된다.

## 11.1 서버 운영체제 관점

Windows Server는 기업 네트워크에서 여러 서버 역할을 수행한다.

```text
Windows Server
 ├─ AD DS: 사용자/컴퓨터 중앙 관리
 ├─ DNS: 이름 해석
 ├─ DHCP: IP 자동 할당
 ├─ IIS: 웹 서버
 ├─ File Server: 파일 공유
 ├─ Hyper-V: 가상화
 └─ WAC: 웹 기반 서버 관리
```

---

## 11.2 관리 관점

서버 관리는 다음 방식으로 수행할 수 있다.

| 방식             | 설명              |
| -------------- | --------------- |
| Server Manager | GUI 기반 역할/기능 관리 |
| PowerShell     | 명령어 기반 자동화 관리   |
| WAC            | 웹 기반 원격 통합 관리   |
| MMC            | 개별 관리 콘솔        |

---

## 11.3 보안 관점

Windows Server에서 중요한 보안 요소는 다음과 같다.

| 보안 요소            | 설명              |
| ---------------- | --------------- |
| AD 인증            | 도메인 사용자 인증      |
| GPO              | 중앙 보안 정책 적용     |
| NTFS 권한          | 파일 접근 제어        |
| 방화벽              | 네트워크 접근 제어      |
| HTTPS            | 웹 통신 암호화        |
| SMB 보안           | 파일 공유 보안        |
| PowerShell 권한 관리 | 관리자 권한 기반 서버 관리 |

---

# 12. 최종 암기 문장

아래 문장만 외워도 전체 흐름을 잡을 수 있다.

> Windows Server 2022는 기업 네트워크에서 사용자, 컴퓨터, 파일, 웹, 가상화, 보안 정책을 중앙 관리하기 위한 Microsoft 서버 운영체제다.

> AD DS는 사용자와 컴퓨터를 중앙 관리하는 디렉터리 서비스이고, DC는 AD DS가 설치되어 인증을 담당하는 서버다.

> Hyper-V는 하나의 물리 서버에서 여러 가상 머신을 실행하게 해주는 Microsoft 가상화 기술이다.

> IIS는 Windows Server에서 웹 사이트와 웹 애플리케이션을 제공하는 웹 서버 역할이다.

> PowerShell은 Windows Server를 명령어와 스크립트로 관리하고 자동화하는 도구다.

> NTFS는 Windows의 기본 파일 시스템이고, ReFS는 대용량 스토리지와 데이터 복원력에 초점을 둔 파일 시스템이다.

> 역할 기반 설치는 서버에 필요한 기능을 Role과 Feature 단위로 선택 설치하는 방식이다.

> Windows Admin Center는 브라우저에서 Windows Server를 원격으로 관리하는 웹 기반 관리 도구다.
