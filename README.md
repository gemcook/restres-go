# restres-go

REST APIの定型レスポンスを提供します。

## インストール

次のコマンドでインストールします。

```sh
go get "github.com/gemcook/restres-go/erres"
```

`dep` を使う場合、は次のコマンドを使用します。

```sh
dep ensure -add "github.com/gemcook/restres-go/erres"
```

## 使用例

1. エラーコード、エラータイプ、エラーメッセージを渡して、`New`します。

    ```go
    res := erres.New("Unexpected", erres.ErrorUnknown, "something wrong")
    ```

2. エラーメッセージは複数渡すことが出来ます。
    さらに、`string`, `error`, `fmt.Stringer`に限らず、どのような型も渡すことが出来ます。

    ```go
    res := erres.New("Unexpected", erres.ErrorUnknown, "something wrong", errors.New("segmentation fault"), errorObject{123})
    ```

3. エラーを複数渡すことも出来ます。

    ```go
    res := erres.New("Unexpected", erres.ErrorUnknown, "something wrong")
    res.Append("NetworkChanged", erres.ErrorUnknown, "network change detected")
    ```
