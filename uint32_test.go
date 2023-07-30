package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestUint32(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint32("TEST_UINT32")
	require.Equal(t, uint32(1), v)
}

func TestUint32Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint32("TEST_UINT32_NOT_EXISTS", 1)
	require.Equal(t, uint32(1), v)
}

func TestUint32Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustUint32("TEST_UINT32")
	require.Equal(t, uint32(1), v)
}

func TestUint32MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustUint32("TEST_UINT32_NOT_EXISTS")
	})
}

func TestUint32s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint32s("TEST_UINT32S")
	require.Equal(t, []uint32{1, 2}, v)
}

func TestUint32sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint32s("TEST_UINT32S_NOT_EXISTS", []uint32{1, 2})
	require.Equal(t, []uint32{1, 2}, v)
}

func TestUint32sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustUint32s("TEST_UINT32S")
	require.Equal(t, []uint32{1, 2}, v)
}

func TestUint32sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustUint32s("TEST_UINT32S_NOT_EXISTS")
	})
}
