package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestUint(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint("TEST_UINT")
	require.Equal(t, uint(1), v)
}

func TestUintDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint("TEST_UINT_NOT_EXISTS", 1)
	require.Equal(t, uint(1), v)
}

func TestUintMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustUint("TEST_UINT")
	require.Equal(t, uint(1), v)
}

func TestUintMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustUint("TEST_UINT_NOT_EXISTS")
	})
}

func TestUints(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uints("TEST_UINTS")
	require.Equal(t, []uint{1, 2}, v)
}

func TestUintsDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uints("TEST_UINTS_NOT_EXISTS", []uint{1, 2})
	require.Equal(t, []uint{1, 2}, v)
}

func TestUintsMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustUints("TEST_UINTS")
	require.Equal(t, []uint{1, 2}, v)
}

func TestUintsMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustUints("TEST_UINTS_NOT_EXISTS")
	})
}
