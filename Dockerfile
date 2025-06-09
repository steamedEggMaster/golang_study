# 1단계: 빌드 단계
FROM golang:1.24-alpine AS builder

# 필요한 OS 패키지 설치
RUN apk add --no-cache git
# go mod로 Github 등에서 패키지를 다운받기 위함

# 작업 디렉토리 설정
WORKDIR /app

# go mod 복사 및 의존성 설치
COPY go.mod go.sum ./
RUN go mod download

# 소스 복사 및 빌드
COPY . .
RUN go build -o register ./cmd/register
# ./cmd/register에 있는 main.go를 컴파일해서 실행파일 register 생성

# 2단계: 실행용 이미지 (더 작음)
FROM alpine:3.20
# 실행용 바이너리만 필요하므로 최소한의 실행 환경만 갖춤

WORKDIR /app

RUN apk add --no-cache ca-certificates
# HTTPS 요청에 필요한 CA 인증서들 설치

# godotenv나 기타 실행에 필요한 것만 복사
COPY --from=builder /app/register /app/
# 빌드 스테이지에서 생성된 실행파일만 복사
COPY config/env ./config/env
# 환경 설정 파일들을 컨테이너에 포함시킴

# 환경변수 기본값 설정 (없으면 dev로 동작)
ENV ENV=dev

EXPOSE 8080

# 실행
CMD ["./register"]