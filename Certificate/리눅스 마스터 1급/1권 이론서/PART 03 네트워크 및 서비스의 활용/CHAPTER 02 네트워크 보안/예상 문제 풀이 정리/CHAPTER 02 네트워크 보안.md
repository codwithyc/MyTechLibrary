## 문제 1번 : 정답 1

```

네트워크 침해 공격 기법에 대한 설명 중 틀린 것은?

1. Sniffing을 위해 네트워크 모드를 Normal 모드로 변경해야 한다.
2. Spoofing은 IP·MAC 주소 등의 정보를 변조해 역추적을 어렵게 한다.
3. TCP SYN Flooding은 TCP의 3-way handshaking을 악용한다.
4. Mail Bomb은 동일 이메일 주소를 대상으로 대량 메일을 동시에 발송한다.

```

## 문제 2번 : 정답 1

```

공격 대상에 IP 패킷을 보낼 때 발신자와 수신자 IP를 모두 공격 대상의 IP로 설정하는 DoS 공격은?

1. Land Attack
2. Teardrop Attack
3. Ping of Death
4. NTP 증폭 공격

```

## 문제 3번 : 정답 2

```

2개의 네트워크 카드를 내장한 Bastion으로 외부와 내부 네트워크를 각각 연결하고 내부 IP가 노출되지 않도록 하는 방화벽은?

1. 배스천 호스트
2. 듀얼(이중) 홈 게이트웨이
3. 스크린 호스트 게이트웨이
4. 스크린 서브넷 게이트웨이

```

## 문제 4번 : 정답 3

```

IDS(Intrusion Detection System)에 대한 설명 중 틀린 것은?

1. 네트워크 및 서비스에 대한 공격을 실시간으로 탐지한다.
2. HIDS는 서버에 직접 설치되어 다양한 정보를 이용해 침입을 탐지한다.
3. 오용 탐지는 정상 패턴과 비교해 이상한 패턴을 공격으로 판단한다.
4. 데이터를 수집·분석해 침입 탐지, 추적 및 보고 기능을 제공한다.

```

## 문제 5번 : 정답 1

```

리눅스에서 사용할 수 있는 공개형 IDS로 스니퍼, 전처리기, 탐지 엔진, 로깅으로 구성된 소프트웨어는?

1. Snort
2. iptables
3. ifconfig
4. John the Ripper

```

## 문제 6번 : 정답 3

```

다음 Snort 룰에 대한 설명으로 잘못된 것은?

alert tcp any any -> 192.168.10.0/24 80
(msg:"passwd detected"; content:"passwd"; nocase; sid:1000001;)

1. 점검하는 프로토콜은 TCP이다.
2. 192.168.10.0/24의 80번 포트로 유입되는 패킷을 점검한다.
3. PASSWD 문자열은 점검하지 않는다.
4. 룰의 식별자는 1000001이다.

```

## 문제 7번 : 정답 2

```

라우터로 사용되는 호스트를 통과하며, 해당 호스트가 목적지가 아닌 패킷을 대상으로 하는 iptables 체인은?

1. INPUT
2. FORWARD
3. OUTPUT
4. PREROUTING

```

## 문제 8번 : 정답 4

```

iptables 정책 파일을 읽어 적용하는 명령으로 알맞은 것은?

1. iptables-restore firewall.sh
2. iptables-save -s firewall.sh
3. iptables-save > firewall.sh
4. iptables-restore < firewall.sh

```

## 문제 9번 : 정답 2

```

다음 iptables 설정에 대한 설명으로 잘못된 것은?

iptables -t nat -A POSTROUTING -o eth0 -j SNAT --to 222.235.10.7

1. -t nat으로 nat 테이블을 선택한다.
2. Destination IP를 변경하는 설정이다.
3. -o eth0은 eth0으로 나가는 패킷을 대상으로 한다.
4. -A POSTROUTING은 POSTROUTING 체인에 정책을 추가한다.

```

## 문제 10번 : 정답 3

```

iptables의 -j 옵션과 함께 사용하는 항목 설명으로 잘못된 것은?

1. ACCEPT : 패킷 허가
2. REJECT : 패킷을 거부하며 TCP는 Reset 메시지 전송
3. DROP : 패킷을 거부하며 UDP는 Port Unreachable 메시지 전송
4. LOG : 패킷을 syslog에 전달해 기록

```

## 문제 11번 : 정답 2

```

리눅스 방화벽에 대한 설명으로 잘못된 것은?

1. 최근 리눅스는 firewalld를 방화벽 규칙 관리 데몬으로 이용한다.
2. firewalld 규칙 적용을 위해 반드시 서비스를 재시작해야 한다.
3. firewalld는 영역(Zone)과 서비스(Service)로 규칙을 관리한다.
4. firewall-cmd로 방화벽 규칙을 설정할 수 있다.

```

## 문제 12번 : 정답 1

```

firewall-cmd 옵션 설명으로 잘못된 것은?

1. --permanent : 영구 규칙을 설정하지만 --reload를 실행하면 제거됨
2. --timeout : 정해진 시간 동안만 규칙 유지
3. --zone : 규칙을 적용할 영역 지정
4. --list-all : 사용 가능한 서비스·포트 목록 출력

```

## 문제 13번 : 정답 4

```

다음 firewall-cmd 명령 설명으로 잘못된 것은?

firewall-cmd --add-forward-port=port=80:proto=tcp:toport=8080

1. firewall-cmd로 포트 포워딩을 설정한다.
2. TCP 80번 포트를 대상으로 한다.
3. 포워딩 대상 포트는 8080이다.
4. --zone을 생략했으므로 자동으로 Trusted 영역이 대상이 된다.

```

## 문제 14번 : 정답 4

```

다음 firewall-cmd 옵션 설명으로 잘못된 것은?

1. --add-service=http : HTTP 서비스 추가
2. --add-port=21/tcp : TCP 21번 포트 추가
3. --remove-service=https : HTTPS 서비스 제거
4. --add-port=8080-8090/tcp : 연속되지 않는 형식으로 오류 발생

```
