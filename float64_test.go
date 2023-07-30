package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestFloat64(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Float64("TEST_FLOAT64")
	require.Equal(t, float64(6.4), v)
}

func TestFloat64Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Float64("TEST_FLOAT64_NOT_EXISTS", 6.4)
	require.Equal(t, float64(6.4), v)
}

func TestFloat64Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustFloat64("TEST_FLOAT64")
	require.Equal(t, float64(6.4), v)
}

func TestFloat64MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustFloat64("TEST_FLOAT64_NOT_EXISTS")
	})
}

func TestFloat64s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Float64s("TEST_FLOAT64S")
	require.Equal(t, []float64{6.4, 4.6}, v)
}

func TestFloat64sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Float64s("TEST_FLOAT64S_NOT_EXISTS", []float64{6.4, 4.6})
	require.Equal(t, []float64{6.4, 4.6}, v)
}

func TestFloat64sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustFloat64s("TEST_FLOAT64S")
	require.Equal(t, []float64{6.4, 4.6}, v)
}

func TestFloat64sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustFloat64s("TEST_FLOAT64S_NOT_EXISTS")
	})
}
