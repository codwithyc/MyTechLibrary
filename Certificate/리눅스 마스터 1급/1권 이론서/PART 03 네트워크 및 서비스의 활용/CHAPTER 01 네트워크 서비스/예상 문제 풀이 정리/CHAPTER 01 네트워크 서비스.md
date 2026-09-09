## 문제 1번 : 정답 1

```

HTTP 요청을 헤더와 바디로 구분하고 요청 내용을 바디에 담아 서버에 전송하며, 주로 폼을 이용한 데이터 전송에 사용하는 메서드는?

1. POST
2. PUT
3. GET
4. OPTIONS

```

## 문제 2번 : 정답 3

```

다음 중 HTTP 상태 코드와 설명이 잘 연결된 것은?

1. 100 : 요청 성공
2. 200 : 리다이렉션
3. 404 : 존재하지 않는 리소스 요청(Not Found)
4. 500 : 리소스에 대한 권한 없음(Forbidden)

```

## 문제 3번 : 정답 2

```

아파치 웹 서버의 모듈 적재 방법 중 DSO(Dynamic Shared Object) 방식의 설명으로 잘못된 것은?

1. 웹 서버 실행 후 사용자 요청이 있을 때 필요한 모듈을 적재한다.
2. 새 모듈 추가 시 아파치 소스 코드를 다시 컴파일해 적용한다.
3. 요청 후 모듈 적재 시간이 필요하여 응답이 다소 느릴 수 있다.
4. 모듈 추가 및 삭제가 편리하다.

```

## 문제 4번 : 정답 4

```

아파치 httpd.conf의 세부 항목에 대한 설명으로 적절한 것은?

1. ServerRoot : HTML 등 웹 콘텐츠 저장 디렉터리 지정
2. ServerTokens : 암호화 통신에 사용할 키 저장
3. KeepAlive : On이면 연결당 요청과 응답을 한 번만 처리
4. UserDir : 일반 사용자의 웹 디렉터리 지정

```

## 문제 5번 : 정답 1

```

MySQL을 소스 코드 컴파일 방식으로 설치할 때 다음 (　) 안에 들어갈 명령은?

$ cd mysql-5.7.28/
$ (　) -DCMAKE_INSTALL_PREFIX=/usr/local/mysql \
  -DDEFAULT_CHARSET=utf8 \
  -DDEFAULT_COLLATION=utf8_general_ci \
  -DENABLED_LOCAL_INFILE=1 \
  -DMYSQL_DATADIR=/usr/local/mysql/data

1. cmake
2. make
3. config
4. build

```

## 문제 6번 : 정답 4

```

NIS 서버 및 클라이언트 관련 명령 설명 중 잘못된 것은?

1. ypwhich : 로그인 인증에 사용한 NIS 서버 조회
2. ypcat : NIS 서버 맵 파일 내용 확인
3. yppasswd : NIS 사용자 비밀번호 변경
4. yptest : NIS 서비스 품질을 확인하기 위해 네트워크 응답 속도 측정

```

## 문제 7번 : 정답 2

```

다음 설명에 해당하는 서비스는?

디렉터리 서비스를 조회·수정하는 TCP 기반 응용 프로토콜이다. 엔트리가 트리 구조로 구성되고 각 엔트리는 여러 이름=값 속성으로 구성된다.

1. NIS
2. LDAP
3. SELinux
4. MongoDB

```

## 문제 8번 : 정답 3

```

다음 중 NIS 설명으로 알맞은 것은?

1. nisdomainname으로 설정한 도메인명은 재부팅 후에도 유지된다.
2. ypxfrd는 NIS 서버와 클라이언트의 시간을 동기화한다.
3. /etc/sysconfig/network에 NIS 도메인명을 설정해 부팅 시 적용할 수 있다.
4. NIS 패키지를 설치하면 /etc/rc.d/init.d에 관련 데몬 스크립트가 설치된다.

```

## 문제 9번 : 정답 4

```

다음 중 삼바(SAMBA)에 대한 설명으로 틀린 것은?

1. 리눅스와 윈도우 간 디렉터리 및 파일 공유에 사용한다.
2. nmblookup으로 NetBIOS 이름에 대응하는 IP 주소를 조회할 수 있다.
3. smbstatus로 현재 접속 중인 클라이언트 목록을 확인할 수 있다.
4. smb.conf에서 접근 가능한 사용자를 지정하는 항목은 valid-users이다.

```

