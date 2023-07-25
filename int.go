package env

import "github.com/metadiv-io/env/types"

/*
Int returns the value of the environment variable named by the key.
*/
func Int(key string, defaultValue ...int) int {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt()
}

/*
Ints returns the list of values of the environment variable named by the key.
*/
func Ints(key string, defaultValue ...[]int) []int {
	m := types.NewManager(key, defaultValue...)
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
	m := types.NewMustManager[int](key)
	return m.Value.ToInt()
}

/*
MustInts returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInts(key string) []int {
	m := types.NewMustManager[int](key)
	return m.Value.ToInts()
}
