# Gin の Docker プロジェクトのサンプル

## アーキテクチャーの選択

各自の環境に応じて使用するベースイメージのアーキテクチャを選択する。

Dockerfile

```Dockerfile
# amdかarmを選択する
# FROM golang:1.19.3-alpine
FROM arm64v8/golang:1.19.3-alpine
```

```Dockerfile
ARG CGO_ENABLED=0
ARG GOOS=linux
# ARG GOARCH=amd64
ARG GOARCH=arm64
```

docker-compose.yml

```yml
services:
  db:
    # amdかarmを選択する
    # image: postgres:15.1-alpine
    image: arm64v8/postgres:15.1-alpine
```

## ビルド

```
docker compose build
```

## 起動

```
docker compose up -d
```

## パッケージ追加

- パッケージインストール

```
go get -u <パッケージ名>
```

- Dockre image の再ビルド

```
docker compose build
```

## ホットリロード

プログラムを編集し、保存すると自動で Go の再ビルドと実行が行われる。

# 注意

衝動的に作ったので、不備や間違ってる部分がある可能性があります。また、DB の設定や API の基本的な記述はまだありません。
