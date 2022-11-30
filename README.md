# Gin の Docker プロジェクトのサンプル

## .env ファイルの作成

`.env.sample`をコピーし、コピーしたファイルの名前を`.env`に書き換える

## アーキテクチャーの選択

各自の環境に応じて使用するベースイメージのアーキテクチャを選択する。

Dockerfile
```Dockerfile
ARG CGO_ENABLED=0
ARG GOOS=linux
# アーキテクチャを指定する
# ARG GOARCH=amd64
ARG GOARCH=arm64
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
