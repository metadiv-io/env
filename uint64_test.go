package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestUint64(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint64("TEST_UINT64")
	require.Equal(t, uint64(1), v)
}

func TestUint64Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint64("TEST_UINT64_NOT_EXISTS", 1)
	require.Equal(t, uint64(1), v)
}

func TestUint64Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustUint64("TEST_UINT64")
	require.Equal(t, uint64(1), v)
}

func TestUint64MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustUint64("TEST_UINT64_NOT_EXISTS")
	})
}

func TestUint64s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint64s("TEST_UINT64S")
	require.Equal(t, []uint64{1, 2}, v)
}

func TestUint64sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint64s("TEST_UINT64S_NOT_EXISTS", []uint64{1, 2})
	require.Equal(t, []uint64{1, 2}, v)
}

func TestUint64sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustUint64s("TEST_UINT64S")
	require.Equal(t, []uint64{1, 2}, v)
}

func TestUint64sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustUint64s("TEST_UINT64S_NOT_EXISTS")
	})
}
