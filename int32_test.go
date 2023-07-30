package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestInt32(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int32("TEST_INT32")
	require.Equal(t, int32(100), v)
}

func TestInt32Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int32("TEST_INT32_NOT_EXISTS", 100)
	require.Equal(t, int32(100), v)
}

func TestInt32Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustInt32("TEST_INT32")
	require.Equal(t, int32(100), v)
}

func TestInt32MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustInt32("TEST_INT32_NOT_EXISTS")
	})
}

func TestInt32s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int32s("TEST_INT32S")
	require.Equal(t, []int32{100, 1000}, v)
}

func TestInt32sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int32s("TEST_INT32S_NOT_EXISTS", []int32{100, 1000})
	require.Equal(t, []int32{100, 1000}, v)
}

func TestInt32sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustInt32s("TEST_INT32S")
	require.Equal(t, []int32{100, 1000}, v)
}

func TestInt32sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustInt32s("TEST_INT32S_NOT_EXISTS")
	})
}
