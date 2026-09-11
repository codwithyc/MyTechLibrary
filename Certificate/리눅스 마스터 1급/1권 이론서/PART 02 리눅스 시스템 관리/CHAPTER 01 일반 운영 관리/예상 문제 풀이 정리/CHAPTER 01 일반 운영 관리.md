## 문제 1번 : 정답 2

```
다음 (　　) 안에 들어갈 내용으로 알맞은 것은?

[jhduser@www ~]$ umask (　　)
[jhduser@www ~]$ touch a.txt
[jhduser@www ~]$ mkdir aaa
[jhduser@www ~]$ ls -l
-rw-r--r-- 1 jhduser jhduser    0 09:11 a.txt
drwxr-xr-x 2 jhduser jhduser 4096 09:22 aaa
1. 011
2. 022
3. 644
4. 755
```

```
r = 4
w = 2
x = 1
```

| 권한    |    계산 | 숫자 |
| ----- | ----: | -: |
| `---` |     0 |  0 |
| `--x` |     1 |  1 |
| `-w-` |     2 |  2 |
| `-wx` |   2+1 |  3 |
| `r--` |     4 |  4 |
| `r-x` |   4+1 |  5 |
| `rw-` |   4+2 |  6 |
| `rwx` | 4+2+1 |  7 |

그래서 문제의 결과를 보면 
```
-rw-r--r--  a.txt
drwxr-xr-x  aaa
```

숫자로 바꾸면
```
a.txt
rw- r-- r--
 6   4   4
→ 644

aaa
rwx r-x r-x
 7   5   5
→ 755
```

그런데 문제는 `umask`에 뭘 넣는냐를 묻고 있는 것입니다.  
일반적으로 새 파일과 새 디럭터리의 기본 권한은 
```
일반 파일: 666
디렉터리: 777
```
입니다.

왜 파일이 `777`이 아니냐면, 새로 만든 일반 파일에 처음부터 실행 권한 `x`를 주지 않기 때문입니다.

이제 `umask 022`를 적요하면 
```
파일 기본값
666
022  ← 막을 권한
---
644
```

그리고 디렉터리는 

```
디렉터리 기본값
777
022
---
755
```

문제 결과와 정확하게 같습니다.
```
touch a.txt
→ 644

mkdir aaa
→ 755
```

따라서 정답은 `022`입니다.

## 문제 2번 : 정답 3

```
다음과 같이 /etc/passwd 파일만으로 사용자 계정을 관리하였으나 보안상의 문제로 인해 다시 /etc/shadow 파일에서 사용자 패스워드를 관리하려고 할 때 사용하는 명령으로 알맞은 것은?

[root@www ~]# grep jhduser /etc/passwd
jhduser:$6$uBH9Mnug$dCJ9i9DJx4NfQiTDeIghpmsGw3cPGJuXgq53VotVLQ2sAuTPVN3Wc0ETPuBahNlbqUebIqX9mGHPO.G4athft/:502:502::/home/jhduser:/bin/bash
[root@www ~]#
1. pwck
2. pwunconv
3. pwconv
4. vipw
```

먼저 리눅스에게는 여러 사용자 계정이 있습니다
```
root
jhduser
yuchan
alice
```

리눅스는 이 사용자 정보를 어디엔가 저장해야 합니다.  
대표적인 파일이 `/etc/passwd` 입니다.

`/etc/passwd/`에는 한 사용자에 대한 정보가 한 줄씩 들어 있습니다.

예를 들어 `jhduser:x:502:502::/home/jhduser:/bin/bash` 이 한줄을 `:` 기준으로 나누면 다음과 같습니다.
`jhduser : x : 502 : 502 : : /home/jhduser : /bin/bash`

각각 의미는 

| 순서 | 값               | 의미            |
| -- | --------------- | ------------- |
| 1  | `jhduser`       | 사용자명          |
| 2  | `x`             | 패스워드 관련 정보    |
| 3  | `502`           | UID, 사용자 번호   |
| 4  | `502`           | GID, 기본 그룹 번호 |
| 5  | 비어 있음           | 사용자 설명        |
| 6  | `/home/jhduser` | 홈 디렉터리        |
| 7  | `/bin/bash`     | 로그인 셸         |

여기서 이번 문제에서 제일 중요한 건 **2번째 필드**입니다.  
옛날에는 `/etc/passwd/`에 패스워드 해시까지 직접 넣어서 관리했습니다.

