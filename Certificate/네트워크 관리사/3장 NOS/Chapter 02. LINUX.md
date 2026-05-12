# LINUX 핵심 요약본

네트워크관리사 2급 관점에서 Linux는 단순히 “무료 운영체제”가 아니라, **서버 운영, 네트워크 관리, 파일 권한 관리, 프로세스 관리, 서비스 운영**을 이해하기 위한 핵심 운영체제입니다.

특히 시험에서는 다음 흐름으로 이해하면 좋습니다.

```text
Linux
 ├─ 운영체제 개요
 ├─ Linux 구조
 ├─ 파일 시스템과 권한
 ├─ 기본 명령어
 ├─ vi 편집기
 └─ 서버 활용
     ├─ 프로세스
     ├─ 데몬
     ├─ RPM
     └─ SAMBA
```

---

# 1. LINUX 개요

## 1.1 Linux란?

**Linux**는 Unix 계열의 오픈소스 운영체제입니다.

정확히 말하면 Linux는 운영체제 전체를 의미하기도 하지만, 엄밀하게는 **커널 Kernel**을 의미합니다.
일반적으로 우리가 사용하는 Ubuntu, CentOS, Rocky Linux, Debian, Fedora 같은 것은 Linux 커널에 여러 프로그램과 패키지 관리 도구를 결합한 **Linux 배포판**입니다.

```text
Linux Kernel
+ Shell
+ 기본 명령어
+ 패키지 관리자
+ 파일 시스템
+ 네트워크 도구
+ 시스템 서비스
= Linux 배포판
```

---

## 1.2 Linux의 일반적 특징

| 특징          | 설명                                               |
| ----------- | ------------------------------------------------ |
| 오픈소스        | 소스 코드가 공개되어 자유롭게 사용, 수정, 배포 가능                   |
| 다중 사용자      | 여러 사용자가 동시에 시스템 사용 가능                            |
| 다중 작업       | 여러 프로그램을 동시에 실행 가능                               |
| 높은 안정성      | 장시간 서버 운영에 적합                                    |
| 강력한 네트워크 기능 | 서버, 라우터, 방화벽, 파일 공유 등 네트워크 기능에 강함                |
| 다양한 배포판     | Ubuntu, Debian, Fedora, Rocky Linux 등 목적별 배포판 존재 |
| 낮은 비용       | 라이선스 비용 부담이 적음                                   |
| 보안성         | 권한 기반 구조와 빠른 보안 패치 생태계                           |
| CLI 중심 관리   | 명령어 기반 관리에 강함                                    |
| 이식성         | 다양한 CPU, 장비, 서버 환경에서 사용 가능                       |

---

## 1.3 Linux의 기술적 특징

| 기술적 특징     | 설명                                         |
| ---------- | ------------------------------------------ |
| 계층적 파일 시스템 | `/` 루트 디렉터리를 기준으로 모든 파일과 장치를 관리            |
| 모든 것은 파일   | 일반 파일, 디렉터리, 장치, 프로세스 정보도 파일처럼 취급          |
| POSIX 호환   | Unix 계열 표준 인터페이스를 따름                       |
| 강력한 권한 관리  | 사용자, 그룹, 기타 사용자 기준 권한 관리                   |
| Shell 제공   | bash, zsh 등의 셸을 통해 명령 실행                   |
| 파이프와 리다이렉션 | 명령어 결과를 연결하거나 파일로 저장 가능                    |
| 프로세스 기반 동작 | 실행 중인 프로그램을 프로세스로 관리                       |
| 데몬 서비스     | 백그라운드 서비스 형태로 서버 기능 제공                     |
| 패키지 관리     | RPM, YUM, DNF, APT 등을 통해 소프트웨어 설치/관리       |
| 네트워크 서버 기능 | SSH, FTP, DNS, DHCP, Web, SAMBA 등 서버 운영 가능 |

---

## 1.4 Linux가 서버에서 많이 사용되는 이유

Linux는 서버 환경에서 매우 많이 사용됩니다. 이유는 다음과 같습니다.

| 이유       | 설명                      |
| -------- | ----------------------- |
| 안정성      | 장시간 재부팅 없이 운영 가능        |
| 보안성      | 사용자/그룹/권한 체계가 명확함       |
| 비용       | 무료 또는 저비용으로 서버 구축 가능    |
| 자동화      | Shell Script로 관리 자동화 가능 |
| 원격 관리    | SSH를 통한 원격 접속이 편리함      |
| 네트워크 친화적 | 서버, 라우터, 방화벽, NAS 등에 적합 |
| 경량성      | 불필요한 GUI 없이 서버 운영 가능    |

---

## 1.5 Linux 배포판

Linux 배포판은 Linux 커널에 여러 도구와 패키지 시스템을 결합한 운영체제입니다.

| 계열         | 배포판                                          | 특징                    |
| ---------- | -------------------------------------------- | --------------------- |
| Debian 계열  | Debian, Ubuntu                               | 사용 편의성, APT 패키지 관리    |
| Red Hat 계열 | RHEL, CentOS, Rocky Linux, AlmaLinux, Fedora | 서버 환경에서 많이 사용, RPM 기반 |
| Arch 계열    | Arch Linux, Manjaro                          | 최신 패키지 중심             |
| SUSE 계열    | openSUSE, SUSE Linux Enterprise              | 기업 환경 사용              |
| 경량 배포판     | Alpine Linux                                 | 컨테이너 환경에서 자주 사용       |

네트워크관리사 2급에서는 특히 **Red Hat 계열의 RPM, YUM** 개념이 자주 연결됩니다.

---

# 2. LINUX 구조

## 2.1 Linux 전체 구조

Linux는 크게 다음 구조로 이해할 수 있습니다.

```text
사용자 User
   ↓
응용 프로그램 Application
   ↓
Shell
   ↓
Kernel
   ↓
Hardware
```

| 구성 요소    | 설명                            |
| -------- | ----------------------------- |
| 사용자      | 명령을 입력하거나 프로그램을 실행하는 사람       |
| 응용 프로그램  | vi, gcc, ssh, apache, samba 등 |
| Shell    | 사용자 명령을 해석해 커널에 전달            |
| Kernel   | 하드웨어와 시스템 자원을 관리하는 핵심         |
| Hardware | CPU, 메모리, 디스크, 네트워크 카드 등      |

---

## 2.2 Kernel

**Kernel**은 운영체제의 핵심입니다.

커널은 하드웨어와 소프트웨어 사이에서 다음 역할을 수행합니다.

| 역할        | 설명                              |
| --------- | ------------------------------- |
| 프로세스 관리   | 실행 중인 프로그램 관리                   |
| 메모리 관리    | RAM 할당 및 회수                     |
| 파일 시스템 관리 | 파일 읽기/쓰기 처리                     |
| 장치 관리     | 디스크, 키보드, 네트워크 카드 등 제어          |
| 네트워크 관리   | TCP/IP 통신 처리                    |
| 시스템 호출 제공 | 프로그램이 커널 기능을 사용할 수 있도록 인터페이스 제공 |

쉽게 말하면 커널은 **운영체제의 엔진**입니다.

---

## 2.3 Shell

**Shell**은 사용자와 커널 사이의 명령어 해석기입니다.

사용자가 명령어를 입력하면 Shell이 이를 해석해서 커널에 전달합니다.

```text
사용자 명령어 입력
→ Shell이 해석
→ Kernel에 요청
→ Kernel이 하드웨어/시스템 자원 처리
→ 결과 출력
```

대표적인 Shell은 다음과 같습니다.

