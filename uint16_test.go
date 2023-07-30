package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestUint16(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint16("TEST_UINT16")
	require.Equal(t, uint16(1), v)
}

func TestUint16Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint16("TEST_UINT16_NOT_EXISTS", 1)
	require.Equal(t, uint16(1), v)
}

func TestUint16Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustUint16("TEST_UINT16")
	require.Equal(t, uint16(1), v)
}

func TestUint16MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustUint16("TEST_UINT16_NOT_EXISTS")
	})
}

func TestUint16s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint16s("TEST_UINT16S")
	require.Equal(t, []uint16{1, 2}, v)
}

func TestUint16sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Uint16s("TEST_UINT16S_NOT_EXISTS", []uint16{1, 2})
	require.Equal(t, []uint16{1, 2}, v)
}

func TestUint16sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustUint16s("TEST_UINT16S")
	require.Equal(t, []uint16{1, 2}, v)
}

func TestUint16sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustUint16s("TEST_UINT16S_NOT_EXISTS")
	})
}
