package env

/*
String returns the value of the environment variable named by the key.
*/
func String(key string, defaultValue ...string) string {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToString()
}

/*
Strings returns the list of values of the environment variable named by the key.
*/
func Strings(key string, defaultValue ...[]string) []string {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToStrings()
}

/*
MustString returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustString(key string) string {
	m := newMustManager[string](key)
	return m.Value.ToString()
}

/*
MustStrings returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustStrings(key string) []string {
	m := newMustManager[string](key)
	return m.Value.ToStrings()
}