| Shell | 설명                       |
| ----- | ------------------------ |
| sh    | 기본 Unix Shell            |
| bash  | Linux에서 가장 많이 사용하는 Shell |
| zsh   | 확장 기능이 많은 Shell          |
| csh   | C 언어 스타일 Shell           |
| ksh   | Korn Shell               |

시험에서는 보통 **bash**를 대표 Shell로 기억하면 됩니다.

---

## 2.4 파일 시스템 계층 구조

Linux는 Windows처럼 `C:\`, `D:\` 드라이브 중심이 아닙니다.

Linux는 `/`를 최상위로 하는 하나의 트리 구조를 사용합니다.

```text
/
├── bin
├── boot
├── dev
├── etc
├── home
├── lib
├── media
├── mnt
├── opt
├── proc
├── root
├── sbin
├── tmp
├── usr
└── var
```

---

## 2.5 주요 디렉터리

| 디렉터리     | 설명                      |
| -------- | ----------------------- |
| `/`      | 최상위 루트 디렉터리             |
| `/bin`   | 기본 사용자 명령어 저장           |
| `/sbin`  | 시스템 관리 명령어 저장           |
| `/boot`  | 부팅 관련 파일 저장             |
| `/dev`   | 장치 파일 저장                |
| `/etc`   | 시스템 설정 파일 저장            |
| `/home`  | 일반 사용자 홈 디렉터리           |
| `/root`  | root 관리자 홈 디렉터리         |
| `/lib`   | 공유 라이브러리 저장             |
| `/usr`   | 응용 프로그램, 라이브러리, 문서 저장   |
| `/var`   | 로그, 메일, 캐시 등 변동 데이터 저장  |
| `/tmp`   | 임시 파일 저장                |
| `/proc`  | 커널/프로세스 정보 제공 가상 파일 시스템 |
| `/mnt`   | 임시 마운트 지점               |
| `/media` | USB, CD-ROM 등 자동 마운트 지점 |
| `/opt`   | 추가 소프트웨어 설치 위치          |

---

## 2.6 Linux 구조 시험 포인트

| 키워드     | 의미             |
| ------- | -------------- |
| Kernel  | 운영체제 핵심        |
| Shell   | 명령어 해석기        |
| bash    | Linux 대표 Shell |
| `/`     | 루트 디렉터리        |
| `/etc`  | 설정 파일          |
| `/home` | 일반 사용자 홈       |
| `/root` | 관리자 홈          |
| `/var`  | 로그 등 변동 파일     |
| `/dev`  | 장치 파일          |
| `/proc` | 프로세스/커널 정보     |

---

# 3. LINUX 파일 시스템

## 3.1 Linux 파일 시스템 개념

Linux에서는 거의 모든 것을 파일로 취급합니다.

```text
일반 문서 → 파일
디렉터리 → 파일
하드디스크 → 장치 파일
터미널 → 장치 파일
프로세스 정보 → 가상 파일
네트워크 설정 → 설정 파일
```

이런 구조 때문에 Linux 관리는 곧 **파일 관리와 권한 관리**라고 볼 수 있습니다.

---

## 3.2 Linux 파일 종류

`ls -l` 명령어를 실행하면 맨 앞 글자로 파일 종류를 확인할 수 있습니다.

```bash
-rw-r--r--  1 user user  1000 May 12 memo.txt
drwxr-xr-x  2 user user  4096 May 12 docs
lrwxrwxrwx  1 user user    10 May 12 link -> memo.txt
```

| 맨 앞 문자 | 파일 종류    | 설명                |
| ------ | -------- | ----------------- |
| `-`    | 일반 파일    | 텍스트, 실행 파일, 이미지 등 |
| `d`    | 디렉터리     | 폴더                |
| `l`    | 심볼릭 링크   | 바로가기              |
| `b`    | 블록 장치 파일 | 디스크 장치            |
| `c`    | 문자 장치 파일 | 키보드, 터미널 등        |
| `p`    | 파이프      | 프로세스 간 통신         |
| `s`    | 소켓       | 네트워크 통신용 파일       |

---

## 3.3 Linux 대표 파일 시스템

| 파일 시스템 | 설명                        |
| ------ | ------------------------- |
| ext2   | 오래된 Linux 파일 시스템, 저널링 없음  |
| ext3   | ext2에 저널링 기능 추가           |
| ext4   | 현재 많이 사용하는 Linux 파일 시스템   |
| XFS    | 대용량 파일과 고성능 서버 환경에 적합     |
| Btrfs  | 스냅샷, 체크섬 등 고급 기능 제공       |
| swap   | 메모리 부족 시 디스크 일부를 메모리처럼 사용 |

시험에서는 보통 **ext4**, **XFS**, **swap** 정도를 중심으로 보면 됩니다.

---

## 3.4 파일 시스템 관리

## 3.4.1 마운트 Mount

Linux에서 디스크나 파티션을 사용하려면 특정 디렉터리에 연결해야 합니다.
이 작업을 **마운트 mount**라고 합니다.

```text
디스크 장치 /dev/sdb1
→ /data 디렉터리에 연결
→ /data 경로로 디스크 사용 가능
```

### 수동 마운트

```bash
mount /dev/sdb1 /data
```

### 마운트 해제

```bash
umount /data
```

주의할 점은 명령어가 `unmount`가 아니라 **umount**입니다.

---

## 3.4.2 자동 마운트 설정 `/etc/fstab`

서버가 재부팅되어도 자동으로 디스크가 마운트되도록 하려면 `/etc/fstab` 파일을 설정합니다.

```text
/dev/sdb1   /data   ext4   defaults   0   0
```

| 항목          | 의미         |
| ----------- | ---------- |
| `/dev/sdb1` | 장치명        |
| `/data`     | 마운트 위치     |
| `ext4`      | 파일 시스템 타입  |
| `defaults`  | 기본 옵션      |
| `0`         | dump 백업 여부 |
| `0`         | fsck 검사 순서 |

---

## 3.4.3 디스크 사용량 확인

### 전체 파일 시스템 사용량

```bash
df -h
```

| 옵션   | 의미                |
| ---- | ----------------- |
| `df` | 파일 시스템 디스크 사용량 확인 |
| `-h` | 사람이 읽기 쉬운 단위로 출력  |

### 특정 디렉터리 용량 확인

```bash
du -sh /var/log
```

| 옵션   | 의미               |
| ---- | ---------------- |
| `du` | 디렉터리/파일 사용량 확인   |
| `-s` | 합계만 출력           |
| `-h` | 사람이 읽기 쉬운 단위로 출력 |

---

## 3.4.4 파일 시스템 검사

```bash
fsck /dev/sdb1
```

`fsck`는 파일 시스템 오류를 검사하고 복구하는 명령어입니다.

주의할 점은 일반적으로 **마운트 해제된 상태에서 실행**해야 안전합니다.

---

# 3.5 명령 프롬프트

Linux의 명령 프롬프트는 현재 사용자, 호스트, 현재 위치, 권한 상태를 보여줍니다.

예시:

```bash
user@server:~$
root@server:/etc#
```

| 기호  | 의미             |
| --- | -------------- |
| `$` | 일반 사용자         |
| `#` | root 관리자 사용자   |
| `~` | 현재 사용자의 홈 디렉터리 |
| `/` | 루트 디렉터리        |

---

## 3.6 절대 경로와 상대 경로

## 절대 경로

루트 `/`부터 시작하는 경로입니다.

```bash
/home/user/test.txt
/etc/hosts
/var/log/messages
```

## 상대 경로

현재 위치를 기준으로 이동하는 경로입니다.

```bash
./test.txt
../backup
docs/readme.txt
```

