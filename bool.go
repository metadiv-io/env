package env

/*
Bool returns the value of the environment variable named by the key.
*/
func Bool(key string, defaultValue ...bool) bool {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToBool()
}

/*
Bools returns the list of values of the environment variable named by the key.
*/
func Bools(key string, defaultValue ...[]bool) []bool {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToBools()
}

/*
MustBool returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustBool(key string) bool {
	m := newMustManager[bool](key)
	return m.Value.ToBool()
}

/*
MustBools returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustBools(key string) []bool {
	m := newMustManager[bool](key)
	return m.Value.ToBools()
}
