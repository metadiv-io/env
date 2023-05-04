# env

## Installation

```bash
go get -u github.com/metadiv-io/env
```

## Highlights

### One variable

* env.String(key string, defaultValue ...string) string

* env.Int(key string, defaultValue ...int) int

* env.Int64(key string, defaultValue ...int64) int64

* env.Int32(key string, defaultValue ...int32) int32

* env.Int16(key string, defaultValue ...int16) int16

* env.Int8(key string, defaultValue ...int8) int8

* env.Float64(key string, defaultValue ...float64) float64

* env.Float32(key string, defaultValue ...float32) float32

* env.Uint(key string, defaultValue ...uint) uint

* env.Uint64(key string, defaultValue ...uint64) uint64

* env.Uint32(key string, defaultValue ...uint32) uint32

* env.Uint16(key string, defaultValue ...uint16) uint16

* env.Uint8(key string, defaultValue ...uint8) uint8

* env.Bool(key string, defaultValue ...bool) bool

* env.Complex128(key string, defaultValue ...complex128) complex128

* env.Complex64(key string, defaultValue ...complex64) complex64

* env.Byte(key string, defaultValue ...byte) byte

* env.Rune(key string, defaultValue ...rune) rune

### Multiple variables

* env.Strings(key string, defaultValue ...string) []string

* env.Ints(key string, defaultValue ...int) []int

* env.Int64s(key string, defaultValue ...int64) []int64

* env.Int32s(key string, defaultValue ...int32) []int32

* env.Int16s(key string, defaultValue ...int16) []int16

* env.Int8s(key string, defaultValue ...int8) []int8

* env.Float64s(key string, defaultValue ...float64) []float64

* env.Float32s(key string, defaultValue ...float32) []float32

* env.Uints(key string, defaultValue ...uint) []uint

* env.Uint64s(key string, defaultValue ...uint64) []uint64

* env.Uint32s(key string, defaultValue ...uint32) []uint32

* env.Uint16s(key string, defaultValue ...uint16) []uint16

* env.Uint8s(key string, defaultValue ...uint8) []uint8

* env.Bools(key string, defaultValue ...bool) []bool

* env.Complex128s(key string, defaultValue ...complex128) []complex128

* env.Complex64s(key string, defaultValue ...complex64) []complex64

* env.Bytes(key string, defaultValue ...byte) []byte

* env.Runes(key string, defaultValue ...rune) []rune