예를 들어 문제에 나온 것처럼 `jhduser:$6$uBH9Mnug$...:502:502::/home/jhduser:/bin/bash`  
이걸 보면
```
jhduser : $6$uBH9Mnug$... : 502 : 502 : ...
          ↑
      패스워드 해시
```

즉 실제 패스워드의 해시값이 `/etc/passwd` 안에 있는 상태입니다.  
그런데 문제가 있습니다. `/etc/passwd/`는 프로그램이 사용자 정보를 확인해야 해서 일반 사용자도 읽을 수 있는 파일입니다.

그래서 현대적인 형태는 보통 이렇게 돼 
```
/etc/passwd
jhduser:x:502:502::/home/jhduser:/bin/bash
        ↑
```
여기서 `x`는 실제 패스워드는 여기 없고 `/etc/shadow`에 있습니다.

그리고 실제 패스워드 해시는 
```
/etc/shadow
jhduser:$6$uBH9Mnug$...
```
쪽에 들어가집니다.

즉 구조가 다음과 같이 변경되었습니다.
```
예전 방식

/etc/passwd
├─ 사용자명
├─ UID
├─ GID
└─ 패스워드 해시까지 저장
```

보안 때문에 

```
현재 방식

/etc/passwd
├─ 사용자명
├─ UID
├─ GID
└─ 패스워드 자리에는 x

/etc/shadow
└─ 실제 패스워드 해시 저장
```

이제 문제로 돌아가면, 문제에서는 현재 `/etc/passwd`가 다음과 같이 되어 있습니다.
```
jhduser:$6$uBH9Mnug$...:502:502::/home/jhduser:/bin/bash
```


## 문제 3번 : 정답 1

```
다음 (　　) 안에 들어갈 내용으로 알맞은 것은?

# tar ( ㄱ ) xvf mysql-boost-5.7.22.tar.gz
# tar ( ㄴ ) xvf httpd-2.2.34.tar.bz2
# tar ( ㄷ ) xvf php-5.36.tar.xz
1. ㄱ z, ㄴ j, ㄷ J
2. ㄱ z, ㄴ J, ㄷ j
3. ㄱ J, ㄴ j, ㄷ z
4. ㄱ j, ㄴ J, ㄷ z
```

## 문제 4번 : 정답 4

```
다음은 기존에 생성되어 있는 backup.tar 파일에 추가로 파일을 묶은 후 확인하는 과정이다. (　　) 안에 들어갈 내용으로 알맞은 것은?

# tar ( ㄱ ) backup.tar lin.txt joon.txt
# tar ( ㄴ ) backup.tar
1. ㄱ cvf, ㄴ tvf
2. ㄱ tvf, ㄴ rvf
3. ㄱ cvf, ㄴ rvf
4. ㄱ rvf, ㄴ tvf
```

## 문제 5번 : 정답 2

```
다음 중 프로세스의 우선순위를 변경하거나 모니터링할 때 사용하는 명령으로 틀린 것은?

1. top
2. jobs
3. nice
4. renice
```

## 문제 6번 : 정답 4

```
월, 수, 금요일 오전 9시 30분에 백업 스크립트가 동작하도록 cron을 설정하는 과정이다. 다음 (　　) 안에 들어갈 내용으로 알맞은 것은?

# vi /etc/crontab
(　　) /etc/backup.sh
1. 21 30 * * 1,3,5
2. 30 21 * * 1,3,5
3. 9 30 * * 1,3,5
4. 30 9 * * 1,3,5
```

## 문제 7번 : 정답 1

```
다음 중 사용자를 추가할 때 사용자의 홈 디렉터리에 기본으로 복사할 파일이 위치한 곳은?

1. /etc/skel
2. /etc/shadow
3. /etc/login.defs
4. /etc/default/useradd
```

## 문제 8번 : 정답 2

```
다음 설명과 같은 경우에 실행하는 명령으로 가장 알맞은 것은?

jhduser의 패스워드를 오랜 시간 변경하지 않아 임의로 패스워드를 만료 설정한다.

1. passwd -d jhduser
2. passwd -e jhduser
3. passwd -l jhduser
4. passwd -r jhduser
```

## 문제 9번 : 정답 4

