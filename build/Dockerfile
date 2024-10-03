# Goのバージョンを指定したベースイメージを使用
FROM golang:1.23.2-alpine

# 必要なパッケージのインストール
RUN apk update && apk add --no-cache git

# 環境変数の設定（Goのモジュール管理のため）
ENV GO111MODULE=on

# ホスト上のコードをコンテナにコピー
WORKDIR /app
COPY . .

# 開発中に便利なツールのインストール
RUN go install golang.org/x/tools/cmd/goimports@latest
