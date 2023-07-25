package env

import "github.com/metadiv-io/env/types"

/*
Byte returns the value of the environment variable named by the key.
*/
func Byte(key string, defaultValue ...byte) byte {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToByte()
}

/*
Bytes returns the list of values of the environment variable named by the key.
*/
func Bytes(key string, defaultValue ...[]byte) []byte {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToBytes()
}

/*
MustByte returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustByte(key string) byte {
	m := types.NewMustManager[byte](key)
	return m.Value.ToByte()
}

/*
MustBytes returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustBytes(key string) []byte {
	m := types.NewMustManager[byte](key)
	return m.Value.ToBytes()
}
