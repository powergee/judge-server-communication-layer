# judge-server-communication-layer

DMOJ의 judge-server와 Golang으로 통신하기 위한 테스트용 리포지토리입니다.

judge-server 프로세스에 Ping을 전송하거나, 채점을 명령하거나, 현재 채점 현황을 얻는 등의 작업을 할 수 있습니다.

# judge-server 실행 및 통신 방법

*아래의 과정은 https://docs.dmoj.ca/#/judge/setting_up_a_judge?id=judge-side-setup 의 내용을 참조하여 작성하였습니다.*

judge-server는 윈도우를 지원하지 않고 리눅스 상에서 동작하므로, 아래 과정은 리눅스에서 실행하시거나, 윈도우의 WSL 상에서 실행하셔야 합니다.

## judge-server 설정 파일 만들기

judge-server를 설치 및 실행하기에 앞서, 특정 디렉토리에 judge-server의 세부 설정을 담고 있는 yml 파일을 만들고, 같은 디렉토리에 채점할 문제의 입출력 파일을 준비해야 합니다. **이 문서에서는 `/problems/judge.yml` 파일에 세부 설정을 작성하고, `/problems` 디렉토리에 각 문제들의 입출력 파일을 만드는 것으로 하겠습니다.**

### 1. judge.yml 파일 만들기

`judge.yml` 파일은 judge-server의 세부 설정을 담고 있는 파일로, 아래와 같은 정보를 담고 있습니다.

1. `id`: 채점기의 ID
2. `key` 채점기의 인증용 키 (단, 이 정보는 이 프로젝트에서는 활용하지 않으므로 임의로 작성하셔도 됩니다.)
3. `problem_storage_root`: 각 문제의 입출력 데이터를 저장하고 있는 디렉토리 경로

예를 들어, `/problems/judge.yml`의 내용을 아래와 같이 작성할 수 있습니다.

```yml
id: TestJudge
key: key_for_testing
problem_storage_root:
  - /problems
```

### 2. 채점할 문제들의 입출력 데이터 추가

각 문제들의 입출력 데이터는 위에서 설정한 `problem_storage_root` 경로(이 문서의 예시에서는 `/problems` 폴더) 아래에 저장해야 합니다. 저장 형식은 아래와 같습니다.

예를 들어, [이 문제의](https://dmoj.ca/problem/aplusb) 문제 데이터를 "aplusb"라는 문제 ID로 저장하고 싶다고 가정 하겠습니다. 문제 ID가 "aplusb"인 문제의 입출력 데이터를 저장하려면, `/problems/aplusb` 디렉토리를 만들고, 그 안에 아래와 같은 파일들을 작성해야 합니다.

1. `{입출력 데이터 번호}.in`: 입력 데이터를 나타냅니다.
    - `1.in`의 내용을 아래와 같이 작성할 수 있습니다.
    ```
    2
    1 1
    -1 0
    ```
2. `{입출력 데이터 번호}.out`: 출력 데이터를 나타냅니다.
    - `1.out`의 내용을 아래와 같이 작성할 수 있습니다.
    ```
    2
    -1
    ```
3. `init.yml`: 테스트케이스 목록을 정의합니다.
    - `1.in`, `1.out`, `2.in`, `2.out`와 같은 입출력 데이터가 있다면, `init.yml` 파일의 내용은 아래와 같습니다.
    ```yml
    test_cases:
    - {in: 1.in, out: 1.out, points: 0}
    - {in: 2.in, out: 2.out, points: 0}
    ```
    - 모든 `points` 값을 0으로 설정하면, 채점 도중 단 하나의 테스트 케이스에서라도 오답이 생기면 즉시 채점을 중단하게 됩니다. 반면에 `points` 값이 0이 아닌 값으로 설정되어있다면 하나가 틀리더라도 모든 테스트 케이스를 채점합니다.

## judge-server 실행

실행 방법에는 Docker로 실행하는 방법과 Python 3의 `pip`를 통해 실행하는 방법이 있는데, 여기에서는 Docker로 실행하는 방법을 소개하겠습니다. `pip`로 실행하는 방법은 다음 링크를 참고해주시기 바랍니다.

- https://docs.dmoj.ca/#/judge/setting_up_a_judge?id=judge-side-setup

### Docker 설치

실행 환경에 맞게 Docker를 설치해주세요.

- Ubuntu: https://docs.docker.com/engine/install/ubuntu/
- WSL 2: https://docs.microsoft.com/ko-kr/windows/wsl/tutorials/wsl-containers

### make 설치

`make` 명령 실행을 위해 `sudo apt install gcc make` 를 실행해주세요.

### Docker 컨테이너 빌드

judge-server 리포지토리를 clone 하여 컨테이너를 빌드해야 합니다.

```bash
git clone --recursive https://github.com/DMOJ/judge.git
cd judge/.docker
sudo make judge-tier1
```

이 예시에서는 `judge-tier1`을 빌드하고 있는데, tier 1은 tier 1, 2, 3 중 가장 지원하는 언어 수가 적은 컨테이너입니다. 만약 더 많은 언어를 채점하려면 `judge-tier2`, `judge-tier3`를 고려해야 합니다. 지원하는 언어의 구체적인 목록은 아래 Dockerfile에서 확인할 수 있습니다.

- judge-tier1: https://github.com/DMOJ/runtimes-docker/blob/master/tier1/Dockerfile
- judge-tier2: https://github.com/DMOJ/runtimes-docker/blob/master/tier2/Dockerfile
- judge-tier3: https://github.com/DMOJ/runtimes-docker/blob/master/tier3/Dockerfile

## judge-server-communication-layer 실행

통신 레이어를 먼저 실행합니다.

```
go run main.go
```

이렇게 하면 `Listening:9999...` 라는 메세지와 함께 9999번 포트에서 judge-server가 연결할 때까지 블록되게 됩니다.

## judge-server 컨테이너 실행

위에서 빌드한 Docker 컨테이너를 실행합니다.

```
sudo docker run \
    --name judge_name \
    --network host \
    -v /home/hyeon/problems:/problems \
    --cap-add=SYS_PTRACE \
    -d \
    --restart=always \
    dmoj/judge-tier1:latest \
    run -p "9999" -c /problems/judge.yml "localhost"
```

이렇게 실행하면 컨테이너가 `judge_name`이라는 이름으로 실행되고, 나중에 이 컨테이너를 종료하거나 삭제하거나 재시작해야 할 일이 있다면 `judge_name`이라는 이름으로 접근할 수 있습니다.

정상적으로 실행되었다면 약 30초~1분 내에 judge-server-communication-layer와 연결되고, 통신 레이어를 실행하고 있는 표준 출력에 아래와 같은 내용이 나타납니다.

```
Handshaking is succeeded!

1. ping
2. get-current-submission
3. submission-request
4. terminate-submission
5. disconnect
```

이제부터 judge-server를 직접 제어할 수 있습니다.