## 문제 10번 : 정답 3

```

NFS 서버의 접근 권한을 다음과 같이 설정했을 때 알맞은 설명은?

/data 192.168.5.0/24(rw,root_squash)

1. 192.168.5.24만 읽기·쓰기 가능하며 root 권한 인정
2. 192.168.5.24만 읽기·쓰기 가능하며 root 권한 불인정
3. 192.168.5.0/24의 모든 호스트가 읽기·쓰기 가능하며 root 권한 불인정
4. 192.168.5.0/24의 모든 호스트가 읽기·쓰기 가능하며 root 권한 인정

```

## 문제 11번 : 정답 4

```

FTP 서비스에 관한 설명으로 틀린 것은?

1. Active 모드는 방화벽 때문에 데이터 연결이 거부될 수 있다.
2. Passive 모드는 서버가 사용할 데이터 포트 번호를 클라이언트에 알려준다.
3. Active 모드에서 서버는 TCP 20번을 데이터 포트로 사용한다.
4. /etc/vsftpd/ftpusers는 FTP에 접근할 수 있는 사용자를 지정한다.

```

## 문제 12번 : 정답 1

```

이메일 서비스 관련 프로토콜 설명으로 틀린 것은?

1. SMTP는 인터넷 이메일 전송에 TCP 21번을 사용한다.
2. POP3는 TCP 110번을 사용한다.
3. IMAP은 TCP 143번을 사용한다.
4. POP3는 수신 후 서버 메일을 삭제할 수 있으나 IMAP은 서버에 남겨둘 수 있다.

```

## 문제 13번 : 정답 2

```

/etc/mail/access 정보를 읽어 관련 DB를 업데이트하는 명령은?

1. mailq -Ac
2. makemap hash /etc/mail/access < /etc/mail/access
3. sendmail -bi
4. m4 sendmail.mc > sendmail.cf

```

## 문제 14번 : 정답 2

```

Primary 서버의 zone 파일을 백업하고 장애 시 보조 DNS로 사용하는 서버는?

1. Primary Name Server
2. Secondary Name Server
3. Caching Name Server
4. Proxy Server

```

## 문제 15번 : 정답 4

```

DNS zone 파일에서 MX 타입을 선언하는 이유는?

1. 도메인 질의 시 IP 주소 조회
2. IP 주소 질의 시 도메인 주소 조회
3. 도메인의 별칭 지정
4. 도메인의 메일 교환 서버 정보 지정

```

## 문제 16번 : 정답 3

```

텍스트 기반 콘솔에서 가상머신 생성·시작·재시작·종료·강제 종료 등을 수행하는 도구는?

1. virt-top
2. virt-manager
3. virsh
4. libvirtd

```

## 문제 17번 : 정답 4

```

다음 중 리눅스 슈퍼데몬에 대한 설명으로 가장 알맞은 것은?

1. inetd는 커널 2.4 이후부터 사용하는 방식이다.
2. 서비스 요청이 빈번하고 빠른 응답이 필요한 경우에 적합하다.
3. telnet, rlogin, ftp 등이 일반적으로 inetd 방식으로 실행된다.
4. tcp_wrapper의 /etc/hosts.allow, /etc/hosts.deny로 접근 제어한다.

```

## 문제 18번 : 정답 4

```

다음 squid.conf 설정의 (　)에 들어갈 내용으로 가장 알맞은 것은?

(　) localnet src 192.168.12.0/255.255.255.0
(　) localnet

1. access, http allow
2. range, http_allow
3. acl, http access allow
4. acl, http_access allow

```

## 문제 19번 : 정답 1

```

dhcpd.conf에서 특정 MAC 주소의 시스템에 고정 IP를 할당할 때 사용하는 항목은?

1. fixed-address
2. default-address
3. range
4. dynamic-address

```

## 문제 20번 : 정답 3

```

다음 중 VNC 관련 명령어 설명으로 잘못된 것은?

1. systemctl start vncserver@:1.service : 디스플레이 1번으로 VNC 서비스 실행
2. vncpasswd : VNC 서버 접속용 패스워드 설정
3. Xvnc : vncserver에 의해 실행되는 VNC Client(Viewer)
4. vncconfig : VNC 서비스 설정 관리

```
