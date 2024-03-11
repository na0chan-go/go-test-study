# go-test-study

# テストコードの書き方

## カバレッジファイルの生成
`go test -cover ./... -coverprofile=cover`

## カバレッジの可視化
`go tool cover -html=cover -o cover.html`
