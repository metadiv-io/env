package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestInt16(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int16("TEST_INT16")
	require.Equal(t, int16(10), v)
}

func TestInt16Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int16("TEST_INT16_NOT_EXISTS", 10)
	require.Equal(t, int16(10), v)
}

func TestInt16Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustInt16("TEST_INT16")
	require.Equal(t, int16(10), v)
}

func TestInt16MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustInt16("TEST_INT16_NOT_EXISTS")
	})
}

func TestInt16s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int16s("TEST_INT16S")
	require.Equal(t, []int16{10, 100}, v)
}

func TestInt16sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int16s("TEST_INT16S_NOT_EXISTS", []int16{10, 100})
	require.Equal(t, []int16{10, 100}, v)
}

func TestInt16sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustInt16s("TEST_INT16S")
	require.Equal(t, []int16{10, 100}, v)
}

func TestInt16sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustInt16s("TEST_INT16S_NOT_EXISTS")
	})
}
