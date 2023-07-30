# 核心庫: env

## 解決問題

Go 原生庫 `os` 獲取的環境變數默認的變數類型只有 `string`， `env` 提供多類型的環境變數獲取，並提供預設值設置，以及檢查必須設罝的環境變數的功能。

## 安裝

```bash
go get -u github.com/metadiv-io/env
```

## 支援類型

* string
* int
* int64
* int32
* int16
* int8
* float64
* float32
* uint
* uint64
* uint32
* uint16
* uint8
* bool
* complex128
* complex64

## 使用

```go
package main

import (
    "fmt"
    "github.com/metadiv-io/env"
)

func main() {
    v := env.String("STRING_VAR", "default") // with default
    v1 := env.MustString("STRING_VAR") // panic if STRING_VAR is not set
    vs := env.Strings("STRING_VARS", []string{"default"}) // list with default
    vs1 := env.MustStrings("STRING_VARS") // panic if STRING_VAR is not set
}
```
