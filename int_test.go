package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int("TEST_INT")
	require.Equal(t, int(10), v)
}

func TestIntDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int("TEST_INT_NOT_EXISTS", 10)
	require.Equal(t, int(10), v)
}

func TestIntMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustInt("TEST_INT")
	require.Equal(t, int(10), v)
}

func TestIntMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustInt("TEST_INT_NOT_EXISTS")
	})
}

func TestInts(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Ints("TEST_INTS")
	require.Equal(t, []int{10, 100}, v)
}

func TestIntsDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Ints("TEST_INTS_NOT_EXISTS", []int{10, 100})
	require.Equal(t, []int{10, 100}, v)
}

func TestIntsMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustInts("TEST_INTS")
	require.Equal(t, []int{10, 100}, v)
}

func TestIntsMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustInts("TEST_INTS_NOT_EXISTS")
	})
}