| 표현   | 의미            |
| ---- | ------------- |
| `.`  | 현재 디렉터리       |
| `..` | 상위 디렉터리       |
| `~`  | 현재 사용자 홈 디렉터리 |
| `/`  | 루트 디렉터리       |

---

# 3.7 권한 관리

Linux 권한은 시험에서 매우 중요합니다.

Linux는 파일과 디렉터리에 대해 다음 세 대상을 기준으로 권한을 부여합니다.

| 대상              | 설명              |
| --------------- | --------------- |
| 사용자 User, Owner | 파일 소유자          |
| 그룹 Group        | 파일 소유 그룹        |
| 기타 Others       | 소유자와 그룹이 아닌 사용자 |

---

## 3.7.1 권한 표시 방식

```bash
-rwxr-xr--
```

이를 나누면 다음과 같습니다.

```text
-   rwx   r-x   r--
│    │     │     │
│    │     │     └─ 기타 사용자 권한
│    │     └─────── 그룹 권한
│    └───────────── 소유자 권한
└────────────────── 파일 종류
```

| 권한  | 의미         | 숫자 |
| --- | ---------- | -: |
| `r` | 읽기 read    |  4 |
| `w` | 쓰기 write   |  2 |
| `x` | 실행 execute |  1 |
| `-` | 권한 없음      |  0 |

---

## 3.7.2 숫자 권한 계산

권한은 숫자로 표현할 수 있습니다.

| 권한    | 계산    | 숫자 |
| ----- | ----- | -: |
| `r--` | 4     |  4 |
| `rw-` | 4+2   |  6 |
| `r-x` | 4+1   |  5 |
| `rwx` | 4+2+1 |  7 |
| `---` | 0     |  0 |

예시:

```bash
chmod 755 script.sh
```

`755`의 의미:

```text
7 = 소유자 rwx
5 = 그룹 r-x
5 = 기타 r-x
```

결과:

```text
-rwxr-xr-x
```

---

## 3.7.3 자주 나오는 권한

|  숫자 | 문자        | 의미                     |
| --: | --------- | ---------------------- |
| 777 | rwxrwxrwx | 모든 사용자 모든 권한           |
| 755 | rwxr-xr-x | 소유자는 모든 권한, 나머지는 읽기/실행 |
| 700 | rwx------ | 소유자만 모든 권한             |
| 644 | rw-r--r-- | 소유자는 읽기/쓰기, 나머지는 읽기    |
| 600 | rw------- | 소유자만 읽기/쓰기             |
| 400 | r-------- | 소유자만 읽기                |

---

## 3.7.4 chmod

파일 권한을 변경하는 명령어입니다.

```bash
chmod 755 file.sh
chmod 644 memo.txt
```

기호 방식도 가능합니다.

```bash
chmod u+x script.sh
chmod g-w file.txt
chmod o+r file.txt
```

| 기호  | 의미        |
| --- | --------- |
| `u` | 사용자 owner |
| `g` | 그룹        |
| `o` | 기타 사용자    |
| `a` | 전체        |
| `+` | 권한 추가     |
| `-` | 권한 제거     |
| `=` | 권한 지정     |

---

## 3.7.5 chown

파일 소유자와 그룹을 변경하는 명령어입니다.

```bash
chown user file.txt
chown user:group file.txt
```

예시:

```bash
chown yuchan:yuchan memo.txt
```

---

## 3.7.6 chgrp

파일의 그룹만 변경하는 명령어입니다.

```bash
chgrp developers memo.txt
```

---

## 3.7.7 디렉터리 권한의 의미

파일 권한과 디렉터리 권한은 의미가 약간 다릅니다.

| 권한 | 파일에서 의미  | 디렉터리에서 의미       |
| -- | -------- | --------------- |
| r  | 파일 내용 읽기 | 디렉터리 목록 보기      |
| w  | 파일 내용 수정 | 디렉터리 내 파일 생성/삭제 |
| x  | 파일 실행    | 디렉터리 안으로 접근 가능  |

디렉터리에서 `x` 권한이 없으면 해당 디렉터리 안으로 진입할 수 없습니다.

---

# 4. LINUX 명령어

Linux 명령어는 시험에서 직접적인 암기 포인트입니다.

크게 다음처럼 나눌 수 있습니다.

```text
Linux 명령어
 ├─ 파일/디렉터리 명령어
 ├─ 시스템 관리 명령어
 ├─ 텍스트 처리 명령어
 └─ 네트워크 명령어
```

---

# 4.1 파일/디렉터리 기본 명령어

| 명령어      | 설명               |
| -------- | ---------------- |
| `pwd`    | 현재 디렉터리 확인       |
| `ls`     | 파일/디렉터리 목록 확인    |
| `cd`     | 디렉터리 이동          |
| `mkdir`  | 디렉터리 생성          |
| `rmdir`  | 빈 디렉터리 삭제        |
| `touch`  | 빈 파일 생성 또는 시간 변경 |
| `cp`     | 파일/디렉터리 복사       |
| `mv`     | 파일 이동 또는 이름 변경   |
| `rm`     | 파일/디렉터리 삭제       |
| `find`   | 파일 검색            |
| `locate` | 파일 빠른 검색         |

---

## 4.1.1 ls

```bash
ls
ls -l
ls -a
ls -al
```

| 옵션   | 설명              |
| ---- | --------------- |
| `-l` | 자세히 보기          |
| `-a` | 숨김 파일 포함        |
| `-h` | 사람이 읽기 쉬운 단위    |
| `-R` | 하위 디렉터리까지 재귀 출력 |

---

## 4.1.2 cp

```bash
cp file1 file2
cp -r dir1 dir2
```

| 옵션   | 설명           |
| ---- | ------------ |
| `-r` | 디렉터리 재귀 복사   |
| `-p` | 권한, 시간 정보 유지 |
| `-i` | 덮어쓰기 전 확인    |

---

## 4.1.3 rm

```bash
rm file.txt
rm -r directory
rm -rf directory
```

| 옵션   | 설명         |
| ---- | ---------- |
| `-r` | 디렉터리 재귀 삭제 |
| `-f` | 강제 삭제      |
| `-i` | 삭제 전 확인    |

주의:

```bash
rm -rf /
```

이런 명령은 시스템을 망가뜨릴 수 있으므로 절대 실행하면 안 됩니다.

---

# 4.2 시스템 관리 명령어

| 명령어         | 설명             |
| ----------- | -------------- |
| `uname`     | 시스템 정보 확인      |
| `hostname`  | 호스트명 확인/변경     |
| `whoami`    | 현재 사용자 확인      |
| `id`        | UID, GID 확인    |
| `who`       | 로그인 사용자 확인     |
| `w`         | 로그인 사용자와 작업 확인 |
| `date`      | 날짜/시간 확인       |
| `cal`       | 달력 출력          |
| `uptime`    | 시스템 가동 시간 확인   |
| `df`        | 파일 시스템 사용량 확인  |
| `du`        | 디렉터리 사용량 확인    |
| `free`      | 메모리 사용량 확인     |
| `top`       | 실시간 프로세스 확인    |
| `ps`        | 프로세스 목록 확인     |
| `kill`      | 프로세스 종료        |
| `shutdown`  | 시스템 종료         |
| `reboot`    | 시스템 재부팅        |
| `systemctl` | systemd 서비스 관리 |

---

## 4.2.1 ps

프로세스 목록을 확인합니다.

```bash
ps
ps -ef
ps aux
```

| 옵션    | 설명                 |
| ----- | ------------------ |
| `-e`  | 모든 프로세스            |
| `-f`  | 자세한 형식             |
| `aux` | BSD 스타일 전체 프로세스 출력 |

