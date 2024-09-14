# Web API Variety Implementation

## 概要

Go, TypeScript, Rustでそれぞれ簡易なWebAPIを実装

## プロジェクト構成

- **go-web-app/**: Goで実装されたWeb API
- **rust-web-app/**: Rustで実装されたWeb API
- **ts-web-app/**: TypeScript（Node.jsとExpressを使用）で実装されたWeb API

## アーキテクチャの説明

### Go Web アプリケーション

**アーキテクチャ**: レイヤードアーキテクチャ  
**説明**:  
Goでの実装は、レイヤードアーキテクチャ的なデザインを採用した。具体的には、以下の層で構成されている：

- **Handlers（ハンドラー層）**: HTTPリクエストを受け取り、レスポンスを生成する。`handlers`ディレクトリ内で定義
- **Services（サービス層）**: ビジネスロジックを実装している`services`ディレクトリ内に定義
- **Models（モデル層）**: アプリケーションで使用されるデータ構造を定義する。`models`ディレクトリ内に作成

コードのモジュール性と拡張性を高め、保守性の高さを目指している

### Rust Web アプリケーション

**アーキテクチャ**: シンプルレイヤードアーキテクチャ  
**説明**:  
Rustでの実装は、レイヤードアーキテクチャに近いシンプルな構成を採用しており、以下の層で構成されている：

- **Handlers（ハンドラー層）**: `actix-web`を使用して、ルートごとにハンドラー関数を定義する
- **Models（モデル層）**: `serde`を用いて、データモデルを定義し、JSONシリアライゼーションをサポートする

Rustでは、アプリケーション全体の性能と安全性を確保するため、型システムと所有権モデルを活用した

### TypeScript Web アプリケーション

**アーキテクチャ**: MVCアーキテクチャ  
**説明**:  
TypeScriptでの実装は、MVC（Model-View-Controller）アーキテクチャに基づいている。以下の要素で構成されている：

- **Models（モデル層）**: データ構造を定義する。`models`ディレクトリに格納
- **Controllers（コントローラー層）**: ビジネスロジックを担当し、リクエストに応じたレスポンスを返す。`controllers`ディレクトリに格納
- **Services（サービス層）**: ビジネスロジックを実行し、コントローラー層から呼び出される。`services`ディレクトリに格納

MVCアーキテクチャにより、各層が分離されているため、アプリケーションの拡張や保守が容易

## セットアップおよび実行手順

### Go版

1. `go-web-app`ディレクトリに移動

```bash
cd go-web-app
```

2. Go Moduleの初期化

```bash
go mod init go-web-app
```

3. アプリケーション起動

```bash
go run main.go
```

### Rust版

rust-web-appディレクトリに移動

```bash
cd rust-web-app
```

アプリケーションをビルドする

```bash
cargo build
```

アプリケーションを実行する

```bash
cargo run
```

### TypeScript版

ts-web-appディレクトリに移動

```bash
cd ts-web-app
```

パッケージインストール

```bash
npm install
```

アプリケーションを実行する

```bash
npm start
```

## API エンドポイント

各アプリケーションは、デモンストレーションのために同様のエンドポイントを実装している

- **GET /**: シンプルな挨拶メッセージ
- **GET /user**: JSON形式でユーザー情報を返す

### APIレスポンス例

- **Go版**:
  - `GET /`: `Hello, Go!`
  - `GET /user`: `{ "id": 1, "name": "Alice" }`

- **Rust版**:
  - `GET /`: `Hello, Rust!`
  - `GET /user`: `{ "id": 1, "name": "Alice" }`

- **TypeScript版**:
  - `GET /`: `Hello, TypeScript!`
  - `GET /user`: `{ "id": 1, "name": "Alice" }`
