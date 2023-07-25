package env

import "github.com/metadiv-io/env/types"

/*
Rune returns the value of the environment variable named by the key.
*/
func Rune(key string, defaultValue ...rune) rune {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToRune()
}

/*
Runes returns the list of values of the environment variable named by the key.
*/
func Runes(key string, defaultValue ...[]rune) []rune {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToRunes()
}

/*
MustRune returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustRune(key string) rune {
	m := types.NewMustManager[rune](key)
	return m.Value.ToRune()
}

/*
MustRunes returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustRunes(key string) []rune {
	m := types.NewMustManager[rune](key)
	return m.Value.ToRunes()
}