---

## 4.2.2 top

실시간으로 CPU, 메모리, 프로세스 상태를 확인합니다.

```bash
top
```

주요 정보:

| 항목      | 설명      |
| ------- | ------- |
| PID     | 프로세스 ID |
| USER    | 실행 사용자  |
| CPU     | CPU 사용률 |
| MEM     | 메모리 사용률 |
| COMMAND | 실행 명령   |

---

## 4.2.3 kill

프로세스를 종료합니다.

```bash
kill PID
kill -9 PID
```

| 신호   | 의미                |
| ---- | ----------------- |
| `15` | 정상 종료 요청, 기본값     |
| `9`  | 강제 종료             |
| `1`  | 재시작 또는 설정 재로드에 사용 |

---

## 4.2.4 systemctl

`systemctl`은 systemd 기반 Linux에서 서비스를 관리하는 명령어입니다.

```bash
systemctl start sshd
systemctl stop sshd
systemctl restart sshd
systemctl status sshd
systemctl enable sshd
systemctl disable sshd
```

| 명령        | 설명            |
| --------- | ------------- |
| `start`   | 서비스 시작        |
| `stop`    | 서비스 중지        |
| `restart` | 서비스 재시작       |
| `status`  | 서비스 상태 확인     |
| `enable`  | 부팅 시 자동 시작    |
| `disable` | 부팅 시 자동 시작 해제 |

---

# 4.3 텍스트 명령어

텍스트 명령어는 로그 분석, 설정 파일 확인, 파일 내용 검색에 자주 사용됩니다.

| 명령어    | 설명               |
| ------ | ---------------- |
| `cat`  | 파일 전체 출력         |
| `more` | 파일을 페이지 단위로 출력   |
| `less` | 파일을 위아래로 이동하며 확인 |
| `head` | 파일 앞부분 출력        |
| `tail` | 파일 뒷부분 출력        |
| `grep` | 특정 문자열 검색        |
| `sort` | 정렬               |
| `uniq` | 중복 제거            |
| `wc`   | 줄, 단어, 문자 수 계산   |
| `cut`  | 특정 필드 추출         |
| `awk`  | 패턴 처리 및 필드 기반 처리 |
| `sed`  | 문자열 치환/편집        |
| `diff` | 파일 차이 비교         |

---

## 4.3.1 cat

```bash
cat file.txt
```

파일 전체 내용을 출력합니다.

---

## 4.3.2 head

```bash
head file.txt
head -n 20 file.txt
```

파일 앞부분을 출력합니다.

---

## 4.3.3 tail

```bash
tail file.txt
tail -n 50 file.txt
tail -f /var/log/messages
```

| 옵션   | 설명                 |
| ---- | ------------------ |
| `-n` | 출력 줄 수 지정          |
| `-f` | 파일 변경 내용을 실시간으로 출력 |

서버 로그 확인할 때 자주 사용합니다.

---

## 4.3.4 grep

특정 문자열을 검색합니다.

```bash
grep "error" app.log
grep -i "error" app.log
grep -r "listen" /etc
```

| 옵션   | 설명              |
| ---- | --------------- |
| `-i` | 대소문자 무시         |
| `-r` | 하위 디렉터리 재귀 검색   |
| `-n` | 줄 번호 출력         |
| `-v` | 해당 문자열을 제외하고 출력 |

---

## 4.3.5 파이프와 리다이렉션

### 파이프 `|`

앞 명령어의 결과를 뒤 명령어의 입력으로 전달합니다.

```bash
ps -ef | grep ssh
```

의미:

```text
전체 프로세스 목록을 출력하고
그중 ssh가 포함된 줄만 검색
```

### 리다이렉션 `>`, `>>`

명령 결과를 파일로 저장합니다.

```bash
ls > list.txt
ls >> list.txt
```

| 기호   | 의미          |
| ---- | ----------- |
| `>`  | 파일 덮어쓰기     |
| `>>` | 파일 뒤에 추가    |
| `<`  | 파일을 입력으로 사용 |
| `2>` | 에러 출력 저장    |

---

# 4.4 네트워크 명령어

| 명령어          | 설명                |
| ------------ | ----------------- |
| `ip addr`    | IP 주소 확인          |
| `ip route`   | 라우팅 테이블 확인        |
| `ifconfig`   | 구버전 네트워크 인터페이스 확인 |
| `ping`       | 네트워크 연결 확인        |
| `traceroute` | 목적지까지 경로 확인       |
| `netstat`    | 네트워크 연결 상태 확인     |
| `ss`         | netstat 대체 명령     |
| `nslookup`   | DNS 질의            |
| `dig`        | DNS 상세 질의         |
| `hostname`   | 호스트명 확인           |
| `curl`       | URL 요청            |
| `wget`       | 파일 다운로드           |
| `ssh`        | 원격 접속             |
| `scp`        | SSH 기반 파일 복사      |
| `ftp`        | FTP 접속            |

---

## 4.4.1 IP 주소 확인

```bash
ip addr
```

또는 줄여서:

```bash
ip a
```

구버전에서는 다음 명령도 사용했습니다.

```bash
ifconfig
```

시험에서는 `ifconfig`가 자주 보일 수 있지만, 실제 최신 Linux에서는 `ip` 명령을 더 많이 사용합니다.

---

## 4.4.2 라우팅 테이블 확인

```bash
ip route
```

기본 게이트웨이 확인 예시:

```text
default via 192.168.0.1 dev eth0
```

의미:

```text
기본 경로는 192.168.0.1 게이트웨이를 통해 나간다.
```

---

## 4.4.3 ping

```bash
ping 8.8.8.8
ping google.com
```

네트워크 연결 상태를 확인합니다.

| 결과    | 의미                         |
| ----- | -------------------------- |
| 응답 있음 | 네트워크 연결 가능                 |
| 응답 없음 | 네트워크 장애, 방화벽 차단, 라우팅 문제 가능 |

---

## 4.4.4 netstat / ss

포트와 연결 상태를 확인합니다.

```bash
netstat -an
netstat -tulnp
ss -tulnp
```

| 옵션   | 설명        |
| ---- | --------- |
| `-t` | TCP       |
| `-u` | UDP       |
| `-l` | LISTEN 상태 |
| `-n` | 숫자 형태 출력  |
| `-p` | 프로세스 표시   |

예시:

```bash
ss -tulnp
```

웹 서버가 열려 있는지 확인할 때 사용합니다.

```text
LISTEN 0 128 0.0.0.0:80
```

---

## 4.4.5 DNS 확인

```bash
nslookup google.com
dig google.com
```

| 명령어        | 설명         |
| ---------- | ---------- |
| `nslookup` | 간단한 DNS 질의 |
| `dig`      | 상세한 DNS 질의 |

---

## 4.4.6 SSH

```bash
ssh user@192.168.0.10
```

SSH는 Linux 서버 원격 관리에 가장 많이 사용되는 프로토콜입니다.

| 항목    | 값                    |
| ----- | -------------------- |
| 프로토콜  | SSH                  |
| 기본 포트 | 22                   |
| 용도    | 원격 로그인, 명령 실행, 파일 복사 |

---

# 5. vi 편집기

## 5.1 vi란?

**vi**는 Linux/Unix 계열에서 사용하는 기본 텍스트 편집기입니다.

서버 환경에서는 GUI 편집기를 사용할 수 없는 경우가 많기 때문에, 설정 파일을 수정할 때 vi를 자주 사용합니다.

예시:

