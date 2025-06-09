# 1단계: 빌드 단계
FROM golang:1.24-alpine AS builder

# 필요한 OS 패키지 설치
RUN apk add --no-cache git ca-certificates
# go mod로 Github 등에서 패키지를 다운받기 위함

# 작업 디렉토리 설정
WORKDIR /app

# go mod 복사 및 의존성 설치
COPY go.mod go.sum ./
RUN go mod download

# 소스 복사 및 빌드
COPY . .
RUN go build \
  -ldflags="-s -w" \
  -trimpath \
  -o register ./cmd/register
# -s : 디버그 심볼 제거 (symbol table) -> 디버깅에 쓰이는 메타데이터 제거
# -w : DWARF 디버깅 정보 제거 -> gdb 등에서 쓰이는 정보 제거
# => Panic 발생 시 stack trace에 정확한 파일명, 줄번호가 안나오게 됨
# trimpath : Go는 기본적으로 빌드한 실행파일에 개발자의 파일 경로를 포함시킴
#            -> 각 개발자/CI 환경마다 경로 다르면 빌드 결과도 다름
#              -> 다른 경로에서 빌드해도 동일한 결과 생성이 가능케함
#            - 보안 이슈 : 소스 구조 노출
# ./cmd/register에 있는 main.go를 컴파일해서 실행파일 register 생성

# 2단계: 실행용 이미지 (더 작음)
FROM alpine:3.20
# 실행용 바이너리만 필요하므로 최소한의 실행 환경만 갖춤

WORKDIR /app

# 기타 실행에 필요한 것만 복사
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/register /app/

EXPOSE 8080

# 실행
CMD ["./register"]

# ##################################
# ########## 이후에 해야 할일 ##########
# 1. scratch 이미지로 하는 방법이 뭔지 알아보기
# 2. 그와 관련된 지식 학습하기 (libc? GCO_ENABLED?)