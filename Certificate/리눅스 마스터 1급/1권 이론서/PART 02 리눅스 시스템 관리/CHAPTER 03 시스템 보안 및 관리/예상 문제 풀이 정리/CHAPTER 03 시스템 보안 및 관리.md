## 문제 1번 : 정답 1

```

다음은 GRUB 패스워드를 설정하는 과정의 일부이다. (　) 안에 들어갈 내용으로 알맞은 것은?

$ grub2-mkpasswd-pbkdf2
Enter password:
Reenter password:
PBKDF2 hash of your password is grub.pbkdf2.sha512...

# vi /etc/grub.d/40_custom
set superusers="root"
(　) root grub.pbkdf2.sha512...

1. password_pbkdf2
2. password
3. password_sha512
4. password_pbkdf2_sha512

```

## 문제 2번 : 정답 1

```

커널 메시지 중 에러 이상의 메시지만 ihduser 사용자의 터미널에 나타나도록 rsyslog.conf 파일에 설정하는 내용으로 알맞은 것은?

1. kern.err    ihduser
2. kern.crit   @ihduser
3. @ihduser    kern.err
4. ihduser     kern.crit

```

## 문제 3번 : 정답 2

```

다음 (　) 안에 사용할 수 있는 명령으로 알맞은 것은?

# ls -l | (　) > home.backup

1. tar 명령과 옵션
2. cpio 명령과 옵션
3. dump 명령과 옵션
4. rsync 명령과 옵션

```

## 문제 4번 : 정답 1

```

다음은 tar 명령으로 증분 백업한 과정이다. 복원 순서를 고르시오.

# tar list -cvfp home.tar /home
# tar list -cvfp home2.tar /home
# tar list -cvfp home3.tar /home
# tar list -cvfp home4.tar /home

1. home.tar → home2.tar → home3.tar → home4.tar
2. home4.tar → home3.tar → home2.tar → home.tar
3. home.tar → home3.tar → home2.tar → home4.tar
4. home.tar → home2.tar → home4.tar → home3.tar

```

## 문제 5번 : 정답 1

```

Project 그룹에 속하지 않은 ihduser에게 Project 그룹 소유의 document 파일 쓰기 권한을 설정하는 명령은?

1. setfacl -m u:ihduser:w document
2. setfacl -x u:ihduser:w document
3. setfacl -m g:ihduser:w document
4. setfacl -x g:ihduser:w document

```

## 문제 6번 : 정답 3

```

다음 설명에 해당하는 백업 도구로 가장 알맞은 것은?

로컬과 원격 호스트 또는 원격 셸 간 파일 복사가 가능하다. 원본과 대상 파일의 차이만 계산해 전송하는 델타 전송 알고리즘으로 속도를 높이며 반복 백업과 미러링에 사용된다.

1. dd
2. cpio
3. rsync
4. dump

```

## 문제 7번 : 정답 4

```

root 권한으로 실행되는 데몬을 최소화하면서 미리 지정한 권한으로 자원 접근을 강제하는 보안 기능은?

1. Nessus
2. GnuPG
3. John the Ripper
4. SELinux

```

## 문제 8번 : 정답 3

```

각 계정에 대한 마지막 로그인 정보를 확인하는 명령으로 알맞은 것은?

1. last
2. lastb
3. lastlog
4. dmesg

```

## 문제 9번 : 정답 4

```

다음 중 cat 명령으로 텍스트 모드 열람이 가능한 로그 파일은?

1. wtmp
2. lastlog
3. btmp
4. secure

```

## 문제 10번 : 정답 1

```

백업 대상이 되는 디렉터리의 조합으로 알맞은 것은?

- ㉠ /etc
- ㉡ /tmp
- ㉢ /usr
- ㉣ /var

1. ㉠, ㉢, ㉣
2. ㉠, ㉡, ㉢
3. ㉡, ㉢, ㉣
4. ㉠, ㉡, ㉢, ㉣

```

## 문제 11번 : 정답 2

```

다음 중 rsync 명령에 관한 설명으로 틀린 것은?

1. 네트워크로 연결된 원격지의 파일들을 동기화하는 유틸리티이다.
2. ssh나 rsh를 이용해 전송할 수 있고 root 권한이 필요하다.
3. rcp에 비해 처리 속도가 우수하고 다양한 기능을 제공한다.
4. 로컬 시스템 백업 시 별도의 서버 설정 없이 사용할 수 있다.

```

## 문제 12번 : 정답 4

```

다음 중 dd 명령에 대한 설명으로 틀린 것은?

1. 파티션이나 디스크 단위 백업에 사용한다.
2. 복사할 블록 단위를 지정할 수 있다.
3. CD 장치에서 ISO 이미지를 추출할 수 있다.
4. 파일 단위로 백업할 수 있다.

```

## 문제 13번 : 정답 4

```

SSH(Secure Shell)의 설명으로 틀린 것은?

1. 원격 복사(scp)를 지원한다.
2. 안전한 파일 전송(sftp)을 지원한다.
3. 기본 사용 포트는 22이다.
4. rsh와 같은 원격 셸을 지원하지 않는다.

```

## 문제 14번 : 정답 2

```

2004년 Rainer Gerhards가 시작했으며 TCP, SSL, TLS, RELP, 멀티 스레드, DB 저장 및 로그 필터링 등을 지원하는 로그 시스템은?

1. logrotate
2. rsyslog
3. lastlog
4. vsyslog

```

## 문제 15번 : 정답 4

```

다음 중 PAM의 control-flag 설명으로 틀린 것은?

1. requisite: 반드시 성공해야 하며 실패 시 즉시 실패 반환
2. required: 반드시 성공해야 하나 실패해도 같은 module-type 검사를 일부 수행한 뒤 실패 반환
3. sufficient: 성공하면 나머지 모듈을 더 이상 검사하지 않음
4. optional: 해당 모듈의 성공은 중요하지만 실패는 중요하지 않음

```

## 문제 16번 : 정답 3

```

가장 최근의 시스템 재부팅 정보 두 항목을 출력하는 명령은?

1. lastb -2 reboot
2. lastb -L 2 reboot
3. last -2 reboot
4. last -L 2 reboot

```

## 문제 17번 : 정답 2

```

다음 중 sudo에 관련된 설명이 아닌 것은?

1. 특정 사용자 또는 그룹에 root 권한을 위임하는 도구이다.
2. 환경설정 파일 편집 시 일반 vi를 사용하는 것이 안전하다.
3. 적용된 사용자는 sudo 명령어 형식으로 root 권한을 대행한다.
4. 환경설정 파일은 /etc/sudoers이다.

```

## 문제 18번 : 정답 4

```

부팅 시 커널 링 버퍼의 메시지가 최종적으로 보관되는 파일은?

1. /var/log/secure
2. /var/log/lastlog
3. /var/log/wtmp
4. /var/log/dmesg

```

## 문제 19번 : 정답 3

```

다음 각 명령어와 관계있는 파일 연결이 틀린 것은?

1. lastlog : /var/log/lastlog
2. dmesg : /var/log/dmesg
3. last : /var/log/tmp
4. lastb : /var/log/btmp

```