```bash
vi /etc/hosts
vi /etc/fstab
vi /etc/samba/smb.conf
```

---

## 5.2 vi의 모드

vi는 일반적인 메모장과 다르게 **모드**가 있습니다.

| 모드       | 설명                      |
| -------- | ----------------------- |
| 명령 모드    | 커서 이동, 삭제, 복사, 저장 명령 수행 |
| 입력 모드    | 실제 텍스트 입력               |
| 마지막 행 모드 | 저장, 종료, 검색, 치환 등 수행     |

처음 vi를 실행하면 **명령 모드**로 시작합니다.

---

## 5.3 vi 모드 전환

```text
명령 모드
 ├─ i, a, o 입력 → 입력 모드
 ├─ : 입력 → 마지막 행 모드
 └─ ESC → 명령 모드 복귀
```

| 키     | 동작                  |
| ----- | ------------------- |
| `i`   | 현재 커서 앞에서 입력        |
| `a`   | 현재 커서 뒤에서 입력        |
| `o`   | 현재 줄 아래 새 줄 생성 후 입력 |
| `ESC` | 명령 모드로 복귀           |
| `:`   | 마지막 행 모드 진입         |

---

## 5.4 저장과 종료

| 명령     | 설명            |
| ------ | ------------- |
| `:w`   | 저장            |
| `:q`   | 종료            |
| `:wq`  | 저장 후 종료       |
| `:q!`  | 저장하지 않고 강제 종료 |
| `:wq!` | 강제 저장 후 종료    |
| `ZZ`   | 저장 후 종료       |

가장 자주 쓰는 명령은 다음 두 개입니다.

```text
:wq
:q!
```

---

## 5.5 커서 이동

| 키     | 설명       |
| ----- | -------- |
| `h`   | 왼쪽       |
| `j`   | 아래       |
| `k`   | 위        |
| `l`   | 오른쪽      |
| `0`   | 줄의 처음    |
| `$`   | 줄의 끝     |
| `gg`  | 문서 처음    |
| `G`   | 문서 끝     |
| `숫자G` | 해당 줄로 이동 |

예시:

```text
10G
```

10번째 줄로 이동합니다.

---

## 5.6 삭제, 복사, 붙여넣기

| 명령     | 설명      |
| ------ | ------- |
| `x`    | 한 글자 삭제 |
| `dd`   | 한 줄 삭제  |
| `숫자dd` | 여러 줄 삭제 |
| `yy`   | 한 줄 복사  |
| `숫자yy` | 여러 줄 복사 |
| `p`    | 붙여넣기    |
| `u`    | 실행 취소   |

예시:

```text
3dd
```

현재 줄부터 3줄 삭제합니다.

```text
5yy
```

현재 줄부터 5줄 복사합니다.

---

## 5.7 검색과 치환

### 검색

```text
/검색어
```

다음 검색 결과로 이동:

```text
n
```

이전 검색 결과로 이동:

```text
N
```

### 치환

```text
:%s/기존문자/새문자/g
```

예시:

```text
:%s/old/new/g
```

문서 전체에서 `old`를 `new`로 바꿉니다.

---

## 5.8 vi 시험 포인트

| 키워드         | 의미            |
| ----------- | ------------- |
| `i`         | 입력 모드         |
| `ESC`       | 명령 모드         |
| `:w`        | 저장            |
| `:q`        | 종료            |
| `:wq`       | 저장 후 종료       |
| `:q!`       | 저장하지 않고 강제 종료 |
| `dd`        | 한 줄 삭제        |
| `yy`        | 한 줄 복사        |
| `p`         | 붙여넣기          |
| `/문자열`      | 검색            |
| `:%s/a/b/g` | 전체 치환         |

---

# 6. LINUX 활용

Linux 활용 부분은 서버 운영과 직접 연결됩니다.

핵심은 다음 네 가지입니다.

```text
Linux 활용
 ├─ 프로세스
 ├─ 데몬
 ├─ RPM
 └─ SAMBA
```

---

# 6.1 프로세스

## 6.1.1 프로세스란?

**프로세스 Process**는 실행 중인 프로그램입니다.

예를 들어 다음은 모두 프로세스입니다.

```text
sshd
httpd
mysqld
bash
vi
smbd
nmbd
```

프로그램은 디스크에 저장된 실행 파일이고, 프로세스는 메모리에서 실행 중인 상태입니다.

| 구분   | 설명         |
| ---- | ---------- |
| 프로그램 | 실행 가능한 파일  |
| 프로세스 | 실행 중인 프로그램 |
| PID  | 프로세스 ID    |
| PPID | 부모 프로세스 ID |

---

## 6.1.2 프로세스 확인

```bash
ps -ef
```

예시:

```text
UID   PID  PPID  CMD
root    1     0  /sbin/init
root  700     1  /usr/sbin/sshd
user 1200   700  -bash
```

| 항목   | 설명         |
| ---- | ---------- |
| UID  | 실행 사용자     |
| PID  | 프로세스 ID    |
| PPID | 부모 프로세스 ID |
| CMD  | 실행 명령      |

---

## 6.1.3 백그라운드 실행

명령 뒤에 `&`를 붙이면 백그라운드에서 실행됩니다.

```bash
long_script.sh &
```

작업 확인:

```bash
jobs
```

백그라운드 작업을 포그라운드로 가져오기:

```bash
fg
```

---

## 6.1.4 프로세스 종료

```bash
kill PID
kill -9 PID
```

| 명령            | 설명       |
| ------------- | -------- |
| `kill PID`    | 정상 종료 요청 |
| `kill -9 PID` | 강제 종료    |

프로세스 이름으로 종료:

```bash
pkill process_name
killall process_name
```

---

# 6.2 데몬

## 6.2.1 데몬이란?

**데몬 Daemon**은 백그라운드에서 계속 실행되며 특정 서비스를 제공하는 프로세스입니다.

보통 Linux의 서버 기능은 데몬으로 동작합니다.

| 데몬         | 역할             |
| ---------- | -------------- |
| `sshd`     | SSH 원격 접속 서비스  |
| `httpd`    | Apache 웹 서버    |
| `mysqld`   | MySQL DB 서버    |
| `named`    | DNS 서버         |
| `dhcpd`    | DHCP 서버        |
| `smbd`     | SAMBA 파일 공유    |
| `nmbd`     | NetBIOS 이름 서비스 |
| `crond`    | 예약 작업 실행       |
| `rsyslogd` | 시스템 로그 기록      |

데몬 이름은 관례적으로 뒤에 `d`가 붙는 경우가 많습니다.

```text
ssh + d = sshd
cron + d = crond
```

---

## 6.2.2 데몬 관리

최신 Linux에서는 `systemctl` 명령으로 데몬을 관리합니다.

```bash
systemctl status sshd
systemctl start sshd
systemctl stop sshd
systemctl restart sshd
systemctl enable sshd
systemctl disable sshd
```

| 명령        | 의미            |
| --------- | ------------- |
| `status`  | 현재 상태 확인      |
| `start`   | 즉시 시작         |
| `stop`    | 즉시 중지         |
| `restart` | 재시작           |
| `enable`  | 부팅 시 자동 시작    |
| `disable` | 부팅 시 자동 시작 해제 |

---

## 6.2.3 데몬 상태 확인 예시

```bash
systemctl status sshd
```

예시:

```text
Active: active (running)
```

의미:

```text
sshd 서비스가 정상 실행 중이다.
```

만약 다음처럼 나오면 서비스가 중지된 상태입니다.

```text
Active: inactive (dead)
```

---

# 6.3 RPM

## 6.3.1 RPM이란?