```
다음 설명에 해당하는 명령으로 알맞은 것은?

애플리케이션은 /lib/libudev.so.0 라이브러리를 링크하고 있다. 실제로 이 라이브러리는 /lib/libudev.so.0.5.1을 링크하고 있다. 최근 몇 가지 버그가 수정된 /lib/libudev.so.0.5.2를 애플리케이션 수정 없이 적용하려 한다.

1. ln -h /lib/libudev.so.0 /lib/libudev.so.0.5.2
2. ln -s /lib/libudev.so.0 /lib/libudev.so.0.5.2
3. ln -h /lib/libudev.so.0.5.2 /lib/libudev.so.0
4. ln -s /lib/libudev.so.0.5.2 /lib/libudev.so.0
```

## 문제 10번 : 정답 3

```
다음 그림과 같을 때 jhduser라는 계정이 lin.txt 파일을 삭제할 수 있도록 권한을 설정하는 명령으로 알맞은 것은?

[root@www ~]# ls -ld /data
drwxr-xr-x 2 root root 4096 Jul 4 16:46 /data
[root@www ~]# ls -l /data
total 4
-rw-r--r-- 1 root root 33 Jul 4 16:47 lin.txt
[root@www ~]#
1. chmod u+w /data
2. chmod u+w lin.txt
3. chown jhduser /data
4. chown jhduser.jhduser /data
```

## 문제 11번 : 정답 1

```
다음 중 시스템 전체에서 사용자, 그룹, 다른 사용자의 퍼미션이 모두 읽기와 쓰기만 부여된 파일을 전부 찾는 명령으로 알맞은 것은?

1. find / -type f -perm 666
2. find / -type f -perm 777
3. find / -type -f -perm 666
4. find / -type -f -perm 777
```

## 문제 12번 : 정답 1

```
다음 설명에 해당하는 명령으로 알맞은 것은?

hack.c 파일을 컴파일하여 hack.o라는 목적 파일을 생성한다.

1. gcc -c hack.c
2. gcc -e hack.c
3. gcc -p hack.c
4. gcc -o hack.c
```

## 문제 13번 : 정답 1

```
다음 중 yum을 이용해서 telnet-server 패키지를 설치하는 명령으로 알맞은 것은?

1. yum install telnet-server
2. yum setup telnet-server
3. yum add telnet-server
4. yum inst telnet-server
```

## 문제 14번 : 정답 4

```
다음 중 httpd와 같이 프로세스 이름을 인자값으로 사용하는 명령으로 틀린 것은?

1. killall
2. pkill
3. nice
4. kill
```

## 문제 15번 : 정답 2

```
다음은 rpm 명령을 이용해서 httpd 패키지를 제거하는 과정이다. (　　) 안에 들어갈 내용으로 알맞은 것은?

# rpm ( ㄱ ) httpd
error: Failed dependencies:
        httpd >= 2.2.0 is needed by (installed)
        gnome-user-share-2.28.2-3.el6.i686
# rpm ( ㄱ ) httpd ( ㄴ )
1. ㄱ -e, ㄴ --force
2. ㄱ -e, ㄴ --nodeps
3. ㄱ -r, ㄴ --nodeps
4. ㄱ -r, ㄴ --force
```

## 문제 16번 : 정답 2

```
다음 중 백그라운드에서 실행 중인 작업을 포어그라운드로 전환하는 명령으로 알맞은 것은?

1. bg
2. fg
3. jobs
4. psgrep
```

## 문제 17번 : 정답 4

```
다음 중 파일의 시간 정보를 현재 시간으로 변경하는 명령으로 알맞은 것은?

1. ls
2. info
3. stat
4. touch
```

## 문제 18번 : 정답 2

```
vi 편집기를 이용하여 사용자 혹은 그룹의 디스크 쿼터를 설정하는 명령으로 알맞은 것은?

1. quota
2. edquota
3. repquota
4. quotacheck
```

## 문제 19번 : 정답 4

```
시스템에 로그인한 사용자의 ID들을 출력하는 명령으로 알맞은 것은?

1. w
2. who
3. whoami
4. lslogins
```

## 문제 20번 : 정답 3

```
다음 설명 중 올바른 root 사용자 관리 기법으로 틀린 것은?

1. 무의미하게 장시간 로그인되어 있지 않도록 한다.
2. root 이외에 UID가 0인 사용자를 생성하지 않는다.
3. SSH로 접근 시 root로 직접 로그인을 허용하면 관리가 편리하다.
4. PAM(Pluggable Authentication Modules)을 이용해서 접근을 제어한다.
```
