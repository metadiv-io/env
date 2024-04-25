package env

/*
Float32 returns the value of the environment variable named by the key.
*/
func Float32(key string, defaultValue ...float32) float32 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToFloat32()
}

/*
Float32s returns the list of values of the environment variable named by the key.
*/
func Float32s(key string, defaultValue ...[]float32) []float32 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToFloat32s()
}

/*
MustFloat32 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustFloat32(key string) float32 {
	m := newMustManager[float32](key)
	return m.Value.ToFloat32()
}

/*
MustFloat32s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustFloat32s(key string) []float32 {
	m := newMustManager[float32](key)
	return m.Value.ToFloat32s()
}

/*
Float64 returns the value of the environment variable named by the key.
*/
func Float64(key string, defaultValue ...float64) float64 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToFloat64()
}

/*
Float64s returns the list of values of the environment variable named by the key.
*/
func Float64s(key string, defaultValue ...[]float64) []float64 {
	m := newManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToFloat64s()
}

/*
MustFloat64 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustFloat64(key string) float64 {
	m := newMustManager[float64](key)
	return m.Value.ToFloat64()
}

/*
MustFloat64s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustFloat64s(key string) []float64 {
	m := newMustManager[float64](key)
	return m.Value.ToFloat64s()
}