**RPM, Red Hat Package Manager**는 Red Hat 계열 Linux에서 사용하는 패키지 관리 형식이자 명령어입니다.

Red Hat 계열에는 다음 배포판이 포함됩니다.

```text
RHEL
CentOS
Rocky Linux
AlmaLinux
Fedora
```

RPM 패키지 파일은 보통 `.rpm` 확장자를 가집니다.

```text
httpd-2.4.57-1.el9.x86_64.rpm
```

---

## 6.3.2 패키지란?

패키지는 프로그램 설치에 필요한 파일을 묶어놓은 것입니다.

패키지에는 보통 다음 정보가 들어 있습니다.

| 구성         | 설명                |
| ---------- | ----------------- |
| 실행 파일      | 실제 프로그램           |
| 설정 파일      | `/etc` 아래 설정      |
| 라이브러리      | 프로그램 실행에 필요한 파일   |
| 문서         | 매뉴얼               |
| 의존성 정보     | 필요한 다른 패키지 정보     |
| 설치/삭제 스크립트 | 설치 과정에서 실행되는 스크립트 |

---

## 6.3.3 RPM 명령어

| 명령어              | 설명                    |
| ---------------- | --------------------- |
| `rpm -i 패키지.rpm` | 설치                    |
| `rpm -U 패키지.rpm` | 업그레이드                 |
| `rpm -e 패키지명`    | 삭제                    |
| `rpm -qa`        | 설치된 전체 패키지 목록         |
| `rpm -qi 패키지명`   | 패키지 정보 확인             |
| `rpm -ql 패키지명`   | 패키지가 설치한 파일 목록        |
| `rpm -qf 파일경로`   | 특정 파일이 어느 패키지 소속인지 확인 |
| `rpm -V 패키지명`    | 패키지 검증                |

---

## 6.3.4 RPM 예시

### 패키지 설치

```bash
rpm -i package.rpm
```

### 패키지 업그레이드

```bash
rpm -U package.rpm
```

### 패키지 삭제

```bash
rpm -e httpd
```

### 설치된 패키지 전체 확인

```bash
rpm -qa
```

### 특정 패키지 설치 여부 확인

```bash
rpm -qa | grep samba
```

---

## 6.3.5 RPM의 단점

RPM은 개별 패키지 파일을 직접 설치합니다.
이때 가장 큰 단점은 **의존성 dependency 문제**입니다.

예를 들어 `A.rpm`을 설치하려고 했는데, `B.rpm`과 `C.rpm`이 필요할 수 있습니다.

```text
A 패키지 설치
→ B 패키지가 필요함
→ C 패키지가 필요함
→ 직접 찾아서 설치해야 함
```

이 문제를 해결하기 위해 `yum`, `dnf` 같은 상위 패키지 관리 도구를 사용합니다.

---

## 6.3.6 RPM, YUM, DNF 차이

| 구분  | 설명                             |
| --- | ------------------------------ |
| RPM | 개별 `.rpm` 패키지 설치/관리            |
| YUM | 저장소를 이용해 패키지와 의존성을 자동 관리       |
| DNF | YUM의 개선 버전, 최신 Red Hat 계열에서 사용 |

예시:

```bash
dnf install samba
yum install samba
```

이 명령은 필요한 의존성까지 함께 설치합니다.

---

# 6.4 SAMBA

## 6.4.1 SAMBA란?

**SAMBA**는 Linux/Unix 시스템에서 Windows와 파일 및 프린터를 공유할 수 있게 해주는 서비스입니다.

Windows는 파일 공유에 주로 **SMB/CIFS** 프로토콜을 사용합니다.
Linux는 기본적으로 Windows 파일 공유 방식과 다르기 때문에, Windows 클라이언트와 파일을 공유하려면 SAMBA를 사용합니다.

```text
Windows PC ↔ SMB/CIFS ↔ SAMBA 서버 Linux
```

즉, SAMBA는 Linux 서버를 Windows 파일 서버처럼 사용할 수 있게 해줍니다.

---

## 6.4.2 SAMBA가 필요한 이유

회사에서 Linux 서버와 Windows PC가 함께 사용된다고 가정해보겠습니다.

```text
직원 PC: Windows
파일 서버: Linux
```

Windows 사용자는 파일 탐색기에서 다음처럼 접근하고 싶어 합니다.

```text
\\192.168.0.10\share
\\fileserver\public
```

이때 Linux 서버가 Windows 파일 공유를 제공하려면 SAMBA가 필요합니다.

---

## 6.4.3 SAMBA의 주요 용도

| 용도              | 설명                               |
| --------------- | -------------------------------- |
| 파일 공유           | Linux 디렉터리를 Windows에서 공유 폴더처럼 접근 |
| 프린터 공유          | Linux에 연결된 프린터를 Windows에서 사용     |
| 인증 연동           | 사용자 계정 기반 접근 제어                  |
| Windows 네트워크 참여 | Workgroup 또는 Domain 환경과 연동       |
| NAS 구축          | Linux 서버를 파일 서버/NAS처럼 활용         |

---

## 6.4.4 SAMBA 관련 프로토콜

| 프로토콜/기술 | 설명                                       |
| ------- | ---------------------------------------- |
| SMB     | Server Message Block, Windows 파일 공유 프로토콜 |
| CIFS    | SMB의 확장/구현 명칭으로 자주 사용                    |
| NetBIOS | 예전 Windows 네트워크 이름 서비스                   |
| WINS    | NetBIOS 이름 해석 서비스                        |
| TCP 445 | SMB 직접 통신 포트                             |
| TCP 139 | NetBIOS Session Service                  |
| UDP 137 | NetBIOS Name Service                     |
| UDP 138 | NetBIOS Datagram Service                 |

시험에서는 특히 다음 포트를 기억하면 좋습니다.

| 서비스              |      포트 |
| ---------------- | ------: |
| SMB              | TCP 445 |
| NetBIOS Session  | TCP 139 |
| NetBIOS Name     | UDP 137 |
| NetBIOS Datagram | UDP 138 |

---

## 6.4.5 SAMBA 주요 데몬

SAMBA는 여러 데몬으로 구성됩니다.

| 데몬         | 역할                     |
| ---------- | ---------------------- |
| `smbd`     | 파일 공유, 프린터 공유, 인증 처리   |
| `nmbd`     | NetBIOS 이름 해석, 브라우징 기능 |
| `winbindd` | Windows 도메인 계정 연동      |

핵심은 다음입니다.

```text
smbd = 실제 파일/프린터 공유 담당
nmbd = Windows 이름 서비스 담당
winbindd = Windows 도메인 계정 연동 담당
```

---

## 6.4.6 SAMBA 설정 파일

SAMBA의 대표 설정 파일은 다음입니다.

```bash
/etc/samba/smb.conf
```

이 파일에서 공유 이름, 공유 경로, 접근 권한, 인증 방식을 설정합니다.

---

## 6.4.7 smb.conf 기본 구조

`smb.conf`는 크게 두 부분으로 나뉩니다.

```ini
[global]
    전역 설정

[share]
    공유 폴더 설정
```

예시:

```ini
[global]
    workgroup = WORKGROUP
    server string = Samba Server
    security = user

[public]
    path = /srv/samba/public
    browseable = yes
    writable = yes
    guest ok = yes
```

---

## 6.4.8 [global] 설정

`[global]`은 SAMBA 서버 전체 설정입니다.

| 설정              | 설명               |
| --------------- | ---------------- |
| `workgroup`     | Windows 작업 그룹 이름 |
| `server string` | 서버 설명            |
| `security`      | 인증 방식            |
| `map to guest`  | 게스트 매핑 방식        |
| `interfaces`    | 사용할 네트워크 인터페이스   |
| `hosts allow`   | 접근 허용 IP         |
| `log file`      | 로그 파일 위치         |

