package env

import (
	"os"
	"strconv"
	"strings"
)

func newManager[T any](key string, defaultValue ...T) *manager[T] {
	var d T
	var has bool
	if len(defaultValue) > 0 {
		d = defaultValue[0]
		has = true
	}
	return &manager[T]{
		Key:          key,
		Value:        newValue(os.Getenv(key)),
		DefaultValue: d,
		HasDefault:   has,
		Must:         false,
	}
}

func newMustManager[T any](key string) *manager[T] {
	m := newManager[T](key)
	m.Must = true
	if os.Getenv("GIN_MODE") == "release" || os.Getenv("GIN_MODE") == "debug" {
		if m.Value.ToString() == "" {
			panic("env: key " + key + " not found")
		}
	}
	return m
}

func newValue(v string) value {
	return value(v)
}

type manager[T any] struct {
	Key          string
	Value        value
	DefaultValue T
	HasDefault   bool
	Must         bool
}

func (m *manager[T]) EmptyValue() bool {
	return m.Value == ""
}

func (m *manager[T]) ReturnDefault() T {
	if m.HasDefault {
		return m.DefaultValue
	}
	return *new(T)
}

type value string

func (v value) ToString() string {
	return string(v)
}

func (v value) ToStrings() []string {
	return strings.Split(string(v), ",")
}

func (v value) ToInt() int {
	output, err := strconv.Atoi(string(v))
	if err != nil {
		panic("env: value is not an int.")
	}
	return output
}

func (v value) ToInts() []int {
	var output []int
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToInt())
	}
	return output
}

func (v value) ToInt64() int64 {
	output, err := strconv.ParseInt(string(v), 10, 64)
	if err != nil {
		panic("env: value is not an int64.")
	}
	return output
}

func (v value) ToInt64s() []int64 {
	var output []int64
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToInt64())
	}
	return output
}

func (v value) ToInt32() int32 {
	output, err := strconv.ParseInt(string(v), 10, 32)
	if err != nil {
		panic("env: value is not an int32.")
	}
	return int32(output)
}

func (v value) ToInt32s() []int32 {
	var output []int32
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToInt32())
	}
	return output
}

func (v value) ToInt16() int16 {
	output, err := strconv.ParseInt(string(v), 10, 16)
	if err != nil {
		panic("env: value is not an int16.")
	}
	return int16(output)
}

func (v value) ToInt16s() []int16 {
	var output []int16
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToInt16())
	}
	return output
}

func (v value) ToInt8() int8 {
	output, err := strconv.ParseInt(string(v), 10, 8)
	if err != nil {
		panic("env: value is not an int8.")
	}
	return int8(output)
}

func (v value) ToInt8s() []int8 {
	var output []int8
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToInt8())
	}
	return output
}

func (v value) ToUint() uint {
	output, err := strconv.ParseUint(string(v), 10, 64)
	if err != nil {
		panic("env: value is not an uint.")
	}
	return uint(output)
}

func (v value) ToUints() []uint {
	var output []uint
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToUint())
	}
	return output
}

func (v value) ToUint64() uint64 {
	output, err := strconv.ParseUint(string(v), 10, 64)
	if err != nil {
		panic("env: value is not an uint64.")
	}
	return output
}

func (v value) ToUint64s() []uint64 {
	var output []uint64
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToUint64())
	}
	return output
}

func (v value) ToUint32() uint32 {
	output, err := strconv.ParseUint(string(v), 10, 32)
	if err != nil {
		panic("env: value is not an uint32.")
	}
	return uint32(output)
}

func (v value) ToUint32s() []uint32 {
	var output []uint32
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToUint32())
	}
	return output
}

func (v value) ToUint16() uint16 {
	output, err := strconv.ParseUint(string(v), 10, 16)
	if err != nil {
		panic("env: value is not an uint16.")
	}
	return uint16(output)
}

func (v value) ToUint16s() []uint16 {
	var output []uint16
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToUint16())
	}
	return output
}

func (v value) ToUint8() uint8 {
	output, err := strconv.ParseUint(string(v), 10, 8)
	if err != nil {
		panic("env: value is not an uint8.")
	}
	return uint8(output)
}

func (v value) ToUint8s() []uint8 {
	var output []uint8
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToUint8())
	}
	return output
}

func (v value) ToFloat32() float32 {
	output, err := strconv.ParseFloat(string(v), 32)
	if err != nil {
		panic("env: value is not an float32.")
	}
	return float32(output)
}

func (v value) ToFloat32s() []float32 {
	var output []float32
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToFloat32())
	}
	return output
}

func (v value) ToFloat64() float64 {
	output, err := strconv.ParseFloat(string(v), 64)
	if err != nil {
		panic("env: value is not an float64.")
	}
	return output
}

func (v value) ToFloat64s() []float64 {
	var output []float64
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToFloat64())
	}
	return output
}

func (v value) ToBool() bool {
	output, err := strconv.ParseBool(string(v))
	if err != nil {
		panic("env: value is not an bool.")
	}
	return output
}

func (v value) ToBools() []bool {
	var output []bool
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToBool())
	}
	return output
}

func (v value) ToByte() byte {
	output, err := strconv.ParseUint(string(v), 10, 8)
	if err != nil {
		panic("env: value is not an byte.")
	}
	return byte(output)
}

func (v value) ToBytes() []byte {
	var output []byte
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToByte())
	}
	return output
}

func (v value) ToRune() rune {
	output, err := strconv.ParseInt(string(v), 10, 32)
	if err != nil {
		panic("env: value is not an rune.")
	}
	return rune(output)
}

func (v value) ToRunes() []rune {
	var output []rune
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToRune())
	}
	return output
}

func (v value) ToComplex64() complex64 {
	output, err := strconv.ParseComplex(string(v), 64)
	if err != nil {
		panic("env: value is not an complex64.")
	}
	return complex64(output)
}

func (v value) ToComplex64s() []complex64 {
	var output []complex64
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToComplex64())
	}
	return output
}

func (v value) ToComplex128() complex128 {
	output, err := strconv.ParseComplex(string(v), 128)
	if err != nil {
		panic("env: value is not an complex128.")
	}
	return complex128(output)
}

func (v value) ToComplex128s() []complex128 {
	var output []complex128
	for _, v := range v.ToStrings() {
		output = append(output, newValue(v).ToComplex128())
	}
	return output
}
