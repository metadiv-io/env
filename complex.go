package env

/*
Complex64 returns the value of the environment variable named by the key.
*/
func Complex64(key string, defaultValue ...complex64) complex64 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToComplex64()
}

/*
Complex64s returns the list of values of the environment variable named by the key.
*/
func Complex64s(key string, defaultValue ...[]complex64) []complex64 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToComplex64s()
}

/*
MustComplex64 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustComplex64(key string) complex64 {
	m := newMustManager[complex64](key)
	return m.Value.ToComplex64()
}

/*
MustComplex64s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustComplex64s(key string) []complex64 {
	m := newMustManager[complex64](key)
	return m.Value.ToComplex64s()
}

/*
Complex128 returns the value of the environment variable named by the key.
*/
func Complex128(key string, defaultValue ...complex128) complex128 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToComplex128()
}

/*
Complex128s returns the list of values of the environment variable named by the key.
*/
func Complex128s(key string, defaultValue ...[]complex128) []complex128 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToComplex128s()
}

/*
MustComplex128 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustComplex128(key string) complex128 {
	m := newMustManager[complex128](key)
	return m.Value.ToComplex128()
}

/*
MustComplex128s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustComplex128s(key string) []complex128 {
	m := newMustManager[complex128](key)
	return m.Value.ToComplex128s()
}