예시:

```ini
[global]
    workgroup = WORKGROUP
    server string = Linux File Server
    security = user
```

---

## 6.4.9 공유 폴더 설정

예시:

```ini
[share]
    path = /srv/samba/share
    browseable = yes
    writable = yes
    valid users = user1 user2
```

| 설정               | 설명                  |
| ---------------- | ------------------- |
| `[share]`        | Windows에서 보이는 공유 이름 |
| `path`           | 실제 Linux 디렉터리 경로    |
| `browseable`     | 네트워크 목록에 표시 여부      |
| `writable`       | 쓰기 가능 여부            |
| `read only`      | 읽기 전용 여부            |
| `guest ok`       | 게스트 접근 허용 여부        |
| `valid users`    | 접근 허용 사용자           |
| `write list`     | 쓰기 허용 사용자           |
| `create mask`    | 생성 파일 기본 권한         |
| `directory mask` | 생성 디렉터리 기본 권한       |

---

## 6.4.10 SAMBA 공유 예시 1: 익명 공유

Windows에서 누구나 접근 가능한 공유 폴더를 만들 때 사용하는 형태입니다.

```ini
[public]
    path = /srv/samba/public
    browseable = yes
    writable = yes
    guest ok = yes
    read only = no
```

의미:

| 설정                 | 의미                         |
| ------------------ | -------------------------- |
| `[public]`         | 공유 이름은 public              |
| `path`             | 실제 경로는 `/srv/samba/public` |
| `browseable = yes` | 네트워크에서 보이게 함               |
| `writable = yes`   | 쓰기 가능                      |
| `guest ok = yes`   | 로그인 없이 접근 가능               |
| `read only = no`   | 읽기 전용 아님                   |

하지만 실제 서버에서는 보안상 익명 쓰기 공유는 주의해야 합니다.

---

## 6.4.11 SAMBA 공유 예시 2: 사용자 인증 공유

특정 사용자만 접근할 수 있는 공유입니다.

```ini
[private]
    path = /srv/samba/private
    browseable = yes
    writable = yes
    valid users = yuchan
```

이 경우 `yuchan` 사용자만 접근 가능합니다.

---

## 6.4.12 SAMBA 사용자 계정

SAMBA는 Linux 사용자와 별도로 SAMBA 암호를 관리합니다.

먼저 Linux 사용자가 있어야 합니다.

```bash
useradd yuchan
passwd yuchan
```

그다음 SAMBA 사용자로 등록합니다.

```bash
smbpasswd -a yuchan
```

SAMBA 계정 활성화:

```bash
smbpasswd -e yuchan
```

SAMBA 계정 삭제:

```bash
smbpasswd -x yuchan
```

중요한 점:

```text
Linux 사용자 계정이 있어야 SAMBA 사용자로 등록할 수 있다.
```

---

## 6.4.13 SAMBA 디렉터리 권한

SAMBA 설정에서 쓰기 권한을 줘도, Linux 파일 시스템 권한이 막혀 있으면 쓰기가 안 됩니다.

즉, SAMBA 접근 권한과 Linux 권한을 둘 다 확인해야 합니다.

```text
SAMBA 설정 권한
+ Linux 파일 시스템 권한
= 실제 접근 가능 여부
```

예시:

```bash
mkdir -p /srv/samba/share
chown -R yuchan:yuchan /srv/samba/share
chmod 755 /srv/samba/share
```

또는 그룹 공유라면:

```bash
groupadd smbusers
usermod -aG smbusers yuchan
chown -R root:smbusers /srv/samba/share
chmod -R 2770 /srv/samba/share
```

여기서 `2770`의 `2`는 setgid입니다.
디렉터리에 setgid가 설정되면, 그 안에 새로 생성되는 파일/디렉터리가 부모 디렉터리의 그룹을 상속받습니다.

---

## 6.4.14 SAMBA 서비스 시작

Red Hat 계열 기준:

```bash
systemctl start smb
systemctl start nmb
```

부팅 시 자동 시작:

```bash
systemctl enable smb
systemctl enable nmb
```

상태 확인:

```bash
systemctl status smb
systemctl status nmb
```

Debian/Ubuntu 계열에서는 서비스명이 다를 수 있습니다.

```bash
systemctl status smbd
systemctl status nmbd
```

시험에서는 보통 **smbd, nmbd**를 데몬 이름으로 기억하면 됩니다.

---

## 6.4.15 방화벽 설정

SAMBA가 정상 동작하려면 방화벽에서 SMB 관련 포트가 열려 있어야 합니다.

Red Hat 계열 firewalld 예시:

```bash
firewall-cmd --permanent --add-service=samba
firewall-cmd --reload
```

또는 포트 기준:

```bash
firewall-cmd --permanent --add-port=445/tcp
firewall-cmd --permanent --add-port=139/tcp
firewall-cmd --permanent --add-port=137/udp
firewall-cmd --permanent --add-port=138/udp
firewall-cmd --reload
```

---

## 6.4.16 Windows에서 SAMBA 접속

Windows 파일 탐색기 주소창에 다음처럼 입력합니다.

```text
\\192.168.0.10\public
```

또는 서버 이름으로 접근합니다.

```text
\\fileserver\public
```

사용자 인증 공유라면 계정과 비밀번호를 입력해야 합니다.

```text
사용자명: yuchan
비밀번호: smbpasswd로 설정한 비밀번호
```

---

## 6.4.17 Linux에서 SAMBA 접속 확인

Linux 클라이언트에서 공유 목록을 확인할 수 있습니다.

```bash
smbclient -L //192.168.0.10 -U yuchan
```

공유 폴더 접속:

```bash
smbclient //192.168.0.10/private -U yuchan
```

---

## 6.4.18 SAMBA 설정 검사

`smb.conf` 문법 오류를 검사하려면 다음 명령을 사용합니다.

```bash
testparm
```

설정 파일을 수정한 뒤에는 다음 순서로 확인하면 좋습니다.

```text
smb.conf 수정
→ testparm으로 문법 검사
→ smb/nmb 재시작
→ 방화벽 확인
→ Windows에서 접속 테스트
```

---

## 6.4.19 SAMBA 문제 해결 체크리스트

SAMBA 접속이 안 될 때는 다음 순서로 확인합니다.

| 확인 항목         | 명령/내용                          |                     |
| ------------- | ------------------------------ | ------------------- |
| 서비스 실행 여부     | `systemctl status smb`         |                     |
| 설정 문법         | `testparm`                     |                     |
| 포트 열림 여부      | `ss -tulnp                     | grep smb` 또는 445 확인 |
| 방화벽           | `firewall-cmd --list-all`      |                     |
| 공유 경로 존재 여부   | `ls -ld /srv/samba/share`      |                     |
| Linux 권한      | `chmod`, `chown` 확인            |                     |
| SAMBA 사용자 등록  | `pdbedit -L` 또는 `smbpasswd -a` |                     |
| Windows 접근 경로 | `\\서버IP\공유이름`                  |                     |
| SELinux       | Red Hat 계열에서 차단 가능             |                     |

---

## 6.4.20 SAMBA에서 자주 헷갈리는 부분

### SAMBA 설정 권한과 Linux 권한은 다르다

`smb.conf`에 다음처럼 되어 있어도:

```ini
writable = yes
```

Linux 디렉터리 권한이 쓰기를 막고 있으면 쓰기 불가입니다.

