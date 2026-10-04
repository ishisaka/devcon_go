# miseとDev Containerで作るGo開発環境

このリポジトリは、miseでGoのバージョンを管理し、Visual Studio CodeのDev Containers拡張機能を使って開発できる小さなGoアプリケーションです。Dev Containerを使うと、必要なツールをコンテナ内にそろえ、ホスト環境へのインストールを最小限にして開発できます。

## 前提条件

- Dockerを実行できる環境（Docker Desktopなど）
- Visual Studio Code
- Visual Studio Codeの「Dev Containers」拡張機能

## 開発環境の設定

[`.devcontainer/devcontainer.json`](./.devcontainer/devcontainer.json) はUbuntuベースの開発コンテナを定義しています。コンテナにはmise、Git、GitHub CLIの各Dev Container Featureが追加され、VS Code拡張機能としてmiseとGoがインストールされます。

[`mise.toml`](./mise.toml) でGo 1.27.1を指定しています。コンテナ内ではmiseがこの設定を読み取り、Goのツールチェーンを管理します。Goモジュールの宣言は[`go.mod`](./go.mod)にあります。

## プロジェクト構成

```text
.
├── cmd/
│   └── main.go          # アプリケーションのエントリーポイント
├── pkg/
│   └── hello/
│       └── hello.go     # 挨拶メッセージを返すパッケージ
├── .devcontainer/
│   └── devcontainer.json
├── mise.toml            # 開発ツールのバージョン設定
└── go.mod               # Goモジュールの定義
```

`cmd/main.go`が`pkg/hello`を呼び出して挨拶を表示します。再利用する処理は`pkg`以下に置き、ユーザーへの出力はコマンドのエントリーポイントで行う構成です。

## 始め方

1. このリポジトリをクローンします。

```sh
git clone https://github.com/ishisaka/devcon_go.git
cd devcon_go
```

2. Visual Studio Codeでフォルダーを開きます。

3. コマンドパレットから **Dev Containers: Reopen in Container** を実行します。

4. コンテナが起動したら、ターミナルで以下のコマンドを実行し、miseを使ってGoやツールのインストールを行います。

```sh
mise install
```

5. `.bashrc`を編集して、miseの初期化スクリプトを読み込むようにします。

```sh
echo 'eval "$(/usr/local/bin/mise activate bash)"' >> ~/.bashrc
source ~/.bashrc
```

6. goのツールチェーンが正しくインストールされているか確認します。

```sh
go version
```

7. アプリケーションを実行して、正しく動作するか確認します。

```sh
go run ./cmd/main.go
```

