# syntax=docker/dockerfile:1.4

# Stage 1: Build Stage
FROM golang:1.23 AS builder

WORKDIR /app

# go.modとgo.sumをマウントして依存関係をインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをマウントしてビルド
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 go build -o /bin/server ./cmd

RUN ls /bin/ && ls  

# Stage 2: Runtime Stage
FROM golang:1.23-alpine

WORKDIR /app

# ビルドされた実行ファイルをコピー
COPY --from=builder /bin/server /bin/

# 実行ファイルを起動
CMD ["/bin/server"]
