package env

/*
Int returns the value of the environment variable named by the key.
*/
func Int(key string, defaultValue ...int) int {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt()
}

/*
Ints returns the list of values of the environment variable named by the key.
*/
func Ints(key string, defaultValue ...[]int) []int {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInts()
}

/*
MustInt returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt(key string) int {
	m := newMustManager[int](key)
	return m.Value.ToInt()
}

/*
MustInts returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInts(key string) []int {
	m := newMustManager[int](key)
	return m.Value.ToInts()
}

/*
Int8 returns the value of the environment variable named by the key.
*/
func Int8(key string, defaultValue ...int8) int8 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt8()
}

/*
Int8s returns the list of values of the environment variable named by the key.
*/
func Int8s(key string, defaultValue ...[]int8) []int8 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt8s()
}

/*
MustInt8 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt8(key string) int8 {
	m := newMustManager[int8](key)
	return m.Value.ToInt8()
}

/*
MustInt8s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt8s(key string) []int8 {
	m := newMustManager[int8](key)
	return m.Value.ToInt8s()
}

/*
Int16 returns the value of the environment variable named by the key.
*/
func Int16(key string, defaultValue ...int16) int16 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt16()
}

/*
Int16s returns the list of values of the environment variable named by the key.
*/
func Int16s(key string, defaultValue ...[]int16) []int16 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt16s()
}

/*
MustInt16 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt16(key string) int16 {
	m := newMustManager[int16](key)
	return m.Value.ToInt16()
}

/*
MustInt16s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt16s(key string) []int16 {
	m := newMustManager[int16](key)
	return m.Value.ToInt16s()
}

/*
Int32 returns the value of the environment variable named by the key.
*/
func Int32(key string, defaultValue ...int32) int32 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt32()
}

/*
Int32s returns the value of the environment variable named by the key.
*/
func Int32s(key string, defaultValue ...[]int32) []int32 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt32s()
}

/*
MustInt32 returns the value of the environment variable named by the key.
If the value is empty, it will panic.
*/
func MustInt32(key string) int32 {
	m := newMustManager[int16](key)
	return m.Value.ToInt32()
}

/*
MustInt32s returns the value of the environment variable named by the key.
If the value is empty, it will panic.
*/
func MustInt32s(key string) []int32 {
	m := newMustManager[int16](key)
	return m.Value.ToInt32s()
}

/*
Int64 returns the value of the environment variable named by the key.
*/
func Int64(key string, defaultValue ...int64) int64 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt64()
}

/*
Int64s returns the list of values of the environment variable named by the key.
*/
func Int64s(key string, defaultValue ...[]int64) []int64 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt64s()
}

/*
MustInt64 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt64(key string) int64 {
	m := newMustManager[int64](key)
	return m.Value.ToInt64()
}

/*
MustInt64s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt64s(key string) []int64 {
	m := newMustManager[int64](key)
	return m.Value.ToInt64s()
}
