# go-test-study

## テストコードの書き方

## カバレッジファイルの生成

`go test -cover ./... -coverprofile=cover.out`

## カバレッジの可視化

`go tool cover -html=cover.out -o cover.html`
