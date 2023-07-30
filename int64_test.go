package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestInt64(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int64("TEST_INT64")
	require.Equal(t, int64(1000), v)
}

func TestInt64Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int64("TEST_INT64_NOT_EXISTS", 1000)
	require.Equal(t, int64(1000), v)
}

func TestInt64Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustInt64("TEST_INT64")
	require.Equal(t, int64(1000), v)
}

func TestInt64MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustInt64("TEST_INT64_NOT_EXISTS")
	})
}

func TestInt64s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int64s("TEST_INT64S")
	require.Equal(t, []int64{1000, 10000}, v)
}

func TestInt64sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int64s("TEST_INT64S_NOT_EXISTS", []int64{1000, 10000})
	require.Equal(t, []int64{1000, 10000}, v)
}

func TestInt64sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustInt64s("TEST_INT64S")
	require.Equal(t, []int64{1000, 10000}, v)
}

func TestInt64sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustInt64s("TEST_INT64S_NOT_EXISTS")
	})
}
