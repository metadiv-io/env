package env

/*
Uint returns the value of the environment variable named by the key.
*/
func Uint(key string, defaultValue ...uint) uint {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint()
}

/*
Uints returns the list of values of the environment variable named by the key.
*/
func Uints(key string, defaultValue ...[]uint) []uint {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUints()
}

/*
MustUint returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint(key string) uint {
	m := newMustManager[uint](key)
	return m.Value.ToUint()
}

/*
MustUints returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUints(key string) []uint {
	m := newMustManager[uint](key)
	return m.Value.ToUints()
}

/*
Uint8 returns the value of the environment variable named by the key.
*/
func Uint8(key string, defaultValue ...uint8) uint8 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint8()
}

/*
Uint8s returns the list of values of the environment variable named by the key.
*/
func Uint8s(key string, defaultValue ...[]uint8) []uint8 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint8s()
}

/*
MustUint8 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint8(key string) uint8 {
	m := newMustManager[uint8](key)
	return m.Value.ToUint8()
}

/*
MustUint8s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint8s(key string) []uint8 {
	m := newMustManager[uint8](key)
	return m.Value.ToUint8s()
}

/*
Uint16 returns the value of the environment variable named by the key.
*/
func Uint16(key string, defaultValue ...uint16) uint16 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint16()
}

/*
Uint16s returns the list of values of the environment variable named by the key.
*/
func Uint16s(key string, defaultValue ...[]uint16) []uint16 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint16s()
}

/*
MustUint16 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint16(key string) uint16 {
	m := newMustManager[uint16](key)
	return m.Value.ToUint16()
}

/*
MustUint16s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint16s(key string) []uint16 {
	m := newMustManager[uint16](key)
	return m.Value.ToUint16s()
}

/*
Uint32 returns the value of the environment variable named by the key.
*/
func Uint32(key string, defaultValue ...uint32) uint32 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint32()
}

/*
Uint32s returns the list of values of the environment variable named by the key.
*/
func Uint32s(key string, defaultValue ...[]uint32) []uint32 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint32s()
}

/*
MustUint32 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint32(key string) uint32 {
	m := newMustManager[uint32](key)
	return m.Value.ToUint32()
}

/*
MustUint32s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint32s(key string) []uint32 {
	m := newMustManager[uint32](key)
	return m.Value.ToUint32s()
}

/*
Uint64 returns the value of the environment variable named by the key.
*/
func Uint64(key string, defaultValue ...uint64) uint64 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint64()
}

/*
Uint64s returns the list of values of the environment variable named by the key.
*/
func Uint64s(key string, defaultValue ...[]uint64) []uint64 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint64s()
}

/*
MustUint64 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint64(key string) uint64 {
	m := newMustManager[uint64](key)
	return m.Value.ToUint64()
}

/*
MustUint64s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint64s(key string) []uint64 {
	m := newMustManager[uint64](key)
	return m.Value.ToUint64s()
}
