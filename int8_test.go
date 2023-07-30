package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestInt8(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int8("TEST_INT8")
	require.Equal(t, int8(1), v)
}

func TestInt8Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int8("TEST_INT8_NOT_EXISTS", 1)
	require.Equal(t, int8(1), v)
}

func TestInt8Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustInt8("TEST_INT8")
	require.Equal(t, int8(1), v)
}

func TestInt8MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustInt8("TEST_INT8_NOT_EXISTS")
	})
}

func TestInt8s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int8s("TEST_INT8S")
	require.Equal(t, []int8{1, 10}, v)
}

func TestInt8sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Int8s("TEST_INT8S_NOT_EXISTS", []int8{1, 10})
	require.Equal(t, []int8{1, 10}, v)
}

func TestInt8sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustInt8s("TEST_INT8S")
	require.Equal(t, []int8{1, 10}, v)
}

func TestInt8sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustInt8s("TEST_INT8S_NOT_EXISTS")
	})
}
