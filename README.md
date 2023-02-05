## slinter

シンプルな Lint 機能を提供するパッケージ

## 使用方法

```bash
go install github.com/aoman-n/slinter/cmd/slinter
go vet -vettool=$(which slinter) ./...
```

## 機能

- 関数/メソッドの行数チェック
- 関数/メソッドの引数チェック

## フラグ

`-{analysis名}.{flag名}`

| flag 名  | default | description |
| :------- | :------ | :---------- |
| maxLines | 150     | 最大行数    |
| maxArgs  | 4       | 最大引数    |

```bash
go vet -vettool=$(which slinter) -slinter.maxLines=200 -slinter.maxArgs=8 ./...
```

## 開発時実行方法　

### テストでの実行

```bash
go test
```

### バイナリでの実行

```bash
go install ./cmd/slinter
go vet -vettool=$(which slinter) ./...
```
