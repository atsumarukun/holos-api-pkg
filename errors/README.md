# エラーパッケージ

errorsパッケージはエラーコード及びスタックトレースを保持したエラーを提供します.

## 目次

- [目次](#目次)
- [クイックスタート](#クイックスタート)
- [エラーの初期化](#エラーの初期化)
- [エラー発生箇所の特定](#エラー発生箇所の特定)

## クイックスタート

`github.com/atsumarukun/holos-api-pkg`を依存モジュールに追加し、errorsパッケージをimportして利用します.

```golang
package main

import (
  "fmt"

  "github.com/atsumarukun/holos-api-pkg/errors"
)

func main() {
  err := errors.New(errors.CodeUnknown, "example error")
  fmt.Printf("%+v", err)
  // UNKNOWN: example error
  //
  // /workspace/main.go:10
  // ....
}
```

## エラーの初期化

errors.New関数はエラーコードとエラーメッセージを引数に受け取り、スタックトレースを保持したエラーを返します.

```golang
err := errors.New(errors.codeUnknown, "example error")
```

## エラー発生箇所の特定

フォーマットのverbに%sまたは%vを指定すると、エラーコードとエラーメッセージを結合して返却します.

```golang
fmt.Printf("%s", err)
// UNKNOWN: example error
```

フォーマットのverbに+タグ付きの%+vを指定すると、スタックトレース付きで返却します.
```golang
fmt.Printf("%+v", err)
// UNKNOWN: example error
//
// /workspace/main.go:10
// ....
```
