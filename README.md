# Environment Getter

Getting environment variable with typed value.

## Installation

```bash
go get -u github.com/metadiv-io/env
```

## Getting String

```go
// get string
s := env.String("ENV_KEY", "default")

// get string list
sl := env.StringList("ENV_KEY", []string{"default1", "default2"})

// must get string
ms := env.MustString("ENV_KEY")

// must get string list
msl := env.MustStrings("ENV_KEY")
```

## Getting Int

```go
// get int
i := env.Int("ENV_KEY", 0)
i8 := env.Int8("ENV_KEY", 0)
i16 := env.Int16("ENV_KEY", 0)
i32 := env.Int32("ENV_KEY", 0)
i64 := env.Int64("ENV_KEY", 0)

// get int list
il := env.IntList("ENV_KEY", []int{0, 1, 2})
il8 := env.Int8List("ENV_KEY", []int8{0, 1, 2})
il16 := env.Int16List("ENV_KEY", []int16{0, 1, 2})
il32 := env.Int32List("ENV_KEY", []int32{0, 1, 2})
il64 := env.Int64List("ENV_KEY", []int64{0, 1, 2})

// must get int
mi := env.MustInt("ENV_KEY")
mi8 := env.MustInt8("ENV_KEY")
mi16 := env.MustInt16("ENV_KEY")
mi32 := env.MustInt32("ENV_KEY")
mi64 := env.MustInt64("ENV_KEY")

// must get int list
mil := env.MustInts("ENV_KEY")
mil8 := env.MustInt8s("ENV_KEY")
mil16 := env.MustInt16s("ENV_KEY")
mil32 := env.MustInt32s("ENV_KEY")
mil64 := env.MustInt64s("ENV_KEY")
```

## Getting Uint

```go
// get uint
u := env.Uint("ENV_KEY", 0)
u8 := env.Uint8("ENV_KEY", 0)
u16 := env.Uint16("ENV_KEY", 0)
u32 := env.Uint32("ENV_KEY", 0)
u64 := env.Uint64("ENV_KEY", 0)

// get uint list
ul := env.UintList("ENV_KEY", []uint{0, 1, 2})
ul8 := env.Uint8List("ENV_KEY", []uint8{0, 1, 2})
ul16 := env.Uint16List("ENV_KEY", []uint16{0, 1, 2})
ul32 := env.Uint32List("ENV_KEY", []uint32{0, 1, 2})
ul64 := env.Uint64List("ENV_KEY", []uint64{0, 1, 2})

// must get uint
mu := env.MustUint("ENV_KEY")
mu8 := env.MustUint8("ENV_KEY")
mu16 := env.MustUint16("ENV_KEY")
mu32 := env.MustUint32("ENV_KEY")
mu64 := env.MustUint64("ENV_KEY")

// must get uint list
mul := env.MustUints("ENV_KEY")
mul8 := env.MustUint8s("ENV_KEY")
mul16 := env.MustUint16s("ENV_KEY")
mul32 := env.MustUint32s("ENV_KEY")
mul64 := env.MustUint64s("ENV_KEY")
```

## Getting Float

```go
// get float
f32 := env.Float32("ENV_KEY", 0)
f64 := env.Float64("ENV_KEY", 0)

// get float list
fl32 := env.Float32List("ENV_KEY", []float32{0, 1, 2})
fl64 := env.Float64List("ENV_KEY", []float64{0, 1, 2})

// must get float
mf32 := env.MustFloat32("ENV_KEY")
mf64 := env.MustFloat64("ENV_KEY")

// must get float list
mfl32 := env.MustFloat32s("ENV_KEY")
mfl64 := env.MustFloat64s("ENV_KEY")
```

## Getting Bool

```go
// get bool
b := env.Bool("ENV_KEY", false)

// get bool list
bl := env.BoolList("ENV_KEY", []bool{false, true})

// must get bool
mb := env.MustBool("ENV_KEY")

// must get bool list
mbl := env.MustBools("ENV_KEY")
```

## Getting Complex

```go
// get complex
c64 := env.Complex64("ENV_KEY", 0)
c128 := env.Complex128("ENV_KEY", 0)

// get complex list
cl64 := env.Complex64List("ENV_KEY", []complex64{0, 1, 2})
cl128 := env.Complex128List("ENV_KEY", []complex128{0, 1, 2})

// must get complex
mc64 := env.MustComplex64("ENV_KEY")
mc128 := env.MustComplex128("ENV_KEY")

// must get complex list
mcl64 := env.MustComplex64s("ENV_KEY")
mcl128 := env.MustComplex128s("ENV_KEY")
```
