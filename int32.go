package env

import "github.com/metadiv-io/env/types"

/*
Int32 returns the value of the environment variable named by the key.
*/
func Int32(key string, defaultValue ...int32) int32 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt32()
}
