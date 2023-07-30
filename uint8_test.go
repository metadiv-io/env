package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestUint8(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint8("TEST_UINT8")
	require.Equal(t, uint8(1), v)
}

func TestUint8Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint8("TEST_UINT8_NOT_EXISTS", 1)
	require.Equal(t, uint8(1), v)
}

func TestUint8Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustUint8("TEST_UINT8")
	require.Equal(t, uint8(1), v)
}

func TestUint8MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustUint8("TEST_UINT8_NOT_EXISTS")
	})
}

func TestUint8s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint8s("TEST_UINT8S")
	require.Equal(t, []uint8{1, 2}, v)
}

func TestUint8sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint8s("TEST_UINT8S_NOT_EXISTS", []uint8{1, 2})
	require.Equal(t, []uint8{1, 2}, v)
}

func TestUint8sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustUint8s("TEST_UINT8S")
	require.Equal(t, []uint8{1, 2}, v)
}

func TestUint8sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustUint8s("TEST_UINT8S_NOT_EXISTS")
	})
}