```bash
dr-xr-xr-x /srv/samba/share
```

따라서 둘 다 맞아야 합니다.

---

### 공유 이름과 실제 경로는 다르다

```ini
[public]
    path = /srv/samba/public
```

Windows에서 보이는 이름은:

```text
public
```

Linux 실제 경로는:

```text
/srv/samba/public
```

접속 경로는:

```text
\\192.168.0.10\public
```

---

### Linux 계정과 SAMBA 계정은 연결되지만 암호는 따로 관리된다

Linux 사용자 생성:

```bash
useradd yuchan
```

Linux 비밀번호 설정:

```bash
passwd yuchan
```

SAMBA 비밀번호 설정:

```bash
smbpasswd -a yuchan
```

즉, Linux 계정이 있어야 하지만 SAMBA 인증 정보는 별도로 등록해야 합니다.

---

## 6.4.21 SAMBA 핵심 암기표

| 키워드                   | 의미                             |
| --------------------- | ------------------------------ |
| SAMBA                 | Linux와 Windows 간 파일/프린터 공유 서비스 |
| SMB/CIFS              | Windows 파일 공유 프로토콜             |
| `smbd`                | 파일/프린터 공유 담당                   |
| `nmbd`                | NetBIOS 이름 서비스 담당              |
| `winbindd`            | Windows 도메인 계정 연동              |
| `/etc/samba/smb.conf` | SAMBA 설정 파일                    |
| `smbpasswd`           | SAMBA 사용자 암호 관리                |
| `testparm`            | SAMBA 설정 문법 검사                 |
| TCP 445               | SMB 기본 포트                      |
| TCP 139               | NetBIOS Session                |
| UDP 137               | NetBIOS Name                   |
| UDP 138               | NetBIOS Datagram               |
| `\\IP\공유명`            | Windows에서 SAMBA 공유 접근 방식       |

---

# 7. 전체 핵심 암기표

## 7.1 Linux 기본 개념

| 키워드    | 의미                |
| ------ | ----------------- |
| Linux  | Unix 계열 오픈소스 운영체제 |
| Kernel | 운영체제 핵심           |
| Shell  | 명령어 해석기           |
| bash   | 대표적인 Linux Shell  |
| Root   | 최고 관리자            |
| `/`    | 루트 디렉터리           |
| CLI    | 명령어 기반 인터페이스      |

---

## 7.2 주요 디렉터리

| 디렉터리    | 설명              |
| ------- | --------------- |
| `/etc`  | 설정 파일           |
| `/home` | 일반 사용자 홈        |
| `/root` | root 홈          |
| `/var`  | 로그, 메일, 캐시      |
| `/bin`  | 기본 명령어          |
| `/sbin` | 시스템 관리 명령어      |
| `/dev`  | 장치 파일           |
| `/proc` | 커널/프로세스 정보      |
| `/tmp`  | 임시 파일           |
| `/usr`  | 사용자 프로그램, 라이브러리 |

---

## 7.3 권한

| 권한  |        숫자 | 의미                |
| --- | --------: | ----------------- |
| r   |         4 | 읽기                |
| w   |         2 | 쓰기                |
| x   |         1 | 실행                |
| 755 | rwxr-xr-x | 실행 파일/디렉터리에 자주 사용 |
| 644 | rw-r--r-- | 일반 파일에 자주 사용      |
| 700 | rwx------ | 소유자만 접근           |
| 777 | rwxrwxrwx | 모두 허용, 보안상 위험     |

---

## 7.4 명령어

| 분류     | 주요 명령어                                                         |
| ------ | -------------------------------------------------------------- |
| 파일 관리  | `ls`, `cd`, `pwd`, `cp`, `mv`, `rm`, `mkdir`, `touch`          |
| 권한 관리  | `chmod`, `chown`, `chgrp`                                      |
| 시스템 관리 | `ps`, `top`, `kill`, `df`, `du`, `free`, `systemctl`           |
| 텍스트 처리 | `cat`, `more`, `less`, `head`, `tail`, `grep`, `sort`, `wc`    |
| 네트워크   | `ip`, `ping`, `netstat`, `ss`, `nslookup`, `dig`, `ssh`, `scp` |
| 패키지    | `rpm`, `yum`, `dnf`, `apt`                                     |

---

# 8. 네트워크관리사 2급 관점 최종 정리

Linux는 네트워크관리사 2급에서 다음 관점으로 정리하면 됩니다.

## 8.1 운영체제 관점

Linux는 오픈소스 Unix 계열 운영체제이며, 서버와 네트워크 장비에서 많이 사용됩니다.

```text
Linux = Kernel + Shell + File System + Commands + Network Services
```

---

## 8.2 서버 관리 관점

Linux 서버는 주로 CLI로 관리합니다.

```text
SSH 접속
→ 설정 파일 수정
→ 서비스 재시작
→ 로그 확인
→ 네트워크 상태 확인
```

예시:

```bash
ssh user@server
vi /etc/samba/smb.conf
systemctl restart smb
tail -f /var/log/messages
ss -tulnp
```

---

## 8.3 권한 관리 관점

Linux 보안의 기본은 사용자, 그룹, 권한입니다.

```text
사용자 User
그룹 Group
기타 Others

r = 4
w = 2
x = 1
```

예시:

```bash
chmod 755 script.sh
chown user:group file.txt
```

---

## 8.4 네트워크 서비스 관점

Linux는 다양한 네트워크 서비스를 제공합니다.

| 서비스          | 역할                 |
| ------------ | ------------------ |
| SSH          | 원격 접속              |
| Apache/Nginx | 웹 서버               |
| DNS          | 이름 해석              |
| DHCP         | IP 자동 할당           |
| FTP          | 파일 전송              |
| SAMBA        | Windows와 파일 공유     |
| NFS          | Linux/Unix 간 파일 공유 |

---

# 9. 최종 암기 문장

> Linux는 Unix 계열의 오픈소스 운영체제로, 서버와 네트워크 환경에서 안정성, 보안성, 자동화, 네트워크 기능이 강한 운영체제다.

> Linux 구조는 사용자, Shell, Kernel, Hardware로 이해할 수 있으며, Shell은 명령어를 해석하고 Kernel은 시스템 자원을 관리한다.

> Linux 파일 시스템은 `/` 루트 디렉터리를 기준으로 계층 구조를 가지며, `/etc`는 설정 파일, `/home`은 일반 사용자 홈, `/var`는 로그와 변동 데이터를 저장한다.

> Linux 권한은 사용자, 그룹, 기타 사용자 기준으로 r, w, x 권한을 부여하며, 숫자로는 r=4, w=2, x=1로 계산한다.

> chmod는 권한 변경, chown은 소유자 변경, chgrp는 그룹 변경 명령어다.

> vi는 Linux 기본 편집기로, 명령 모드, 입력 모드, 마지막 행 모드를 구분해서 사용한다.

> 프로세스는 실행 중인 프로그램이고, 데몬은 백그라운드에서 지속적으로 동작하며 서비스를 제공하는 프로세스다.

> RPM은 Red Hat 계열 Linux에서 사용하는 패키지 관리 방식이며, YUM과 DNF는 의존성을 자동으로 처리하는 상위 패키지 관리 도구다.

> SAMBA는 Linux 서버를 Windows 파일 서버처럼 사용할 수 있게 해주는 서비스이며, SMB/CIFS 프로토콜을 사용한다.

> SAMBA의 핵심 데몬은 smbd와 nmbd이고, 설정 파일은 `/etc/samba/smb.conf`, 주요 포트는 TCP 445, TCP 139, UDP 137, UDP 138이다.
